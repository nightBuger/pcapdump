package pcapimpl

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"pcapdump/parser/v2v"
)

type Dumper struct {
	dataSourceHandler *pcap.Handle
	devName           string
	stopCh            chan struct{}
	layerType         gopacket.LayerType
}

type Parser struct {
	Name    string
	Decoder gopacket.DecodeFunc
	TypeId  int
}

func (this *Dumper) ToString() string {
	infoString := fmt.Sprintf("interface name :%s\n", IfEmptyToNil(this.devName))

	return infoString
}

func GetDevNameSlice(string) (devNameSlice []string) {
	devNameSlice, err := getDevNameSlice()
	if err != nil {
		panic(err)
	}
	return
}

func (this *Dumper) SetInterface(devName string) (err error) {
	if this.dataSourceHandler != nil {
		this.dataSourceHandler.Close()
	}
	this.dataSourceHandler = nil
	this.devName = ""
	this.dataSourceHandler, err = openDev(devName)
	if err != nil {
		return errors.New("无法打开设备 " + err.Error())
	}
	this.devName = devName
	return
}

func (this *Dumper) Run() {
	if this.dataSourceHandler == nil {
		this.emitError("请先指定一个数据源(一个网卡或者pcap文件),再执行dump run命令")
		return
	}
	if this.stopCh != nil {
		this.emitError("已经启动了抓包,请勿重复开启")
		return
	}
	this.stopCh = make(chan struct{}, 0)
	go this.dumpThread()
}

func (this *Dumper) Stop() {
	if this.stopCh != nil {
		this.stopCh <- struct{}{}
	}
	this.stopCh = nil

}

func (this *Dumper) emitError(errorText string) {
	fmt.Println(errorText)
}
func (this *Dumper) emitInfo(infoText string) {
	fmt.Println(infoText)
}

func (this *Dumper) dumpThread() {
	dataSource := gopacket.NewPacketSource(this.dataSourceHandler, this.layerType)
	this.emitInfo("==========抓包进程启动==========")
	for {
		select {
		case pack := <-dataSource.Packets():
			this.parse(pack)
		case <-this.stopCh:
			goto stop
		}
	}
stop:
	this.emitInfo("==========抓包进程结束==========")
}

func (this *Dumper) parse(pack gopacket.Packet) {

	ethLayer := pack.Layer(v2v.V2VEthType)
	if ethLayer == nil {
		return
	}
	v2v1Layer := pack.Layer(v2v.V2V1LayerType)
	if v2v1Layer == nil {
		return
	}
	v2v2Layer := pack.Layer(v2v.V2V2LayerType)
	if v2v2Layer == nil {
		return
	}

	//转换interface到真实的指针类型
	pEth := ethLayer.(*v2v.V2VEth)
	pV1 := v2v1Layer.(*v2v.V2V1Layer)
	pV2 := v2v2Layer.(*v2v.V2V2Layer)

	//check合法包
	switch pV1.WhichPacketType() {
	case v2v.LinkPacket:
		switch pV2.CmdId {
		case v2v.V2VLink:
		case v2v.V2VLinkRes:
		case v2v.V2VAuth:
		case v2v.V2VAuthRes:
		case v2v.V2VLogin:
		case v2v.V2VLoginRes:
		default:
			return
		}
	case v2v.UniCastPacket:
		switch pV2.CmdId {
		case v2v.V2VHeart:
		case v2v.V2VHeartRes:
		default:
			return
		}
	}
	//filter过滤器
	if pV1.WhichPacketType() != v2v.UniCastPacket {
		return
	}
	// dispaly部分

	fmt.Printf("dst=%s src=%s type=0x%02x cmd=0x%04x\n",
		PrintByteToHex(pEth.EthDst),
		PrintByteToHex(pEth.EthSrc),
		pV1.WhichPacketType(),
		pV2.CmdId)
	return
}

func (this *Dumper) RegisterParser(layerType gopacket.LayerType) {
	this.layerType = layerType
	return
}

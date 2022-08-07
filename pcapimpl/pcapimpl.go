package pcapimpl

import (
	"errors"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
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

	ethLayer := pack.Layer(layers.LayerTypeEthernet)
	if ethLayer == nil {
		return
	}
	fmt.Println("ethLayer header:")
	PrintByteToHex(ethLayer.LayerContents())
	v2vLayer := pack.Layer(this.layerType)
	if v2vLayer == nil {
		fmt.Println("不是视联网包")
		return
	} else {
		fmt.Println("真的是视联网包!!!")
	}

	return
}

func (this *Dumper) RegisterParser(layerType gopacket.LayerType) {
	this.layerType = layerType
	return
}

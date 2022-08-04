package pcapimpl

import (
	"errors"
	"fmt"
	"github.com/google/gopacket/pcap"
)

type Dumper struct {
	dataSourceHandler *pcap.Handle
	devName           string
}

func (this *Dumper) ToString() string {
	infoString := fmt.Sprintf("interface name :%s\n", IfEmptyToNil(this.devName))

	return infoString
}

func (*Dumper) GetDevNameSlice(string) (devNameSlice []string) {
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

	go func() {

	}()
}

func (this *Dumper) emitError(errorText string) {
	fmt.Println(errorText)
}

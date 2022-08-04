package pcapimpl

import (
	"github.com/google/gopacket/pcap"
)

//获得本机上所有网卡列表
func GetDeviceList() (devs []pcap.Interface, err error) {
	devs, err = pcap.FindAllDevs()
	return
}

func openDev(devName string) (*pcap.Handle, error) {
	handler, err := pcap.OpenLive(devName, 65535, true, pcap.BlockForever)
	return handler, err
}

package pcapimpl

import (
	"fmt"
	"github.com/google/gopacket"
	_ "github.com/google/gopacket"
	_ "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"os"
)

func init() {
	return
	devs, err := pcap.FindAllDevs()
	if err != nil {
		println("pcap.FindAllDevs() error:", err.Error())
		os.Exit(1)
	}
	for i := 0; i < len(devs); i++ {
		fmt.Println(devs[i].Name, devs[i].Description)
	}

	handler, err := pcap.OpenLive("\\Device\\NPF_{03ABB600-9AA4-4F1F-9121-6BCE21AD7C89}", 65535, true, pcap.BlockForever)
	if err != nil {
		println("pcap.OpenLive() error:", err.Error())
		os.Exit(1)
	}
	defer handler.Close()
	ds := gopacket.NewPacketSource(handler, handler.LinkType())
	ds.NoCopy = true
	for pack := range ds.Packets() {
		println(pack.Dump())
		networkPack := pack.NetworkLayer()
		if networkPack == nil {
			continue
		}
		printByteToHex(pack.Data())
		fmt.Println("++++++++++++")
		printByteToHex(networkPack.LayerContents())
		fmt.Println("++++++++++++")
		printByteToHex(networkPack.LayerPayload())
		fmt.Println("=======================")
	}
}

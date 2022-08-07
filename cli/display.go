package cli

import (
	"fmt"
	"github.com/google/gopacket/pcap"
	"os"
	"pcapdump/pcapimpl"
)

func showDevList() {
	devs, err := pcapimpl.GetDeviceList()
	if err != nil {
		fmt.Println("无法获取网卡列表,err:", err.Error())
		os.Exit(200)
	}

	for i := 0; i < len(devs); i++ {
		fmt.Printf("==========网卡%d==========\n", i+1)
		PrintDev(devs[i])
	}

}

func PrintDev(inter pcap.Interface) {
	//info := fmt.Sprintf("Name:%s\nDescription:%s\nFlags:%s\n", inter.Name, inter.Description, inter.Flags)
	info := fmt.Sprintf("Name:%s\n", inter.Name)
	for i := 0; i < len(inter.Addresses); i++ {
		a, _ := inter.Addresses[i].Netmask.Size()
		info += fmt.Sprintf("ip%d:%s/%d\n", i+1, inter.Addresses[i].IP, a)
	}
	fmt.Println(info)
}

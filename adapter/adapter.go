package adapter

import (
	"github.com/google/gopacket"
	"pcapdump/parser"
)

type Adapter interface {
	Regist(args ...parser.LayerType)
	Parse(pack gopacket.Packet)
	filter() bool
	display()
}

var autoId = 1000

func golbalLayerId() int {
	autoId++
	if autoId > 1999 {
		panic("注册太多LayerType了,不可能")
	}
	return autoId
}

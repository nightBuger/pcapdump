package adapter

import (
	"github.com/google/gopacket"
	"pcapdump/parser"
)

type V2VLoginParser struct {
}

func (this *V2VLoginParser) Regist(args ...parser.LayerType) {
	var next gopacket.LayerType = -1
	for i := len(args) - 1; i >= 0; i-- {
		next = gopacket.RegisterLayerType(golbalLayerId(), gopacket.LayerTypeMetadata{
			Name:    args[i].GetDecodeName(),
			Decoder: args[i].GenDecodeFunc(next),
		})
	}
}

func (this *V2VLoginParser) Parse(pack gopacket.Packet) {
	//TODO implement me
	panic("implement me")
}

func (this *V2VLoginParser) filter() bool {
	//TODO implement me
	panic("implement me")
}

func (this *V2VLoginParser) display() {
	//TODO implement me
	panic("implement me")
}

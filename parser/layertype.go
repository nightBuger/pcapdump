package parser

import "github.com/google/gopacket"

type LayerType interface {
	GenDecodeFunc(lt gopacket.LayerType) gopacket.DecodeFunc
	GetDecodeName() string
}

package parser

import "github.com/google/gopacket"

type layerType interface {
	GenDecodeFunc(lt gopacket.LayerType) gopacket.DecodeFunc
}

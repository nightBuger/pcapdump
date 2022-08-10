package v2v

import (
	"encoding/binary"
	"errors"
	"github.com/google/gopacket"
)

type V2VEth struct {
	Header  []byte
	Payload []byte

	EthType uint16
	EthSrc  []byte
	EthDst  []byte
}

const (
	ETH0800    uint16 = 0x0800
	ETH0806    uint16 = 0x0806
	ETH909B    uint16 = 0x909B
	ETH8100    uint16 = 0x8100
	ETHUNKNOWN uint16 = 0x0000
)
const (
	EthHeaderSize     uint32 = 14
	EthVlanHeaderSize uint32 = 18
)

func (this *V2VEth) LayerType() gopacket.LayerType { return V2VEthType }
func (this *V2VEth) LayerContents() []byte         { return this.Header }
func (this *V2VEth) LayerPayload() []byte          { return this.Payload }

var (
	V2VEthType = gopacket.RegisterLayerType(1000, gopacket.LayerTypeMetadata{
		Name:    "V2VEthType",
		Decoder: gopacket.DecodeFunc(DecodeV2VEth),
	})
)

func DecodeV2VEth(data []byte, p gopacket.PacketBuilder) error {
	layer := newV2VEthLayer(data)
	if layer == nil {
		return errors.New("")
	}
	p.AddLayer(layer)
	return p.NextDecoder(V2V1LayerType)
}

func getEthType(data []byte) (uint16, uint32) {
	var ethType uint16
	ethType = binary.BigEndian.Uint16(data[12:14])
	headerSize := EthHeaderSize // 不带vlan时 标准
	if ethType == ETH8100 {
		ethType = binary.BigEndian.Uint16(data[16:18])
		headerSize = EthVlanHeaderSize
	}
	switch ethType {
	case ETH0800:
	case ETH0806:
	case ETH909B:
	default:
		ethType = ETHUNKNOWN
	}
	return ethType, headerSize

}

//创建一个v2vEth指针, 并检查以太网类型字段,自适应以带vlan的太网包头长度,同时,只有0x0800 0x0806 0x909b会认为是正确的包,其余的包返回nil
func newV2VEthLayer(data []byte) *V2VEth {
	ethType, size := getEthType(data)
	if ethType == ETHUNKNOWN {
		return nil
	}
	ret := &V2VEth{
		EthSrc: data[6:12],
		EthDst: data[0:6],
	}
	ret.EthType = ethType
	ret.Header = data[0:size]
	ret.Payload = data[size:]
	return ret
}

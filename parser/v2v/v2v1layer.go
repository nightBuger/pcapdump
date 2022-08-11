package v2v

import (
	"errors"
	"github.com/google/gopacket"
)

type V2V1Layer struct {
	Header  []byte
	Payload []byte
}

func (this *V2V1Layer) LayerType() gopacket.LayerType { return V2V1LayerType }
func (this *V2V1Layer) LayerContents() []byte         { return this.Header }
func (this *V2V1Layer) LayerPayload() []byte          { return this.Payload }

var (
	V2V1LayerType = gopacket.RegisterLayerType(1001, gopacket.LayerTypeMetadata{
		Name:    "V2V1LayerType",
		Decoder: gopacket.DecodeFunc(DecodeV2V1Layer),
	})
)

type PacketType byte

const (
	PackTypeField, ePackTypeFieldSize = iota, 1 // byte0
)
const (
	UnknowPacket  PacketType = 0xff // 不认识的包
	LinkPacket    PacketType = 0x10 // 0类连接包
	UniCastPacket PacketType = 0x02 // 2类单播包
)

const headerSize int = 22

func V2V1HeaderSize() int {
	return headerSize
}

func (this *V2V1Layer) WhichPacketType() PacketType {
	switch PacketType(this.Header[PackTypeField]) {
	case LinkPacket:
	case UniCastPacket:
	default:
		return UnknowPacket
	}
	return PacketType(this.Header[PackTypeField])
}

func DecodeV2V1Layer(data []byte, p gopacket.PacketBuilder) error {
	//fmt.Println("在这解析!!!")
	if len(data) < V2V1HeaderSize() {
		return errors.New("")
	}
	var layer V2V1Layer = V2V1Layer{
		Header:  data[0:V2V1HeaderSize()],
		Payload: data[V2V1HeaderSize():],
	}

	if layer.WhichPacketType() == UnknowPacket {
		return errors.New("")
	}
	p.AddLayer(&layer)
	return p.NextDecoder(V2V2LayerType)
}

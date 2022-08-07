package v2v

import (
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
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

func HeaderSize() int {
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
	if len(data) < HeaderSize() {
		return errors.New("V2V交换协议长度不足")
	}
	var layer V2V1Layer = V2V1Layer{
		Header:  data[0:HeaderSize()],
		Payload: data[HeaderSize():],
	}

	if layer.WhichPacketType() == UnknowPacket {

		return errors.New("未知的V2V交换协议")
	}
	p.AddLayer(&layer)
	return p.NextDecoder(layers.LayerTypeEthernet)
}

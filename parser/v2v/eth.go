package v2v

import (
	"github.com/google/gopacket"
)

type V2VEth struct {
	Header  []byte
	Payload []byte
}

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
	var layer V2V1Layer = V2V1Layer{
		Header:  data[0:14],
		Payload: data[14:],
	}

	p.AddLayer(&layer)
	return p.NextDecoder(V2V1LayerType)
}

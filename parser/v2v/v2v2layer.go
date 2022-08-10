package v2v

import (
	"encoding/binary"
	"errors"
	"github.com/google/gopacket"
)

type V2V2Layer struct {
	Header  []byte
	Payload []byte

	CmdId Cmd
}
type Cmd uint16

const (
	V2VUnknownCmd Cmd = 0x0000
	V2VLink       Cmd = 0x0001
	V2VLinkRes    Cmd = 0x1001
	V2VAuth       Cmd = 0x0002
	V2VAuthRes    Cmd = 0x1002
	V2VLogin      Cmd = 0x0003
	V2VLoginRes   Cmd = 0x1003
	V2VHeart      Cmd = 0x2001
	V2VHeartRes   Cmd = 0x3001
)

func (this *V2V2Layer) LayerType() gopacket.LayerType { return V2V2LayerType }
func (this *V2V2Layer) LayerContents() []byte         { return this.Header }
func (this *V2V2Layer) LayerPayload() []byte          { return this.Payload }

var (
	V2V2LayerType = gopacket.RegisterLayerType(1002, gopacket.LayerTypeMetadata{
		Name:    "V2V2LayerType",
		Decoder: gopacket.DecodeFunc(DecodeV2V2Layer),
	})
)

func DecodeV2V2Layer(data []byte, p gopacket.PacketBuilder) error {
	layer := newV2V2Layer(data)
	if layer == nil {
		return errors.New("")
	}
	p.AddLayer(layer)
	return nil
}

func getV2V2Cmd(data []byte) Cmd {
	cmdId := Cmd(binary.BigEndian.Uint16(data[0:2]))
	switch cmdId {
	case V2VLink:
	case V2VLinkRes:
	case V2VAuth:
	case V2VAuthRes:
	case V2VLogin:
	case V2VLoginRes:
	case V2VHeart:
	case V2VHeartRes:
	default:
		cmdId = V2VUnknownCmd
	}
	return cmdId
}

func newV2V2Layer(data []byte) *V2V2Layer {
	if len(data) < 2 {
		return nil
	}
	cmdId := getV2V2Cmd(data)
	if cmdId == V2VUnknownCmd {
		return nil
	}
	ret := &V2V2Layer{
		Header:  data[0:2],
		Payload: data[2:],
	}
	ret.CmdId = cmdId
	return ret
}

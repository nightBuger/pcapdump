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
	V2VUnknownCmd Cmd = 0x0000 // 未知包
	V2VLink       Cmd = 0x0001 // 设备连接 入网请求
	V2VLinkRes    Cmd = 0x1001 // 设备连接响应
	V2VAuth       Cmd = 0x0002 // 设备认证
	V2VAuthRes    Cmd = 0x1002 // 设备认证响应
	V2VLogin      Cmd = 0x0003 // 设备入网
	V2VLoginRes   Cmd = 0x1003 // 设备入网响应
	V2VHeart      Cmd = 0x2001 // 入网心跳
	V2VHeartRes   Cmd = 0x3001 // 入网心跳响应
)

func V2V2HeaderSize() uint32 {
	return 2
}

func (this *V2V2Layer) LayerType() gopacket.LayerType { return V2V2LayerType }
func (this *V2V2Layer) LayerContents() []byte         { return this.Header }
func (this *V2V2Layer) LayerPayload() []byte          { return this.Payload }

func (this *V2V2Layer) GenDecodeFunc(lt gopacket.LayerType) gopacket.DecodeFunc {
	return func(data []byte, p gopacket.PacketBuilder) error {
		layer := newV2V2Layer(data)
		if layer == nil {
			return errors.New("")
		}
		p.AddLayer(layer)
		if lt < 0 {
			return nil
		}
		return p.NextDecoder(lt)
	}
}

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
	if uint32(len(data)) < V2V2HeaderSize() {
		return nil
	}
	cmdId := getV2V2Cmd(data)
	if cmdId == V2VUnknownCmd {
		return nil
	}
	ret := &V2V2Layer{
		Header:  data[0:V2V2HeaderSize()],
		Payload: data[V2V2HeaderSize():],
	}
	ret.CmdId = cmdId
	return ret
}

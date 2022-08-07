package adapter

import "github.com/google/gopacket"

type Adapter interface {
	Parse(pack gopacket.Packet)
	Display()
	Filter() bool
}

package packetSv

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Packet struct {
	Time                   string
	Payload                string
	SrcMAC, DstMAC         string
	SrcAddress, DstAddress string
	Seq, Ack               uint32
	FIN, SYN, ACK          bool
}

type StartCaptureBo struct {
	PacketChan chan *Packet
	Duration   int64
	Device     string
	BPF        string
	Transfer   *Transfer
}

type Transfer struct {
	Flag             bool
	SrcMAC, DstMAC   string
	SrcIP, DstIP     string
	SrcPort, DstPort uint16
}

type Layers struct {
	Loopback *layers.Loopback
	Ethernet *layers.Ethernet
	IPv4     *layers.IPv4
	IPv6     *layers.IPv6
	TCP      *layers.TCP
	UDP      *layers.UDP
	Payload  *gopacket.Payload
}

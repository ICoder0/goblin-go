package networkLayer

import (
	"net"

	"github.com/google/gopacket/layers"
)

type Packet struct {
	Payload    []byte // 数据
	Version    uint8
	IHL        uint8
	TOS        uint8
	Length     uint16
	Id         uint16
	Flags      layers.IPv4Flag
	FragOffset uint16
	TTL        uint8
	Protocol   layers.IPProtocol
	Checksum   uint16
	SrcIP      net.IP
	DstIP      net.IP
	Options    []layers.IPv4Option
	Padding    []byte
}

func Keys() []string {
	return []string{"srcIP", "dstIP"}
}

func (p *Packet) TransferToArray() []string {
	// payload := string(p.Payload)
	srcIP := p.SrcIP.String()
	dstIP := p.DstIP.String()
	return []string{srcIP, dstIP}
}

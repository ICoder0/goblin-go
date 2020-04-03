package linkLayer

import (
	"net"

	"github.com/google/gopacket/layers"
)

type Packet struct {
	Payload        []byte // 数据
	SrcMAC, DstMAC net.HardwareAddr
	EthernetType   layers.EthernetType
	Length         uint16
}

func Keys() []string {
	return []string{"srcMac", "dstMac"}
}

func (p *Packet) TransferToArray() []string {
	// payload := string(p.Payload)
	srcMac := p.SrcMAC.String()
	dstMac := p.DstMAC.String()
	return []string{srcMac, dstMac}
}

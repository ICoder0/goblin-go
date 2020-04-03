package transportLayer

import (
	"strconv"

	"github.com/google/gopacket/layers"
)

type Packet struct {
	Payload                                    []byte // 数据
	SrcPort, DstPort                           uint16 // 源端口，目标端口
	Seq                                        uint32
	Ack                                        uint32
	DataOffset                                 uint8
	FIN, SYN, RST, PSH, ACK, URG, ECE, CWR, NS bool
	Window                                     uint16
	Checksum                                   uint16
	Urgent                                     uint16
	Options                                    []layers.TCPOption
	Padding                                    []byte
	opts                                       [4]layers.TCPOption
}

func Keys() []string {
	return []string{"srcPort", "dstPort", "payload", "seq", "ack", "fin", "syn"}
}

func (p *Packet) TransferToArray() []string {
	payload := string(p.Payload)
	srcPort := strconv.Itoa(int(p.SrcPort))
	dstPort := strconv.Itoa(int(p.DstPort))
	seq := strconv.Itoa(int(p.Seq))
	ack := strconv.Itoa(int(p.Ack))
	fin := strconv.FormatBool(p.FIN)
	syn := strconv.FormatBool(p.SYN)
	return []string{srcPort, dstPort, payload, seq, ack, fin, syn}
}

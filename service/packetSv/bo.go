package packetSv

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
}

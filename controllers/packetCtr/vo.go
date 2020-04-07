package packetCtr

type CaptureForm struct {
	Device     string `json:"device"`
	BPF        string `json:"bpf"`
	Duration   int64  `json:"duration"`
	IsTransfer bool   `json:"isTransfer"`
	DstIP      string `json:"dstIP"`
	DstPort    uint16 `json:"dstPort"`
}

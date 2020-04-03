package deviceCtr

import "net"

// 网卡设备
type Device struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Flags       uint32    `json:"flags"`
	Addresses   []Address `json:"addresses"`
}

// 网卡地址信息
type Address struct {
	IP        net.IP     `json:"ip"`
	Netmask   net.IPMask `json:"netmask"`
	Broadaddr net.IP     `json:"broadaddr"`
	P2P       net.IP     `json:"p2p"`
}

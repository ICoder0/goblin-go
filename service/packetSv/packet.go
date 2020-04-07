package packetSv

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	handle *pcap.Handle
)

// StartCapture 启动抓包
func StartCapture(ctx context.Context, bo *StartCaptureBo) (err error) {
	go func() {
		for {
			bo.Duration--
			if bo.Duration == 0 {
				stopCapture()
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	if handle, err = pcap.OpenLive(bo.Device, 65535, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err = handle.SetBPFFilter(bo.BPF); err != nil {
		stopCapture()
		panic(err)
	} else {
		fmt.Println("-------------------开始嗅探------------------")
		// 获取数据包管道
		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
		// 获取数据包
		go func() {
			for packet := range packetSource.Packets() {
				var (
					pkt      *Packet
					ethernet *layers.Ethernet
					loopback *layers.Loopback
					ipv4     *layers.IPv4
					ipv6     *layers.IPv6
					tcp      *layers.TCP
					udp      *layers.UDP
					payload  *gopacket.Payload
				)

				pkt = new(Packet)
				pkt.Time = time.Now().Format("2006-01-02 15:04:05")

				for _, l := range packet.Layers() {
					switch l.LayerType().String() {
					case "Ethernet":
						ethernet, _ = l.(*layers.Ethernet)
					case "Loopback":
						loopback, _ = l.(*layers.Loopback)
					case "IPv4":
						ipv4, _ = l.(*layers.IPv4)
					case "IPv6":
						ipv6, _ = l.(*layers.IPv6)
					case "TCP":
						tcp, _ = l.(*layers.TCP)
					case "UDP":
						udp, _ = l.(*layers.UDP)
					case "Payload":
						payload, _ = l.(*gopacket.Payload)
					}

				}

				lys := &Layers{
					Loopback: loopback,
					Ethernet: ethernet,
					IPv4:     ipv4,
					IPv6:     ipv6,
					TCP:      tcp,
					UDP:      udp,
					Payload:  payload,
				}

				// 转发
				if bo.Transfer.Flag {
					transfer(bo.Transfer, lys)
				}

				if ethernet != nil {
					pkt.SrcMAC = ethernet.SrcMAC.String()
					pkt.DstMAC = ethernet.DstMAC.String()
				}

				if ipv4 != nil {
					pkt.SrcAddress = ipv4.SrcIP.String()
					pkt.DstAddress = ipv4.DstIP.String()
				} else if ipv6 != nil {
					pkt.SrcAddress = ipv6.SrcIP.String()
					pkt.DstAddress = ipv6.DstIP.String()
				}

				if tcp != nil {
					tcp.Seq = 1
					pkt.Seq = tcp.Seq
					pkt.Ack = tcp.Ack
					pkt.FIN = tcp.FIN
					pkt.SYN = tcp.SYN
					pkt.ACK = tcp.ACK
					pkt.Payload = string(tcp.Payload)
					pkt.SrcAddress += ":" + tcp.SrcPort.String()
					pkt.DstAddress += ":" + tcp.DstPort.String()
				} else if udp != nil {

				}
				bo.PacketChan <- pkt
			}
		}()

	}

	return nil
}

// StopCapture 停止抓包
func stopCapture() {
	handle.Close()
	fmt.Println("-------------------停止嗅探------------------")
}

func transfer(tsf *Transfer, lys *Layers) {
	if lys.TCP.PSH {
		conn, err := net.Dial("tcp", tsf.DstIP+":"+strconv.Itoa(int(tsf.DstPort)))
		if err != nil {
			fmt.Println("创建连接失败：", err.Error())
		} else {
			_, err = conn.Write(lys.TCP.Payload)
			if err != nil {
				fmt.Println("转发失败：", err.Error())
			} else {
				fmt.Println("转发成功")
			}
			_ = conn.Close()
		}
	}
}

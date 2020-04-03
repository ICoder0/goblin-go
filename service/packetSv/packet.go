package packetSv

import (
	"context"
	"fmt"
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
				StopCapture()
				break
			}
			time.Sleep(1 * time.Second)
		}
	}()

	if handle, err = pcap.OpenLive(bo.Device, 1600, true, pcap.BlockForever); err != nil {
		panic(err)
	} else if err = handle.SetBPFFilter(bo.BPF); err != nil {
		StopCapture()
		panic(err)
	} else {

		fmt.Println("--------------------------------------------------------------------------")

		packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

		go func() {
			for packet := range packetSource.Packets() {
				var (
					pkt                                                 *Packet
					ethernet                                            *layers.Ethernet
					ip4Packet                                           *layers.IPv4
					ip6Packet                                           *layers.IPv6
					tcpPacket                                           *layers.TCP
					udpPacket                                           *layers.UDP
					linkLayertype, networkLayertype, transportLayertype gopacket.LayerType
				)

				pkt = new(Packet)
				pkt.Time = time.Now().Format("2006-01-02 15:04:05")

				if packet.LinkLayer() != nil {
					linkLayertype = packet.LinkLayer().LayerType()
					switch linkLayertype.String() {
					case "Ethernet":
						ethernet, _ = packet.Layer(linkLayertype).(*layers.Ethernet)
					}
				}

				if packet.NetworkLayer() != nil {
					networkLayertype = packet.NetworkLayer().LayerType()
					switch networkLayertype.String() {
					case "IPv4":
						ip4Packet, _ = packet.Layer(networkLayertype).(*layers.IPv4)
					case "IPv6":
						ip6Packet, _ = packet.Layer(networkLayertype).(*layers.IPv6)
					}
				}

				if packet.TransportLayer() != nil {
					transportLayertype = packet.TransportLayer().LayerType()
					switch transportLayertype.String() {
					case "TCP":
						tcpPacket, _ = packet.Layer(transportLayertype).(*layers.TCP)
					case "UDP":
						udpPacket, _ = packet.Layer(transportLayertype).(*layers.UDP)
					}
				}

				if ethernet != nil {
					pkt.SrcMAC = ethernet.SrcMAC.String()
					pkt.DstMAC = ethernet.DstMAC.String()
				}

				if ip4Packet != nil {
					pkt.SrcAddress = ip4Packet.SrcIP.String()
					pkt.DstAddress = ip4Packet.DstIP.String()
				} else if ip6Packet != nil {
					pkt.SrcAddress = ip6Packet.SrcIP.String()
					pkt.DstAddress = ip6Packet.DstIP.String()
				}

				if tcpPacket != nil {
					tcpPacket.Seq = 1
					pkt.Seq = tcpPacket.Seq
					pkt.Ack = tcpPacket.Ack
					pkt.FIN = tcpPacket.FIN
					pkt.SYN = tcpPacket.SYN
					pkt.ACK = tcpPacket.ACK
					pkt.Payload = string(tcpPacket.Payload)
					pkt.SrcAddress += ":" + tcpPacket.SrcPort.String()
					pkt.DstAddress += ":" + tcpPacket.DstPort.String()
				} else if udpPacket != nil {

				}

				bo.PacketChan <- pkt
			}
		}()
	}

	return nil
}

// StopCapture 停止抓包
func StopCapture() {
	handle.Close()
	fmt.Println("-------------------停止嗅探------------------")
}

// Forward 转发
func Forward(ctx context.Context) {

}

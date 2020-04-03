package deviceSv

import (
	"context"

	"github.com/google/gopacket/pcap"
)

func FindAllDevice(ctx context.Context) (devices []pcap.Interface, err error) {
	if devices, err = pcap.FindAllDevs(); err != nil {
		return []pcap.Interface{}, err
	}
	return devices, nil
}

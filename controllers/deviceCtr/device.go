package deviceCtr

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/gopacket/pcap"
	"goblin-go/controllers"
	"goblin-go/service/deviceSv"
)

func FindAllDevice(ctx *gin.Context) {

	var (
		err        error
		devices    []Device
		allDevices []pcap.Interface
	)

	if allDevices, err = deviceSv.FindAllDevice(ctx.Copy().Request.Context()); err != nil {
		// TODO: 2020/4/2 错误处理
		ctx.JSON(http.StatusOK, controllers.Response{
			Status: 400,
			Msg:    "获取设备列表失败",
		})
	}

	// 处理响应数据
	devices = []Device{}
	for _, d := range allDevices {
		var addresses []Address
		addresses = []Address{}
		for _, a := range d.Addresses {
			address := Address{
				IP:        a.IP,
				Netmask:   a.Netmask,
				Broadaddr: a.Broadaddr,
				P2P:       a.P2P,
			}
			addresses = append(addresses, address)
		}

		device := Device{
			Name:        d.Name,
			Description: d.Description,
			Flags:       d.Flags,
			Addresses:   addresses,
		}
		devices = append(devices, device)
	}

	ctx.JSON(http.StatusOK, controllers.Response{
		Status: 200,
		Msg:    "请求成功",
		Data:   devices,
	})
}

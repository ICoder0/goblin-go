package router

import (
	"github.com/gin-gonic/gin"
	"goblin-go/controllers/deviceCtr"
	"goblin-go/controllers/packetCtr"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	device := v1.Group("/device")
	device.GET("all", deviceCtr.FindAllDevice) // 获取可扫描到的所有网卡

	packet := v1.Group("/packet")
	packet.GET("/ws", packetCtr.WebSocketConnect) // 启动WebSocket
	packet.POST("/capture", packetCtr.Capture)    // 启动抓包

	_ = r.Run(":8085")
}

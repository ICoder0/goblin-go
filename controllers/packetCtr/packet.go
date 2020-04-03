package packetCtr

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"sniffer/controllers"
	"sniffer/service/packetSv"
)

var packetChan chan *packetSv.Packet

func InitPacketChan() {
	if packetChan == nil {
		packetChan = make(chan *packetSv.Packet, 1000)
	}
}

func Capture(ctx *gin.Context) {
	var (
		err  error
		form *CaptureForm
	)

	// 数据绑定
	form = new(CaptureForm)
	if err = ctx.ShouldBind(form); err != nil {
		fmt.Println(fmt.Sprintf("参数出错：%s", err.Error()))
		ctx.JSON(http.StatusOK, controllers.Response{
			Status: 400,
			Msg:    "参数出错",
		})
		return
	}

	InitPacketChan()
	startCaptureBo := &packetSv.StartCaptureBo{
		PacketChan: packetChan,
		Duration:   form.Duration,
		Device:     form.Device,
		BPF:        form.BPF,
	}
	if err = packetSv.StartCapture(ctx.Copy().Request.Context(), startCaptureBo); err != nil {
		ctx.JSON(http.StatusOK, controllers.Response{
			Status: 400,
			Msg:    "抓包失败:" + err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, controllers.Response{
		Status: 200,
		Msg:    "请求成功",
	})
}

func WebSocketConnect(ctx *gin.Context) {
	InitPacketChan()

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// 创建 websocket 连接
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, controllers.Response{
			Status: 400,
			Msg:    "请求失败:" + err.Error(),
		})
		return
	}
	defer conn.Close()

	fmt.Println("-------------------开启 WebSocket------------------")
	for {
		pkt, ok := <-packetChan
		if !ok {
			break
		}

		// 格式化数据
		p, _ := json.Marshal(pkt)
		if err = conn.WriteMessage(websocket.TextMessage, p); err != nil {
			log.Println(err)
			return
		}
	}
	close(packetChan)
}

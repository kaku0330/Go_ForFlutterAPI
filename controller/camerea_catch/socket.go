package camereacatch

import (
	"encoding/base64"
	"fmt"
	"go_api/util"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func GetPic(ctx *gin.Context) {
	ws, err := util.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ws.WriteJSON("fail to send data")
		return
	}
	defer ws.Close()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("error:", err)
			break
		}
		if err != nil {
			ws.WriteJSON("fail to send data")
		}
		imageData, err := base64.StdEncoding.DecodeString(string(msg))
		if err != nil {
			fmt.Println("decode base64:", err)
			continue
		}
		filePath := fmt.Sprintf("./doc/frame_%s.jpg", time.Now().Local().Format("2006_01_02_15_04_05"))

		err = os.WriteFile(filePath, imageData, 0644)
		if err != nil {
			fmt.Println("write file:", err)
			continue
		}
		fmt.Println("ok")
	}

	defer fmt.Println("Socket關閉")

}

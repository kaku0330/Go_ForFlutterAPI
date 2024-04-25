package users

import (
	conn "go_api/DB"
	handler "go_api/controller"

	"github.com/gin-gonic/gin"
)

type findaccount struct {
	Row bool `json:"row"`
}

func Get(ctx *gin.Context) {
	db := conn.ConnectDB()
	res := findaccount{}
	g := handler.GinContext{Ctx: ctx}
	var count int64
	result := db.Debug().Table("FlutterTestUse.flutter_useraccount").
		Where("username = ? AND password = ?", ctx.Query("username"), ctx.Query("password")).
		Count(&count)

	if result.Error != nil {
		g.SendResponse(500, "查詢資料失敗", nil)
		return
	}

	if count == 1 {
		res.Row = true
	} else {
		res.Row = false
	}

	g.SendResponse(200, "查詢資料成功", res)

}

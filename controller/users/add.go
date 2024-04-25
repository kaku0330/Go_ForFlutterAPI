package users

import (
	conn "go_api/DB"
	handler "go_api/controller"
	model "go_api/models"

	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) {
	db := conn.ConnectDB()
	res := model.User{}
	g := handler.GinContext{Ctx: ctx}
	ctx.ShouldBind(&res)

	result := db.Debug().Table("FlutterTestUse.flutter_useraccount").
		Create(&res)

	if result.Error != nil {
		g.SendResponse(500, "新增資料失敗", nil)
		return
	}

	g.SendResponse(200, "新增資料成功", res)

}

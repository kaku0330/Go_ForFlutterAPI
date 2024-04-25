package satisfactionsurvey

import (
	conn "go_api/DB"
	handler "go_api/controller"

	"go_api/models"

	"github.com/gin-gonic/gin"
)

func Post(ctx *gin.Context) {
	db := conn.ConnectDB()
	res := models.Servey{}

	g := handler.GinContext{Ctx: ctx}
	ctx.ShouldBind(&res)
	result := db.Debug().Table("FlutterTestUse.satisfaction_survey").
		Create(&res)

	if result.Error != nil {
		g.SendResponse(500, "新增資料失敗", nil)
		return
	}

	g.SendResponse(200, "新增資料成功", res)

}

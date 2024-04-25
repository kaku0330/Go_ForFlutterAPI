package users

import (
	conn "go_api/DB"
	handler "go_api/controller"
	model "go_api/models"

	"github.com/gin-gonic/gin"
)

func Update(ctx *gin.Context) {
	db := conn.ConnectDB()
	res := model.User{}
	g := handler.GinContext{Ctx: ctx}
	ctx.ShouldBind(&res)

	result := db.Debug().Table("FlutterTestUse.flutter_useraccount").
		Where("username = ?", res.Username).
		Updates(map[string]interface{}{
			"username": res.Username,
			"password": res.Password,
		})

	if result.RowsAffected == 0 {
		res.Username = ""
		res.Password = ""
		g.SendResponse(200, "查無資料", res)
		return
	}

	if result.Error != nil {
		g.SendResponse(500, "查詢資料失敗", nil)
		return
	}

	result = db.Debug().Table("FlutterTestUse.flutter_useraccount").
		Where("username = ?", res.Username).
		Updates(map[string]interface{}{
			"username": res.Username,
			"password": res.Password,
		})

	if result.Error != nil {
		g.SendResponse(500, "資料更改失敗", nil)
		return
	}

	g.SendResponse(200, "資料更改成功", res)

}

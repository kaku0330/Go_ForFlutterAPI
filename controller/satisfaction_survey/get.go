package satisfactionsurvey

import (
	conn "go_api/DB"
	handler "go_api/controller"
	"go_api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DataSetCount struct {
	Servey []models.Servey `json:"servey"`
	Count  int64           `json:"count"`
}

var result *gorm.DB

func Get(ctx *gin.Context) {
	db := conn.ConnectDB()
	data := []models.Servey{}
	res := DataSetCount{}
	g := handler.GinContext{Ctx: ctx}

	result = db.Debug().Table("FlutterTestUse.satisfaction_survey").
		Find(&data)

	if result.Error != nil {
		g.SendResponse(500, "查詢資料失敗", nil)
		return
	}
	res.Servey = data
	g.SendResponse(200, "查詢資料成功", res)

}

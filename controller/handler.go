package handler

import "github.com/gin-gonic/gin"

type GinContext struct {
	Ctx *gin.Context
}

type response struct {
	Status uint16      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func (g *GinContext) SendResponse(status uint16, msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Status: status,
		Msg:    msg,
		Data:   data,
	})
}

package ceshi

import (
	"github.com/gin-gonic/gin"
	//"web-gin/amqp/producer"
	"net/http"
)

func SendMq(ctx *gin.Context)  {
	//goods_id := ctx.DefaultQuery("goods_id", "0")

	//producer.TestGoods(goods_id)

	ctx.JSON(http.StatusOK, nil)
}

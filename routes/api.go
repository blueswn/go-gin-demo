package routes

import (
	"github.com/gin-gonic/gin"
	"log"
)


func LoadApiRouter(e *gin.Engine)  {
	e.GET("/v1", func(context *gin.Context) {
		log.Panic("13213")
	})
}

package api

import (
	"artisan/app/handler"
	"github.com/gin-gonic/gin"
)

func Ping(r gin.IRouter) {
	r.GET("/ping", handler.Ping)
}

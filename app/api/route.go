package api

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func InitHttp() {
	r := gin.Default()
	gin.SetMode(viper.GetString("http.mode"))
	RouteTable(r)
	err := r.Run(":" + viper.GetString("http.port")) // listen and serve on 0.0.0.0:8080
	if err != nil {
		panic(err)
	}
}

func RouteTable(r gin.IRouter) {
	Ping(r)
}

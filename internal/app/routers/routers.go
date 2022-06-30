package routers

import (
	"github.com/gin-gonic/gin"

	"memorable/internal/app/services"

	"net/http"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("web/templete/*")

	v1 := r.Group("/forever")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})

		v1.GET("/520cjx", services.DrawHeart)
	}

	return r
}

package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DrawHeart(c *gin.Context)  {
	c.HTML(http.StatusOK,"beatingheart.html",nil)
}

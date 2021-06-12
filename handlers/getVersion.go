package handlers

import (
	"net/http"

	"git.target.com/DhimanGhosh/GoPrograming/vget/services"
	"github.com/gin-gonic/gin"
)

func GetVersionHandler(c *gin.Context) {
	s, err := services.GetVersion("1")
	if err != nil {
		c.String(http.StatusOK, err.Error())
	}

	c.String(http.StatusOK, *s)
}

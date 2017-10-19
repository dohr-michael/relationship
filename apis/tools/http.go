package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func OptionHandler(c *gin.Context) {
	c.Status(http.StatusOK)
}

func NoCacheMiddleware(c *gin.Context) {
	c.Header("Cache-Control", "no-store, must-revalidate, no-cache, max-age=0")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "Mon, 01 Jan 0001 00:00:00 GMT")
	c.Next()
}

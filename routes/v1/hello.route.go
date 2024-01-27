// Package v1 hello.route.go
package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var response = gin.H{"message": "Hello world", "status": http.StatusOK}

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, response)
}

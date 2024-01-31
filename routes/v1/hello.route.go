// Package v1 hello.route.go
package v1

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

var status = http.StatusOK
var response = gin.H{"response": "Hello world", "status": status}

// HelloHandler
// ROUTE: /v1/hello
// Jar health check
func HelloHandler(c *gin.Context) {
    c.JSON(status, response)
}

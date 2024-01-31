// main.go
package main

import (
    "github.com/gin-gonic/gin"
    "jstnf/the-jar-api/data"
    "jstnf/the-jar-api/routes/v1"
    "log"
    "os"
)

func main() {
    data.Init()

    router := gin.Default()
    router.GET("/v1/hello", v1.HelloHandler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "3001"
    }

    log.Printf("Server is running on port %s", port)
    log.Fatal(router.Run(":" + port))
}

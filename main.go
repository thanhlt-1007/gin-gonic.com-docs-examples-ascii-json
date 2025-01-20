package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/someJSON", func(context *gin.Context) {
        data := map[string]interface{} {
            "lang": "GO语言",
            "tag":  "<br>",
        }
        context.AsciiJSON(
            http.StatusOK,
            data,
        )
    })

    router.Run(":8080")
}

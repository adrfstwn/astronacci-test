package router

import (
    "github.com/gin-gonic/gin"
    "backend/internal/handler"
)

func Setup(h *handler.VoucherHandler) *gin.Engine {
    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "POST, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    })

    api := r.Group("/api")
    {
        api.POST("/check", h.Check)
        api.POST("/generate", h.Generate)
    }

    return r
}
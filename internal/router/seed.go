package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/sabrinababakulova/golang-crud/internal/database"
)

func Seed(r *gin.Engine) {
    r.GET("/seed", func(c *gin.Context) {
        database.Migrate()
        c.JSON(http.StatusOK, gin.H{"message": "Seeded successfully!"})
    })
}

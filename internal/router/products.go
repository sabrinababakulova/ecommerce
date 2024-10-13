package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/sabrinababakulova/golang-crud/internal/database"
    "github.com/sabrinababakulova/golang-crud/internal/types"

)

func SetupRoutes(r *gin.Engine) {
    r.GET("/test", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "testing successful!"})
    })

    r.GET("/products", func(c *gin.Context) {
      products := database.GetProducts()
      c.JSON(http.StatusOK, gin.H{"data": products})
    })

    r.POST("/products", func(c *gin.Context) {
      var product types.ProductJson
        
        if err := c.ShouldBindJSON(&product); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        result, error := database.CreateProduct(product)
        if error != nil {
          panic(error)
        }
        c.JSON(http.StatusOK, gin.H{"message": "Product created successfully!", "affected": result})
    })

    r.DELETE("/products/:id", func(c *gin.Context) {
        id := c.Param("id")

        result, error := database.DeleteProduct(id)

        if error != nil {
          panic(error)
        }

        c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully!", "affected": result})
    })
}

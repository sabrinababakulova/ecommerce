package main

import (
    "github.com/gin-gonic/gin"
    "github.com/sabrinababakulova/golang-crud/internal/router"
    "github.com/sabrinababakulova/golang-crud/internal/database"
)
func main() {
  database.ConnectDatabase()
  r := gin.Default()
  router.SetupRoutes(r)
  router.Seed(r)
  r.Run(":8080")
  defer database.Db.Close()
}


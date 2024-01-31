package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"gopkg.in/mgo.v2/bson"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.Default()
	// router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, bson.M{"Application": "Lifecycle test", "Status": "Up"})
	})
	router.Run(":" + port)
}
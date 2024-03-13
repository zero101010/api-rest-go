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
		port = "8080"
	}
	router := gin.Default()
	// router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, bson.M{"Application": " Go test", "Status": "Up", "version": "3"})
	})
	router.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, bson.M{"test": "Running"})
	})
	router.Run(":" + port)
}
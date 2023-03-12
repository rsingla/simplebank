package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World, Yohaan is fuddu!")
	})

	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World, Rajeev is awesome!")
	})

	r.Run()
}

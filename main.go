package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gvquiroz/oauth-challenge/countutils"
	"github.com/djimenez/iconv-go"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	// healthcheck
	r.GET("/ping", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, "pong")
	})
	r.GET("/", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(200, "")
	})

	r.GET("/count", func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")

		input := c.Query("input")
		output,_ := iconv.ConvertString(input, "latin1", "utf-8")
		result := countutils.CountDuplicatedChars(output)

		c.JSON(200, result)
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":4567")
}
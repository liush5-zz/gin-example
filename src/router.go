package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func startGin() {
	r := gin.Default()

	RouteRegister(r)

	r.Run(":8080")
}

func RouteRegister(r *gin.Engine) {
	r.GET("/ping", otherMw(), loginTest)

	admin := r.Group("/liu")
	admin.Use(LoggerMw())
	{
		admin.GET("/login", otherMw(), loginTest)
	}

}

/////////////////

func loginTest(c *gin.Context) {

	c.String(200, " loginTest:pong")
}

func LoggerMw() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		c.String(200, " LoggerMw: hhh")
		log.Println(status)

		c.Next()
	}
}

func otherMw() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// after request
		latency := time.Since(t)
		log.Print(latency)

		c.String(200, " otherMw: ping")

		c.Next()
	}
}

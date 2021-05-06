package main

import (
	"log"
	"net/http"
	"time"

	"github.com/huruji/eva/eva"
)

func onlyForV2() eva.HandlerFunc {
	return func(c *eva.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := eva.New()
	r.Use(eva.Logger()) // global midlleware
	r.GET("/", func(c *eva.Context) {
		c.HTML(http.StatusOK, "<h1>Hello eva</h1>")
	})
	r.Static("/assets", "./static")

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *eva.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}

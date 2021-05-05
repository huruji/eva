package main

import (
	"net/http"

	"github.com/huruji/eva/eva"
)

func main() {
	r := eva.New()
	r.GET("/", func(c *eva.Context) {
		c.HTML(http.StatusOK, "<h1>hello</h1>")
	})
	r.GET("/hello", func(c *eva.Context) {
		c.String(http.StatusOK, "hello %s, you are at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/user", func(c *eva.Context) {
		c.JSON(http.StatusOK, eva.H{
			"username": "huruji",
		})
	})
	r.Run(":9999")
}

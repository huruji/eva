package main

import (
	"net/http"

	"github.com/huruji/eva/eva"
)

func main() {
	r := eva.New()
	r.GET("/", func(c *eva.Context) {
		c.HTML(http.StatusOK, "<h1>Hello eva</h1>")
	})

	r.GET("/hello", func(c *eva.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *eva.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *eva.Context) {
		c.JSON(http.StatusOK, eva.H{"filepath": c.Param("filepath")})
	})
	r.Run(":9999")
}

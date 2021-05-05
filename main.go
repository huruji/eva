package main

import (
	"net/http"

	"github.com/huruji/eva/eva"
)

func main() {
	r := eva.New()
	r.GET("/", func(c *eva.Context) {
		c.HTML(http.StatusOK, "<h1>index page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *eva.Context) {
			c.HTML(http.StatusOK, "<h1>hello v1</h1>")
		})
		v1.GET("/hello", func(c *eva.Context) {
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *eva.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *eva.Context) {
			c.JSON(http.StatusOK, eva.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":9999")
}

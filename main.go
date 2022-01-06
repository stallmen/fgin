package main

import (
	"fmt"
	"gee/fastest"
	"net/http"
	"reflect"
)

func main() {
	f := fastest.New()
	f.GET("/", func(c *fastest.Context) {
		c.String(http.StatusOK,"hehe")
	})

	f.GET("/xx", func(c *fastest.Context) {
		x := c.Get("id")
		fmt.Println(x,reflect.TypeOf(x))
		c.Json(200, map[string]string{
			"id":"123",
			"name":"lisi",
		})
	})

	f.POST("/hehe", func(c *fastest.Context) {
		x := c.Post("name")
		fmt.Println(x,reflect.TypeOf(x))
	})


	f.RUN(":8888")
}

func test(args ...string)  {
	fmt.Println(args)
}
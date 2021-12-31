package main

import (
	"fmt"
	"gee/lee"
	"net/http"
)

func main() {
	f := lee.Ins()
	f.GET("/", func(w http.ResponseWriter, r *http.Request) {
		test("1","2","3")
		fmt.Println(r.URL.Query())
	})

	f.RUN(":8888")
}

func test(args ...string)  {
	fmt.Println(args)
}
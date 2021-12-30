package main

import (
	"gee/fgin"
	"net/http"
)

func main() {
	f := fgin.Ins()
	f.GET("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("nihao"))
	})

	f.RUN(":8888")
}

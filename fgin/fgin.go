package fgin

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	_ = iota
	ONE
)

type HandleFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandleFunc
}

func Ins() *Engine {
	return &Engine{make(map[string]HandleFunc)}
}

//路由添加
func (e *Engine) addRoute(method, route string, handle HandleFunc) {
	var str strings.Builder
	str.Grow(len(method) + len(route) + ONE)

	_, err := str.WriteString(method)
	if err != nil {
		log.Fatal("路由注册错误")
	}
	_, err = str.WriteString("-")
	if err != nil {
		log.Fatal("路由注册错误")
	}

	_, err = str.WriteString(route)
	if err != nil {
		log.Fatal("路由注册错误")
	}

	e.router[str.String()] = handle
}

func (e *Engine) GET(route string, handle HandleFunc) {
	e.addRoute("GET", route, handle)
}

func (e *Engine) POST(route string, handle HandleFunc) {
	e.addRoute("POST", route, handle)
}

//启动服务
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var key strings.Builder
	key.Grow(len(r.Method) + len(r.URL.Path) + ONE)

	_, err := key.WriteString(r.Method)
	if err != nil {
		log.Fatal("路由注册错误")
	}
	_, err = key.WriteString("-")
	if err != nil {
		log.Fatal("路由注册错误")
	}
	_, err = key.WriteString(r.URL.Path)
	if err != nil {
		log.Fatal("路由注册错误")
	}

	if handle, ok := e.router[key.String()]; ok {
		handle(w, r)
	} else {
		fmt.Fprintf(w, "PATH %s 404 NOT FOUND", key.String())
	}
}

func (e *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

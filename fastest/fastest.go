package fastest

import (
	"net/http"
)

const (
	_ = iota
	One
)

type HandleFunc func(ctx *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method, route string, handle HandleFunc)  {
	e.router.addRoute(method,route,handle)
}


func (e *Engine) GET(route string, handle HandleFunc) {
	e.addRoute("GET", route, handle)
}

func (e *Engine) POST(route string, handle HandleFunc) {
	e.addRoute("POST", route, handle)
}

//接管实现
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := NewContext(w,r)
	e.router.handle(c)
}

// RUN 启动服务
func (e *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

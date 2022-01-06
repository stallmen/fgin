package fastest

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type router struct {
	handler map[string]HandleFunc
}

func newRouter() *router {
	return &router{
		handler: make(map[string]HandleFunc),
	}
}

//添加路由
func (r *router) addRoute(method, route string, handle HandleFunc)  {
	var str strings.Builder
	str.Grow(len(method) + len(route) + One)

	_, err := str.WriteString(method)
	if err != nil {
		log.Fatal("addRoute fail")
	}
	_, err = str.WriteString("-")
	if err != nil {
		log.Fatal("addRoute fail")
	}

	_, err = str.WriteString(route)
	if err != nil {
		log.Fatal("addRoute fail")
	}

	r.handler[str.String()] = handle
}

//请求接管
func (r *router) handle(c *Context)  {
	var key strings.Builder
	key.Grow(len(c.Method) + len(c.Request.URL.Path) + One)

	_, err := key.WriteString(c.Method)
	if err != nil {
		log.Fatal("route read fail")
	}
	_, err = key.WriteString("-")
	if err != nil {
		log.Fatal("route read fail")
	}
	_, err = key.WriteString(c.Request.URL.Path)
	if err != nil {
		log.Fatal("route read fail")
	}

	if handler, ok := r.handler[key.String()]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound,fmt.Sprintf( "PATH %s 404 NOT FOUND", key.String()))
	}
}










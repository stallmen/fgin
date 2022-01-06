package fastest

import (
	"encoding/json"
	"net/http"
	"strings"
)


type Context struct {
	Response http.ResponseWriter
	Request *http.Request
	Path string
	Method string
	StatusCode int
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context  {
	return &Context{
		Response: w,
		Request: r,
		Path: r.URL.Path,
		Method: r.Method,
	}
}


// Get GET参数获取,支持2个参数,第二个为默认值(任意类型)
func (c *Context) Get(arg string) string {
	return strings.TrimSpace(c.Request.URL.Query().Get(arg))
}

// Post post参数获取
func (c *Context) Post(arg string) string {
	return strings.TrimSpace(c.Request.PostFormValue(arg))
}

// SetHeader 请求头设置
func (c *Context) SetHeader(key,value string)  {
	c.Response.Header().Set(key,value)
}

// Status 状态码设置
func (c *Context) Status(code int)  {
	c.StatusCode = code
	c.Response.WriteHeader(code)
}


//字符串返回
func (c *Context) String(code int,str string)  {
	c.SetHeader("Content-Type","text/plain")
	c.Status(code)
	if _,err := c.Response.Write([]byte(str));err != nil {
		http.Error(c.Response,err.Error(),http.StatusInternalServerError)
	}
}

// Json json返回
func (c *Context) Json(code int,v interface{})  {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)

	if err := json.NewEncoder(c.Response).Encode(v);err != nil {
		http.Error(c.Response,err.Error(),http.StatusInternalServerError)
	}
}




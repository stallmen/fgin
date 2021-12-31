package lee

import "net/http"


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
func (c *Context) Get(args ...interface{}) interface{} {
	lenth := len(args)

	if lenth == Zero {
		return nil
	} else {
		arg,ok := args[0].(string)

		if !ok {
			return nil
		}

		if lenth == One {
			return c.Request.URL.Query().Get(arg)
		} else if lenth == Two {
			if c.Request.URL.Query().Get(arg) == EmptyString {
				return args[1]
			}

			return c.Request.URL.Query().Get(arg)
		}

		return nil
	}
}





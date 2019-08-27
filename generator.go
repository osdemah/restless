package restless

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
)

var httpRespType = reflect.TypeOf(new(HttpResponse)).Elem()

func generateHandler(v reflect.Value) (func(*gin.Context), error) {
	m := v.Type()
	if m.NumIn() != 0 {
		return nil, errors.New("Input paramters are not supported yet!")
	}

	if m.NumOut() > 1 {
		return nil, errors.New("Only one http response is supported!")
	}

	if m.NumOut() == 1 && !m.Out(0).Implements(httpRespType) {
		return nil, errors.New("Handler response should implement HttpResponse interface!")
	}

	return func(c *gin.Context) {
		out := v.Call([]reflect.Value{})
		if len(out) == 1 {
			code := int(out[0].MethodByName("Code").Call([]reflect.Value{})[0].Int())
			resp := out[0].MethodByName("Response").Call([]reflect.Value{})[0].Interface()
			c.JSON(code, resp)
		} else {
			c.Status(http.StatusNoContent)
		}
	}, nil
}

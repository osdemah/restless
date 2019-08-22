package restless

import (
	"errors"
	"fmt"
	"net/http"
	"path"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

type HttpResponse interface {
	Code() int
	Response() interface{}
}

func CompileAndRun(h interface{}) error {
	httpRespType := reflect.TypeOf(new(HttpResponse)).Elem()

	router := gin.Default()

	t := reflect.TypeOf(h)
	v := reflect.ValueOf(h)
	// Kind returns the specific kind of this type.
	if t.Kind() != reflect.Struct {
		return errors.New("Can not parse non-struct type handlers")
	}

	pathField, has := t.FieldByName(StructHttpPrefixVariableName)
	prefix := ""
	if has {
		if pathField.Type.Name() != "string" {
			return fmt.Errorf("%s must be string", StructHttpPrefixVariableName)
		}
		prefix = v.FieldByName(StructHttpPrefixVariableName).String()
	}

	for i := 0; i < t.NumMethod(); i++ {
		method := t.Method(i)
		name := method.Name
		if !strings.HasPrefix(name, HttpHandlerMethodNamePrefix) {
			continue
		}

		if method.Type.NumIn() != 1 {
			return errors.New("Input paramters are not supported yet!")
		}

		httpMethod, httpPath, err := extractMethod(strings.TrimPrefix(name, HttpHandlerMethodNamePrefix))

		fullPath := path.Join("/", prefix, httpPath)
		err = nil
		router.Handle(strings.ToUpper(httpMethod), fullPath, func(c *gin.Context) {
			out := v.MethodByName(name).Call([]reflect.Value{})
			if len(out) > 1 {
				err = errors.New("Only one http response is supported!")
				return
			}

			if len(out) == 1 {
				if !out[0].Type().Implements(httpRespType) {
					err = errors.New("Handler response should implement HttpResponse interface")
					return
				}

				code := int(out[0].MethodByName("Code").Call([]reflect.Value{})[0].Int())
				resp := out[0].MethodByName("Response").Call([]reflect.Value{})[0].Interface()
				c.JSON(code, resp)
			} else {
				c.Status(http.StatusNoContent)
			}
		})

		if err != nil {
			return err
		}
	}
	return router.Run(":8080")
}

package restless

import (
	"errors"
	"fmt"
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

		hdl, err := generateHandler(v.MethodByName(name))

		if err != nil {
			return err
		}

		httpMethod, httpPath, err := extractMethod(strings.TrimPrefix(name, HttpHandlerMethodNamePrefix))

		fullPath := path.Join("/", prefix, httpPath)
		err = nil
		router.Handle(strings.ToUpper(httpMethod), fullPath, hdl)
	}
	return router.Run(":8080")
}

package restless

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
)

type HttpResponse interface {
	Code() int
	Response() interface{}
}

func CompileAndRun(h interface{}) error {
	router := gin.Default()

	v := reflect.ValueOf(h)
	// Kind returns the specific kind of this type.
	if v.Kind() != reflect.Struct {
		return errors.New("Can not parse non-struct type handlers")
	}

	prefix, err := memeberToString(v, StructHttpPrefixVariableName)
	if err != nil {
		return err
	}

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		name := v.Type().Method(i).Name

		httpMethod, httpPath, err := extractHttpMethodPathWithPrefix(name, prefix)
		if err != nil {
			return err
		}
		if len(httpMethod) == 0 && len(httpPath) == 0 {
			// Non-Handler functions will be ignored!
			continue
		}

		hdl, err := generateHandler(method)
		if err != nil {
			return err
		}

		router.Handle(httpMethod, httpPath, hdl)
	}
	return router.Run(":8080")
}

package restless

import (
	"errors"
	"fmt"
	"path"
	"reflect"
	"strings"
)

func toCamel(str string) string {
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}

func extractHttpMethodPath(funcName string) (httpMethod, httpPath string, err error) {
	for _, v := range httpMethods {
		if strings.HasPrefix(funcName, v) {
			httpURL := strings.TrimPrefix(funcName, v)
			elements := camelExtract.FindAllString(httpURL, -1)
			if len(elements) == 0 && len(httpURL) > 0 {
				return "", "", errors.New("Malformed function name, The Function name should comply with HttpGetCamelCase format")
			}
			for i := 0; i < len(elements); i++ {
				elements[i] = strings.ToLower(elements[i])
			}
			return strings.ToUpper(v), path.Join(elements...), nil
		}
	}
	return "", "", errors.New("Malformed function name, A HTTP Method name must come after Http in function name")
}

func memeberToString(v reflect.Value, f string) (string, error) {
	v = v.FieldByName(f)
	if !v.IsValid() {
		return "", nil
	}

	if v.Kind() != reflect.String {
		return "", fmt.Errorf("%s must be string", v.Kind())
	}

	return v.String(), nil
}

func extractHttpMethodPathWithPrefix(funcName, prefix string) (httpMethod, httpPath string, err error) {
	if !strings.HasPrefix(funcName, HttpHandlerMethodNamePrefix) {
		// This is not malformed functionm, it's just not supposed to be a handler
		return "", "", nil
	}

	httpMethod, httpPath, err = extractHttpMethodPath(strings.TrimPrefix(funcName, HttpHandlerMethodNamePrefix))
	if err != nil {
		return "", "", err
	}

	return strings.ToUpper(httpMethod), path.Join("/", prefix, httpPath), nil
}

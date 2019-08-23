package restless

import (
	"errors"
	"path"
	"strings"
)

func toCamel(str string) string {
	return strings.ToUpper(str[:1]) + strings.ToLower(str[1:])
}

func extractMethod(funcName string) (method, url string, err error) {
	for _, v := range methods {
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

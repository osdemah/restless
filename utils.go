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
			elements := camelExtract.FindAllString(strings.TrimPrefix(funcName, v), -1)
			for i := 0; i < len(elements); i++ {
				elements[i] = strings.ToLower(elements[i])
			}
			return strings.ToUpper(v), path.Join(elements...), nil
		}
	}
	return "", "", errors.New("HTTP Method name must comes after Http in function name")
}

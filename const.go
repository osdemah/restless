package restless

import (
	"net/http"
	"regexp"
)

var camelExtract = regexp.MustCompile(`[A-Z][a-z0-9]*`)
var httpMethods = [9]string{
	toCamel(http.MethodGet),
	toCamel(http.MethodPost),
	toCamel(http.MethodDelete),
	toCamel(http.MethodPatch),
	toCamel(http.MethodPut),
	toCamel(http.MethodHead),
	toCamel(http.MethodConnect),
	toCamel(http.MethodOptions),
	toCamel(http.MethodTrace),
}

const (
	StructHttpPrefixVariableName = "Prefix"
	HttpHandlerMethodNamePrefix  = "Http"
)

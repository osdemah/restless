package restless

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToCamelLowerCase(t *testing.T) {
	assert.Equal(t, "Abcd", toCamel("abcd"), "Lower case word should be converted to a Camel word!")
}

func TestToCamelUpperCase(t *testing.T) {
	assert.Equal(t, "Abcd", toCamel("ABCD"), "Upper case word should be converted to a Camel word!")
}

func TestToCamelMixedCase(t *testing.T) {
	assert.Equal(t, "Abcd", toCamel("aBcD"), "Mixed case word should be converted to a Camel world!")
}

func TestToCamelCamelCase(t *testing.T) {
	assert.Equal(t, "Abcd", toCamel("Abcd"), "Camel word shouldn't be changed!")
}

func TestExtractMethodBaisc(t *testing.T) {
	method, url, err := extractHttpMethodPath("GetBasic")
	assert.NoError(t, err, "GetBasic is a valid method name")
	assert.Equal(t, "GET", method, "Method should be extracted correctly!")
	assert.Equal(t, "basic", url, "URL should be extracted correctly!")
}

func TestExtractMethodTwoWords(t *testing.T) {
	method, url, err := extractHttpMethodPath("PostHelloWorld")
	assert.NoError(t, err, "PostHelloWorld is a valid method name")
	assert.Equal(t, "POST", method, "Method should be extracted correctly!")
	assert.Equal(t, "hello/world", url, "URL should be extracted correctly!")
}

func TestExtractMethodMultipleWords(t *testing.T) {
	method, url, err := extractHttpMethodPath("PatchHelloWorldBasicTest")
	assert.NoError(t, err, "PatchHelloWorldBasicTest is a valid method name")
	assert.Equal(t, "PATCH", method, "Method should be extracted correctly!")
	assert.Equal(t, "hello/world/basic/test", url, "URL should be extracted correctly")
}

func TestExtractMethodAllHttpMethods(t *testing.T) {
	methods := []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}
	for _, m := range methods {
		c := toCamel(m) + "HelloWorldBasicTest"
		method, url, err := extractHttpMethodPath(c)
		assert.NoError(t, err, c+" is a valid method name")
		assert.Equal(t, m, method, "Method should be extarcted correctly!")
		assert.Equal(t, "hello/world/basic/test", url, "URL should be extracted correctly!")
	}
}

func TestExtractMethodMalformedMethod(t *testing.T) {
	method, url, err := extractHttpMethodPath("MalformedHelloWorldBasicTest")
	assert.EqualError(t, err, "Malformed function name, A HTTP Method name must come after Http in function name", "Expected Error!")
	assert.Equal(t, method, "", "It shouldn't return any method for the Malformed function name!")
	assert.Equal(t, url, "", "It shouldn't return any URL for the Malformed function name!")
}

func TestExtractMethodMalformedURL(t *testing.T) {
	method, url, err := extractHttpMethodPath("Deleteabcd")
	assert.EqualError(t, err, "Malformed function name, The Function name should comply with HttpGetCamelCase format", "Expected Error!")
	assert.Equal(t, method, "", "It shouldn't return any method for the Malformed function name!")
	assert.Equal(t, url, "", "It shouldn't return any URL for the Malformed function name!")
}

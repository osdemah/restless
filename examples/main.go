package main

import (
	"github.com/osdemah/restless"
	"log"
)

type BasicResponse struct {
}

func (b BasicResponse) Code() int {
	return 200
}

func (b BasicResponse) Response() interface{} {
	return "Hello World!"
}

type BasicHandler struct {
}

func (b BasicHandler) HttpGetHello() BasicResponse {
	return BasicResponse{}
}

func main() {
	log.Fatal(restless.CompileAndRun(BasicHandler{}))
}

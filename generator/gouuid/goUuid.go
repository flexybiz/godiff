package main

import (
	"C"
	"fmt"
	"github.com/satori/go.uuid"
)

//export GoUuid
func GoUuid() *C.char {
	return C.CString(fmt.Sprintf("%s", uuid.NewV4()))
}

func main() {}

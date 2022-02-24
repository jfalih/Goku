package main

import (
	"fmt"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	fmt.Println("Hello World")
}

package main

import (
	"Otra"
	"fmt"
)

func main() {

	fmt.Println("hello from main.go")
	something := Otra.OtraFunction("wer")
	fmt.Println(something)
}

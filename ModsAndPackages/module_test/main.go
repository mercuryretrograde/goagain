package main

import (
	"example.com/module_test/package_00"
	"fmt"
)

var filePath = "/home/che_deane/gopath/goagain/ModsAndPackages/module_test/main.go"

func main() {
	emp := package_00.Struct00{Name: "Sam", Age: 31, Salary: 2000}

	fmt.Printf("\nfile path = '%s'\n\n", filePath)
	fmt.Printf("\nemp struct Name = '%s'\n\n", emp.Name)

}

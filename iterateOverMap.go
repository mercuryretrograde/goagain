package main

import (
	"fmt"
)

func getAllKeys(sample map[string]string) []string {
	var keys []string
	for k := range sample {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	sample := map[string]string{
		"a": "w",
		"b": "x",
		"c": "y",
		"d": "z",
	}

	fmt.Println("Iterating over all keys and values")

	for k, v := range sample {
		fmt.Printf("key: %s value: %s\n", k, v)
	}

	fmt.Println("Iterating over only keys")

	for k := range sample {
		fmt.Printf("key: %s\n", k)
	}

	fmt.Println("Iterating over only values")

	for _, v := range sample {
		fmt.Printf("value: %s\n", v)
	}

	fmt.Println("Get list of all keys")
	keys := getAllKeys(sample)
	fmt.Println(keys)

}

package main

import "fmt"

func main() {
	nico := map[string]string{"name": "nico", "age": "12"}
	for key, _ := range nico {
		fmt.Println(key)
	}
}

// map 은 key: value가 있다
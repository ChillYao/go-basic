package main

import (
	"fmt"
)

func main() {
	animals := []string{"dog", "cat"}
	for i := 0; i < len(animals); i++ {
		fmt.Println(animals[i])
	}

	for index, value := range animals {
		fmt.Println(index, value)
	}
}

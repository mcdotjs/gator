package main

import (
	"fmt"
	"github.com/mcdotjs/blog_aggregator/internal/config"
)

func main() {
	fileContent, err := config.Read()
	if err != nil {
		fmt.Errorf("Problem with reading file")
	}
	fmt.Println("first read", fileContent)
	fileContent.SetUser("Mirko")

	fileContent, err = config.Read()
	if err != nil {
		fmt.Errorf("Problem with reading file")
	}
	fmt.Println("second read", fileContent)

}

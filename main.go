package main

import (
	"fmt"

	"github.com/ozlevka/devoptools/config"
)

func main() {
	fmt.Println("Start")
	conf:= config.Read()

	fmt.Printf("This is config %v\n", conf)
}

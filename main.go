package main

import (
	"log"
	"os"

	"github.com/sganon/computorV1/parser"
)

func main() {
	err := parser.ParseEquation(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
}

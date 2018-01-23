package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sganon/computorV1/parser"
	"github.com/sganon/computorV1/reducer"
)

func main() {
	equation, err := parser.ParseEquation(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(equation)
	reducer.ReduceEquation(&equation)
	fmt.Println(equation)
	err = reducer.ValidateEquation(equation)
	if err != nil {
		log.Fatal(err)
	}
}

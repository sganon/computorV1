package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sganon/computorV1/parser"
	"github.com/sganon/computorV1/reducer"
	"github.com/sganon/computorV1/resolver"
)

func main() {
	equation, err := parser.ParseEquation(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	reducer.ReduceEquation(&equation)
	fmt.Printf("Reduced :")
	reducer.PrintEquation(equation)
	err = reducer.ValidateEquation(&equation)
	if err != nil {
		log.Fatal(err)
	}
	resolver.ResolveEquation(equation)
}

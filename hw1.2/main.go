package main

import (
	"fmt"
	"hw1.2/calc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Enter the expression to the args")
	}
	input := os.Args[1]
	result, err := calc.Calc(input)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(result)
}

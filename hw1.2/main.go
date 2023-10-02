package main

import (
	"fmt"
	"hw1.2/calc"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Println("enter the expression to the args")
		return
	}
	input := os.Args[1]
	result, err := calc.Calc(input)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(result)
}

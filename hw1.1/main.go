package main

import (
	"flag"
	"fmt"
	"hw1.1/flagsparser"
	"hw1.1/iostrings"
	"hw1.1/uniq"
	"log"
)

func main() {
	options, err := flagsparser.ParseFlags()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if len(flag.Args()) > 2 {
		flag.Usage()
		log.Println("too much arguments")
		return
	}

	inputStrings, err := iostrings.Read()
	if err != nil {
		log.Println(err.Error())
		return
	}

	outputStrings, err := uniq.Uniq(inputStrings, options)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if err = iostrings.Write(outputStrings); err != nil {
		fmt.Println(err.Error())
		return
	}
}

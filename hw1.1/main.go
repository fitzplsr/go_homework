package main

import (
	"flag"
	"hw1.1/flagsparser"
	"hw1.1/iostrings"
	"hw1.1/uniq"
	"log"
)

func main() {
	options, err := flagsparser.ParseFlags()
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(flag.Args()) > 2 {
		flagsparser.Usage()
		log.Fatal("too much arguments")
	}

	inputStrings, err := iostrings.Read()
	if err != nil {
		log.Fatal(err)
	}

	outputStrings := uniq.Uniq(inputStrings, options)
	if err = iostrings.Write(outputStrings); err != nil {
		log.Fatal(err.Error())
	}
}
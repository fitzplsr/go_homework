package flagsparser

import (
	"errors"
	"flag"
	"fmt"
	"hw1.1/uniq"
)

const (
	cUsage      = "calculate the number of occurrences of the string in the input data"
	dUsage      = "output only those lines that were repeated in the input data"
	uUsage      = "output only those lines that are not repeated in the input data."
	iUsage      = "ignore case of letters"
	fUsage      = "ignore the first num_fields of fields in a row"
	sUsage      = "ignore the first num_chars characters in the string"
	usageSchema = "Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"
)

func Usage() {
	_, err := fmt.Println(usageSchema)
	if err != nil {
		return
	}
	flag.PrintDefaults()
}

func ParseFlags() (uniq.Options, error) {
	cFlagPtr := flag.Bool("c", false, cUsage)
	dFlagPtr := flag.Bool("d", false, dUsage)
	uFlagPtr := flag.Bool("u", false, uUsage)
	iFlagPtr := flag.Bool("i", false, iUsage)
	var numFields uint
	var numChars uint
	flag.UintVar(&numFields, "f", 0, fUsage)
	flag.UintVar(&numChars, "s", 0, sUsage)
	flag.Parse()
	if *cFlagPtr && *dFlagPtr || *cFlagPtr && *uFlagPtr || *uFlagPtr && *dFlagPtr {
		Usage()
		return uniq.Options{}, errors.New("you can't use flags -u -c -d together")
	}
	return uniq.Options{
		C: *cFlagPtr,
		D: *dFlagPtr,
		U: *uFlagPtr,
		I: *iFlagPtr,
		F: numFields,
		S: numChars,
	}, nil
}

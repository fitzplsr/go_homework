package iostrings

import (
	"bufio"
	"flag"
	"log"
	"os"
)

func Read() (input []string, err error) {
	reader := os.Stdin
	inputFile := flag.Arg(0)
	if inputFile != "" {
		var file *os.File
		file, err = os.Open(inputFile)
		defer func(file *os.File) {
			fileCloseErr := file.Close()
			if fileCloseErr != nil {
				log.Println(fileCloseErr.Error())
				return
			}
		}(file)
		if err != nil {
			log.Println(err.Error())
			return
		}
		reader = file
	}

	in := bufio.NewScanner(reader)
	for in.Scan() {
		input = append(input, in.Text())
	}
	return
}

func Write(data []string) (err error) {
	writer := os.Stdout
	outputFile := flag.Arg(1)
	if outputFile != "" {
		var file *os.File
		file, err = os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				log.Println(err.Error())
				return
			}
		}(file)
		if err != nil {
			return
		}
		writer = file
	}

	out := bufio.NewWriter(writer)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			log.Println(err.Error())
			return
		}
	}(out)

	for _, line := range data {
		if _, err = out.WriteString(line + "\n"); err != nil {
			return
		}
	}
	return
}

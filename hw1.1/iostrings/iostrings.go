package iostrings

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
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
				log.Fatal("Can't close file")
			}
		}(file)
		if err != nil {
			err = errors.New("can't open file")
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
				log.Fatal("Can't close file")
			}
		}(file)
		if err != nil {
			err = fmt.Errorf("Can't append or create file: %s\n", outputFile)
			return
		}
		writer = file
	}

	out := bufio.NewWriter(writer)
	defer func(out *bufio.Writer) {
		err := out.Flush()
		if err != nil {
			log.Fatal("Error flushing buffer")
		}
	}(out)

	for _, line := range data {
		if _, err = out.WriteString(line + "\n"); err != nil {
			return errors.New("error writing file")
		}
	}
	return
}

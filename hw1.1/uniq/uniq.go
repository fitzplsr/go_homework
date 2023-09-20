package uniq

import (
	"errors"
	"strconv"
	"strings"
)

type Options struct {
	C bool
	D bool
	U bool
	I bool
	F int
	S int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func equalStrings(a string, b string) bool {
	return a == b
}

func Uniq(input []string, options Options) (output []string, err error) {
	if options.F < 0 || options.S < 0 {
		err = errors.New("-f -s flags must be positive integer")
		return
	}
	var prevLine, prevCutLine, currentLine string
	var currentLineIndex int
	isFirst := true
	countOfRepeated := 1
	compare := equalStrings
	if options.I {
		compare = strings.EqualFold
	}

	record := func() {
		switch {
		case options.C:
			output = append(output, strconv.Itoa(countOfRepeated)+" "+prevLine)
		case options.U:
			if countOfRepeated == 1 {
				output = append(output, prevLine)
			}
		case options.D:
			if countOfRepeated > 1 {
				output = append(output, prevLine)
			}
		default:
			output = append(output, prevLine)
		}
	}

	for _, line := range input {
		currentLine = line
		if options.F != 0 && len(line) != 0 {
			fields := strings.Fields(line)
			if options.F > len(fields)-1 {
				currentLineIndex = len(line)
				currentLine = ""
			} else {
				currentLineIndex = strings.Index(line, fields[options.F])
				currentLine = line[currentLineIndex:]
			}
		}

		if options.S != 0 && len(line) != 0 && currentLineIndex != len(line) {
			currentLine = string([]rune(currentLine)[min(options.S, len([]rune(currentLine))-1):])
		}

		if isFirst {
			prevLine = line
			prevCutLine = currentLine
			isFirst = false
			continue
		}

		isMatch := compare(currentLine, prevCutLine)

		if !isMatch {
			record()
			prevLine = line
			prevCutLine = currentLine
			countOfRepeated = 0
		}
		countOfRepeated++
	}
	record()
	return
}

package uniq

import (
	"strconv"
	"strings"
)

type Options struct {
	C bool
	D bool
	U bool
	I bool
	F uint
	S uint
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Uniq(input []string, options Options) (output []string) {
	var prevLine, prevCutLine, currentLine string
	var currentLineIndex int
	isFirst := true
	countOfRepeated := 1
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
			if options.F > (uint)(len(fields)-1) {
				currentLineIndex = len(line)
				currentLine = ""
			} else {
				currentLineIndex = strings.Index(line, fields[options.F])
			}
		}
		if len(line) != 0 && currentLineIndex != len(line) {
			currentLine = line[min(currentLineIndex+(int)(options.S), len(line)-1):]
		}
		if isFirst {
			prevLine = line
			prevCutLine = currentLine
			isFirst = false
			continue
		}
		isMatch := currentLine == prevCutLine
		if options.I {
			isMatch = strings.EqualFold(currentLine, prevCutLine)
		}
		if !isMatch {
			record()
			prevLine = line
			prevCutLine = currentLine
			countOfRepeated = 1
			continue
		}
		countOfRepeated++
	}
	record()
	return
}

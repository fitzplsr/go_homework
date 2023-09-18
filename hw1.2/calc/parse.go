package calc

import (
	"regexp"
	"strconv"
	"unicode"
)

func isOperation(c byte) (result bool) {
	switch c {
	case '-', '*', '/', '+', '(', ')':
		result = true
	}
	return
}

func isUnaryMinus(stack *Stack) bool {
	return stack.Length() == 0 || stack.Back().Operation != 0 && stack.Back().Operation != ')'
}

func parseInput(stack *Stack, str string) (err error) {
	re := regexp.MustCompile(`^[+\-]?(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)`)
	var symbol byte
	isNumberOrder := true
	for i := 0; i < len([]rune(str)); i++ {
		if symbol = str[i]; symbol == ' ' {
			continue
		}
		if isOperation(str[i]) {
			if symbol == '-' && isUnaryMinus(stack) {
				symbol = '~'
			}
			stack.Push(Item{Operation: symbol})
			isNumberOrder = true
		} else if isNumberOrder && unicode.IsDigit(rune(symbol)) { //
			numStr := re.Find([]byte((str[i:])))
			var num float64
			num, err = strconv.ParseFloat(string(numStr), 64)
			if err != nil {
				break
			}
			i += len(numStr) - 1
			stack.Push(Item{Number: num})
			isNumberOrder = false
		} else {
			err = &IncorrectExpressionError{"error parse expression"}
			break
		}
	}
	return
}

package calc

func Calc(input string) (result float64, err error) {
	var stack Stack
	if err = parseInput(&stack, input); err != nil {
		return
	}

	stack.Reverse()
	if err = convertToPostfix(&stack); err != nil {
		return
	}

	if result, err = compute(&stack); err != nil {
		return
	}
	return
}

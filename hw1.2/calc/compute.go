package calc

import (
	"fmt"
)

type ComputeError struct {
	message string
}

func (e *ComputeError) Error() string {
	return fmt.Sprintf("Error computing expression: %s", e.message)
}

func doUnary(result *Stack, taken byte, item Item) (err error) {
	if taken == '~' {
		item.Number = (-1) * item.Number
	} else {
		err = &ComputeError{"undefined operator"}
		return
	}
	result.Push(item)
	return
}

func doBinary(result *Stack, taken byte, first Item, second Item) (err error) {
	var item Item
	switch taken {
	case '-':
		item.Number = second.Number - first.Number
	case '+':
		item.Number = second.Number + first.Number
	case '*':
		item.Number = second.Number * first.Number
	case '/':
		if first.Number == 0 {
			err = &ComputeError{"zero division error"}
			return
		}
		item.Number = second.Number / first.Number
	}
	result.Push(item)
	return
}

func doOperation(result *Stack, taken byte) (err error) {
	if getPriority(taken) == 3 {
		if result.Length() > 0 {
			item, _ := result.Pop()
			err = doUnary(result, taken, item)
		} else {
			err = &ComputeError{"error doing operation"}
		}
	} else if result.Length() > 1 {
		first, _ := result.Pop()
		second, _ := result.Pop()
		err = doBinary(result, taken, first, second)
	} else {
		err = &ComputeError{"error doing operation"}
	}
	return
}

func compute(stack *Stack) (resultNum float64, err error) {
	var result Stack
	for err == nil && stack.Length() > 0 {
		if stack.Back().Operation == 0 {
			item, _ := stack.Pop()
			result.Push(item)
		} else {
			taken, _ := stack.Pop()
			err = doOperation(&result, taken.Operation)
		}
	}
	if result.Length() > 1 {
		err = &ComputeError{"missed operator"}
	}
	if err != nil {
		return
	}
	item, _ := result.Pop()
	resultNum = item.Number
	return
}

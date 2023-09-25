package calc

import (
	"fmt"
)

type ComputeError struct {
	message string
}

func (e *ComputeError) Error() string {
	return fmt.Sprintf("error computing expression: %s", e.message)
}

func doUnary(result *Stack, taken byte, item Item) error {
	switch taken {
	case '~':
		item.Number = (-1) * item.Number
		result.Push(item)
	default:
		return &ComputeError{"undefined operator"}
	}
	return nil
}

func doBinary(result *Stack, taken byte, first Item, second Item) error {
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
			return &ComputeError{"zero division error"}
		}
		item.Number = second.Number / first.Number
	default:
		return &ComputeError{"undefined operator"}
	}
	result.Push(item)
	return nil
}

func doOperation(result *Stack, taken byte) error {
	if priorities[taken] == 3 {
		if result.Length() < 1 {
			return &ComputeError{"error doing operation"}
		}
		item, _ := result.Pop()
		return doUnary(result, taken, item)
	}

	if result.Length() < 2 {
		return &ComputeError{"error doing operation"}
	}

	first, _ := result.Pop()
	second, _ := result.Pop()
	return doBinary(result, taken, first, second)
}

func compute(stack *Stack) (resultNum float64, err error) {
	var result Stack
	for err == nil && stack.Length() > 0 {
		if stack.Back().Operation == 0 {
			item, _ := stack.Pop()
			result.Push(item)
			continue
		}
		taken, _ := stack.Pop()
		err = doOperation(&result, taken.Operation)
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

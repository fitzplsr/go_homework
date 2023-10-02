package calc

import (
	"fmt"
)

var priorities = map[byte]int{
	'(': 0,
	')': 0,
	'+': 1,
	'-': 1,
	'*': 2,
	'/': 2,
	'~': 3,
}

type IncorrectExpressionError struct {
	message string
}

func (e *IncorrectExpressionError) Error() string {
	return fmt.Sprintf("incorrect expression: %s", e.message)
}

func takeFromStack(operation byte, priority int) bool {
	if priority == 0 {
		return operation != '('
	}
	return priorities[operation] >= priority
}

func takeLower(operations *Stack, result *Stack, priority int) (err error) {
	for operations.Length() > 0 && takeFromStack(operations.Back().Operation, priority) {
		item := operations.Pop()
		result.Push(item)
	}
	if priority == 0 && operations.Length() == 0 {
		err = &IncorrectExpressionError{"empty operations stack"}
	}
	return
}

func moveItem(result *Stack, operations *Stack, item Item, openBracket *bool) error {
	switch item.Operation {
	case 0:
		*openBracket = false
		result.Push(item)
	case '(':
		*openBracket = true
		operations.Push(item)
	case ')':
		if *openBracket {
			return &IncorrectExpressionError{"empty inside brackets"}
		}
		fallthrough
	default:
		err := takeLower(operations, result, priorities[item.Operation])
		if err != nil {
			return &IncorrectExpressionError{"incorrect bracket order"}
		}
		if item.Operation == ')' {
			operations.Pop()
		} else {
			*openBracket = false
			operations.Push(item)
		}
	}
	return nil
}

func convertToPostfix(stack *Stack) (err error) {
	var result, operations Stack
	openBracket := false
	for err == nil && stack.Length() > 0 {
		item := stack.Pop()
		err = moveItem(&result, &operations, item, &openBracket)
	}
	for err == nil && operations.Length() > 0 {
		item := operations.Pop()
		if item.Operation == '(' {
			err = &IncorrectExpressionError{"incorrect bracket order"}
			break
		}
		result.Push(item)
	}
	for err == nil && result.Length() > 0 {
		item := result.Pop()
		stack.Push(item)
	}
	return
}

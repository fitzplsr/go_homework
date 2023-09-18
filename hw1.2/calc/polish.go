package calc

import (
	"fmt"
)

type IncorrectExpressionError struct {
	message string
}

func (e *IncorrectExpressionError) Error() string {
	return fmt.Sprintf("Incorrect expression: %s", e.message)
}

func getPriority(operation byte) (priority int) {
	switch operation {
	case '(', ')':
		priority = 0
	case '+', '-':
		priority = 1
	case '*', '/':
		priority = 2
	case '~':
		priority = 3
	}
	return
}

func takeFromStack(operation byte, priority int) (result bool) {
	if priority == 0 {
		result = operation != '('
	} else {
		result = getPriority(operation) >= priority
	}
	return
}

func takeLower(operations *Stack, result *Stack, priority int) (err error) {
	for operations.Length() > 0 && takeFromStack(operations.Back().Operation, priority) {
		item, _ := operations.Pop()
		result.Push(item)
	}
	if priority == 0 && operations.Length() == 0 {
		err = &IncorrectExpressionError{"empty operations stack"}
	}
	return
}

func moveItem(result *Stack, operations *Stack, item Item, openBracket *bool) (err error) {
	if item.Operation == 0 {
		*openBracket = false
		result.Push(item)
	} else if item.Operation == '(' {
		*openBracket = true
		operations.Push(item)
	} else if item.Operation == ')' && *openBracket {
		err = &IncorrectExpressionError{"empty inside brackets"}
	} else if takeLower(operations, result, getPriority(item.Operation)) == nil {
		if item.Operation == ')' {
			operations.Pop()
		} else {
			*openBracket = false
			operations.Push(item)
		}
	} else {
		err = &IncorrectExpressionError{"incorrect bracket order"}
	}
	return
}

func convertToPostfix(stack *Stack) (err error) {
	var result, operations Stack
	openBracket := false
	for err == nil && stack.Length() > 0 {
		item, _ := stack.Pop()
		err = moveItem(&result, &operations, item, &openBracket)
	}
	for err == nil && operations.Length() > 0 {
		item, _ := operations.Pop()
		if item.Operation == '(' {
			err = &IncorrectExpressionError{"incorrect bracket order"}
			break
		}
		result.Push(item)
	}
	for err == nil && result.Length() > 0 {
		item, _ := result.Pop()
		stack.Push(item)
	}
	return
}

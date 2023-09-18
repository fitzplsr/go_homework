package calc

import "fmt"

type Stack struct {
	items []Item
}

func (s *Stack) Length() int {
	return len(s.items)
}

func (s *Stack) Back() Item {
	return s.items[len(s.items)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Push(item Item) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() (element Item, isEmpty bool) {
	if s.IsEmpty() {
		return
	} else {
		index := len(s.items) - 1
		element := (s.items)[index]
		s.items = (s.items)[:index]
		return element, true
	}
}

func (s *Stack) Reverse() {
	if s.IsEmpty() {
		return
	}
	reversed := make([]Item, 0, len(s.items))
	for i := range s.items {
		reversed = append(reversed, s.items[len(s.items)-1-i])
	}
	s.items = reversed
}

func (s *Stack) Print() {
	for _, item := range s.items {
		fmt.Printf("Operation: %c, Number: %g\n", rune(item.Operation), item.Number)
	}
}

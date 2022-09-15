package main

import "fmt"

// Stack represents a stack data structure
type Stack struct {
	Data []any
}

// Push adds an element to the stack
func (s *Stack) Push(data any) {
	s.Data = append(s.Data, data)
}

// Pop removes the last element from the stack
func (s *Stack) Pop() any {
	if len(s.Data) == 0 {
		return nil
	}
	data := s.Data[len(s.Data)-1]
	s.Data = s.Data[:len(s.Data)-1]
	return data
}

func main() {
	mathExp := "1*2+(5*3)-(2+(3*5))"
	fmt.Println(checkParenthesis(mathExp)) // true
}

var (
	// listOfOpenAndCloseParenthesis is a map of open and close parenthesis
	listOfOpenAndCloseParenthesis = map[string]string{"(": ")", "{": "}", "[": "]"}
	// listOfCloseAndOpenParenthesis is a map of close and open parenthesis
	listOfCloseAndOpenParenthesis = map[string]string{")": "(", "}": "{", "]": "["}
)

// checkParenthesis checks if the parenthesis in a math expression are balanced
func checkParenthesis(exp string) bool {
	// Create a stack
	s := &Stack{Data: []any{}}
	// Loop through the math expression
	for _, v := range exp {
		// Convert the rune to string
		// Check if the current character is an open parenthesis
		if _, ok := listOfOpenAndCloseParenthesis[string(v)]; ok {
			// If it is, push it to the stack
			s.Push(string(v))
		} else {
			// If it is not, check if it is a close parenthesis
			if openTag, ok := listOfCloseAndOpenParenthesis[string(v)]; ok {
				// If it is, check if the stack is empty or if the last element in the stack is not the corresponding open parenthesis
				if s.Pop() != openTag {
					// If it is, return false
					return false
				}
			}
		}
	}
	// If the stack is not empty, return false
	if len(s.Data) != 0 {
		return false
	}

	// If the stack is empty, return true
	return true
}

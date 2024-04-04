package main

import ()

type StringStack struct {
	s []string
}

func (s *StringStack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *StringStack) Length() int {
	return len(s.s)
}

func (s *StringStack) Push(str string) {
	s.s = append(s.s, str)
}

func (s *StringStack) Pop() string {

	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[0 : length-1]
	return res

}

func (s *StringStack) Peek() string {

	length := len(s.s)
	res := s.s[length-1]
	return res

}

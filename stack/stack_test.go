package stack

import "testing"

func TestStack(t *testing.T) {
	var s4s Stack
	s4s.push("abc")
	s4s.push("def")
	s4s.push("ghi")

	fPln(s4s.sprint("-"))

	val, ok := s4s.pop()
	fPln(val, ok, s4s)
	fPln(s4s.sprint("-"))

	val, ok = s4s.pop()
	s4s.push("jkl")
	fPln(val, ok, s4s)
	fPln(s4s.sprint("-"))

	val, ok = s4s.pop()
	fPln(val, ok, s4s)
	fPln(s4s.sprint("-"))

	val, ok = s4s.pop()
	fPln(val, ok, s4s)
	fPln(s4s.sprint("-"))
}
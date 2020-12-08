package stack

import "testing"

func TestStack(t *testing.T) {
	var s4s Stack
	s4s.Push("abc")
	s4s.Push("def")
	s4s.Push("ghi")
	s4s.Push("jkl")

	tmp := s4s.Copy()

	fPln(s4s.Sprint("-"))

	val, ok := s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	val, ok = s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	s4s.Push("***")
	fPln(s4s, s4s.Sprint("-"))

	val, ok = s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	val, ok = s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	val, ok = s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	val, ok = s4s.Pop()
	fPln(val, ok, s4s, s4s.Sprint("-"))

	fPln(tmp.Sprint("/"))
}

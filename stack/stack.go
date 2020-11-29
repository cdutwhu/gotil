package stack

// Stack :
type Stack []interface{}

// Push :
func (stk *Stack) Push(items ...interface{}) (n int) {
	for i, s := range items {
		*stk = append(*stk, s)
		n = i
	}
	return n + 1
}

func (stk *Stack) len() int {
	return len(*stk)
}

// Len :
func (stk *Stack) Len() int {
	return len(*stk)
}

// Pop :
func (stk *Stack) Pop() (interface{}, bool) {
	if stk.len() > 0 {
		last := (*stk)[stk.len()-1]
		*stk = (*stk)[:stk.len()-1]
		return last, true
	}
	return nil, false
}

// Peek :
func (stk *Stack) Peek() (interface{}, bool) {
	if stk.len() > 0 {
		return (*stk)[stk.len()-1], true
	}
	return nil, false
}

// Clear :
func (stk *Stack) Clear() (n int) {
	n = stk.len()
	stk = &Stack{}
	return n
}

// Sprint :
func (stk *Stack) Sprint(sep string) string {
	sb := sBuilder{}
	for _, ele := range *stk {
		sb.WriteString(fSf("%v", ele))
		sb.WriteString(sep)
	}
	return sTrimRight(sb.String(), sep)
}

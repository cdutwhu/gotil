package stack

// Stack :
type Stack []interface{}

func (stk *Stack) push(items ...interface{}) (n int) {
	for i, s := range items {
		*stk = append(*stk, s)
		n = i
	}
	return n + 1
}

func (stk *Stack) len() int {
	return len(*stk)
}

func (stk *Stack) pop() (interface{}, bool) {
	if stk.len() > 0 {
		last := (*stk)[stk.len()-1]
		*stk = (*stk)[:stk.len()-1]
		return last, true
	}
	return nil, false
}

func (stk *Stack) peek() (interface{}, bool) {
	if stk.len() > 0 {
		return (*stk)[stk.len()-1], true
	}
	return nil, false
}

func (stk *Stack) clear() (n int) {
	n = stk.len()
	stk = &Stack{}
	return n
}

func (stk *Stack) sprint(sep string) string {
	sb := sBuilder{}
	for _, ele := range *stk {
		sb.WriteString(fSf("%v", ele))
		sb.WriteString(sep)
	}
	return sTrimRight(sb.String(), sep)
}

package iter

// N : for i := range N(n)
func N(n int) []struct{} {
	return make([]struct{}, n)
}

// Iter : for i := range Iter(end)/(start,end)/(start,step,end) 
func Iter(params ...int) <-chan int {

	start, end, step := 0, 0, 1
	switch len(params) {
	case 1:
		end = params[0]
	case 2:
		start, end = params[0], params[1]
	case 3:
		start, step, end = params[0], params[1], params[2]
	default:
		failP1OnErr("%v: params' count only can be 1, 2 or 3", fEf("PARAM_INVALID"))
	}

	ch := make(chan int)

	if start > end {

		switch len(params) {
		case 1, 2:
			step = -1
		}

		failP1OnErrWhen(step >= 0, "%v", fEf("step error, must be NEGATIVE"))
		go func() {
			defer close(ch)
			for i := start; i > end; i += step {
				ch <- i
			}
		}()

	} else {

		failP1OnErrWhen(step <= 0, "%v", fEf("step error, must be POSITIVE"))
		go func() {
			defer close(ch)
			for i := start; i < end; i += step {
				ch <- i
			}
		}()
	}

	return ch
}

// Iter2Slc :
func Iter2Slc(params ...int) (slc []int) {
	if len(params) == 1 {
		for i := range N(params[0]) {
			slc = append(slc, i)
		}
		return
	}
	for i := range Iter(params...) {
		slc = append(slc, i)
	}
	return
}

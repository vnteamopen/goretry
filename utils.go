package goretry

func intPow(x int64, degree int) int64 {
	if degree == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	ret := int64(x)
	for i := 2; i <= degree; i++ {
		ret *= x
	}
	return ret
}

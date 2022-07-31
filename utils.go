package goretry

func intPow(x, degree int64) int64 {
	if degree == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	ret := int64(x)
	for i := int64(2); i <= degree; i++ {
		ret *= x
	}
	return ret
}

package stats

// ProbGeom generates the probability for a geometric random variable
// with parameter p to achieve success in the interval of [a, b] trials
// See https://en.wikipedia.org/wiki/Geometric_distribution for more information
func ProbGeom(a int, b int, p float64) (prob float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// probability of failure

// ProbGeom generates the expectation or average number of trials
// for a geometric random variable with parameter p
func ExpGeom(p float64) (exp float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// ProbGeom generates the variance for number for a
// geometric random variable with parameter p
func VarGeom(p float64) (exp float64, err error) { _ = "STUB: not implemented"; return 0, nil }

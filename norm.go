package stats

// NormSample generates random samples from a normal distribution
// with the given mean (loc) and standard deviation (scale).
func NormSample(loc float64, scale float64, size int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// NormPpfRvs generates random variates using the Point Percentile Function.
// For more information please visit: https://demonstrations.wolfram.com/TheMethodOfInverseTransforms/
func NormPpfRvs(loc float64, scale float64, size int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// NormBoxMullerRvs generates random variates using the Box–Muller transform.
// For more information please visit: http://mathworld.wolfram.com/Box-MullerTransformation.html
func NormBoxMullerRvs(loc float64, scale float64, size int) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// u1 and u2 are uniformly distributed random numbers between 0 and 1.

// x1 and x2 are normally distributed random numbers.

// NormPdf is the probability density function.
func NormPdf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormLogPdf is the log of the probability density function.
func NormLogPdf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormCdf is the cumulative distribution function.
func NormCdf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormLogCdf is the log of the cumulative distribution function.
func NormLogCdf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormSf is the survival function (also defined as 1 - cdf, but sf is sometimes more accurate).
func NormSf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormLogSf is the log of the survival function.
func NormLogSf(x float64, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormPpf is the point percentile function.
// This is based on Peter John Acklam's inverse normal CDF.
// algorithm: http://home.online.no/~pjacklam/notes/invnorm/ (no longer visible).
// For more information please visit: https://stackedboxes.org/2017/05/01/acklams-normal-quantile-function/
func NormPpf(p float64, loc float64, scale float64) (x float64) {
	_ = "STUB: not implemented"
	return 0
}

// NormIsf is the inverse survival function (inverse of sf).
func NormIsf(p float64, loc float64, scale float64) (x float64) {
	_ = "STUB: not implemented"
	return 0
}

// NormMoment approximates the non-central (raw) moment of order n.
// For more information please visit: https://math.stackexchange.com/questions/1945448/methods-for-finding-raw-moments-of-the-normal-distribution
func NormMoment(n int, loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormStats returns the mean, variance, skew, and/or kurtosis.
// Mean(‘m’), variance(‘v’), skew(‘s’), and/or kurtosis(‘k’).
// Takes string containing any of 'mvsk'.
// Returns array of m v s k in that order.
func NormStats(loc float64, scale float64, moments string) []float64 {
	_ = "STUB: not implemented"
	return nil
}

// NormEntropy is the differential entropy of the RV.
func NormEntropy(loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormFit returns the maximum likelihood estimators for the Normal Distribution.
// Takes array of float64 values.
// Returns array of Mean followed by Standard Deviation.
func NormFit(data []float64) [2]float64 { _ = "STUB: not implemented"; return nil }

// NormMedian is the median of the distribution.
func NormMedian(loc float64, scale float64) float64 {
	_ = "STUB: not implemented"

	// NormMean is the mean/expected value of the distribution.
	return 0
}

func NormMean(loc float64, scale float64) float64 {
	_ = "STUB: not implemented"

	// NormVar is the variance of the distribution.
	return 0
}

func NormVar(loc float64, scale float64) float64 { _ = "STUB: not implemented"; return 0 }

// NormStd is the standard deviation of the distribution.
func NormStd(loc float64, scale float64) float64 {
	_ = "STUB: not implemented"

	// NormInterval finds endpoints of the range that contains alpha percent of the distribution.
	return 0
}

func NormInterval(alpha float64, loc float64, scale float64) [2]float64 {
	_ = "STUB: not implemented"
	return nil
}

// factorial is the naive factorial algorithm.
func factorial(x int) int { _ = "STUB: not implemented"; return 0 }

// Ncr is an N choose R algorithm.
// Aaron Cannon's algorithm.
func Ncr(n, r int) int { _ = "STUB: not implemented"; return 0 }

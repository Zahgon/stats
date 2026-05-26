package stats

// Series is a container for a series of data
type Series []Coordinate

// Coordinate holds the data in a series
type Coordinate struct {
	X, Y float64
}

// LinearRegression finds the least squares linear regression on data series
func LinearRegression(s Series) (regressions Series, err error) {
	_ = "STUB: not implemented"
	return *new(Series), nil
}

// Placeholder for the math to be done

// Loop over data keeping index in place

// Find gradient and intercept

// Create the new regression series

// ExponentialRegression returns an exponential regression on data series
func ExponentialRegression(s Series) (regressions Series, err error) {
	_ = "STUB: not implemented"
	return *new(Series), nil
}

// LogarithmicRegression returns an logarithmic regression on data series
func LogarithmicRegression(s Series) (regressions Series, err error) {
	_ = "STUB: not implemented"
	return *new(Series), nil
}

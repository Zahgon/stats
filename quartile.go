package stats

// Quartiles holds the three quartile points
type Quartiles struct {
	Q1 float64
	Q2 float64
	Q3 float64
}

// Quartile returns the three quartile points from a slice of data
func Quartile(input Float64Data) (Quartiles, error) {
	_ = "STUB: not implemented"
	return *new(Quartiles), nil
}

// Start by sorting a copy of the slice

// Find the cutoff places depeding on if
// the input slice length is even or odd

// Find the Medians with the cutoff points

// InterQuartileRange finds the range between Q1 and Q3
func InterQuartileRange(input Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Midhinge finds the average of the first and third quartiles
func Midhinge(input Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Trimean finds the average of the median and the midhinge
func Trimean(input Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

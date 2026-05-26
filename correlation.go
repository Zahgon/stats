package stats

// Correlation describes the degree of relationship between two sets of data
func Correlation(data1, data2 Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Pearson calculates the Pearson product-moment correlation coefficient between two variables
func Pearson(data1, data2 Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Spearman calculates the Spearman rank correlation coefficient between two variables.
// It works by ranking the data and then computing the Pearson correlation of the ranks.
// This method handles tied values using fractional (average) ranking.
func Spearman(data1, data2 Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// rankData assigns fractional (average) ranks to the data values.
// Tied values receive the average of the ranks they would have been assigned.
func rankData(data Float64Data) Float64Data {
	_ = "STUB: not implemented"

	// Create index-value pairs and sort by value
	return *new(Float64Data)
}

// Assign fractional ranks handling ties

// Average rank for tied values (ranks are 1-based)

// AutoCorrelation is the correlation of a signal with a delayed copy of itself as a function of delay
func AutoCorrelation(data Float64Data, lags int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

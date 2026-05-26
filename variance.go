package stats

// _variance finds the variance for both population and sample data
func _variance(input Float64Data, sample int) (variance float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Sum the square of the mean subtracted from each number

// When getting the mean of the squared differences
// "sample" will allow us to know if it's a sample
// or population and wether to subtract by one or not

// Variance the amount of variation in the dataset
func Variance(input Float64Data) (sdev float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// PopulationVariance finds the amount of variance within a population
func PopulationVariance(input Float64Data) (pvar float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// SampleVariance finds the amount of variance within a sample
func SampleVariance(input Float64Data) (svar float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Covariance is a measure of how much two sets of data change
func Covariance(data1, data2 Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Calculate sum of squares

// CovariancePopulation computes covariance for entire population between two variables.
func CovariancePopulation(data1, data2 Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

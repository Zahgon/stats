package stats

// Skewness computes the population skewness of the dataset
func Skewness(input Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// PopulationSkewness computes the population skewness using the third
// central moment normalized by the cube of the standard deviation.
func PopulationSkewness(input Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Compute sum of squared and cubed differences from the mean

// SampleSkewness computes the adjusted Fisher-Pearson standardized moment
// coefficient, correcting for bias in small samples.
func SampleSkewness(input Float64Data) (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// Adjusted Fisher-Pearson: G1 = g1 * sqrt(n*(n-1)) / (n-2)

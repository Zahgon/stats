package stats

// Outliers holds mild and extreme outliers found in data
type Outliers struct {
	Mild    Float64Data
	Extreme Float64Data
}

// QuartileOutliers finds the mild and extreme outliers
func QuartileOutliers(input Float64Data) (Outliers, error) {
	_ = "STUB: not implemented"
	return *new(Outliers), nil
}

// Start by sorting a copy of the slice

// Calculate the quartiles and interquartile range

// Calculate the lower and upper inner and outer fences

// Find the data points that are outside of the
// inner and upper fences and add them to mild
// and extreme outlier slices

// Wrap them into our struct

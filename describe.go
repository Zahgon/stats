package stats

// Holds information about the dataset provided to Describe
type Description struct {
	Count                  int
	Mean                   float64
	Std                    float64
	Max                    float64
	Min                    float64
	DescriptionPercentiles []descriptionPercentile
	AllowedNaN             bool
}

// Specifies percentiles to be computed
type descriptionPercentile struct {
	Percentile float64
	Value      float64
}

// Describe generates descriptive statistics about a provided dataset, similar to python's pandas.describe()
func Describe(input Float64Data, allowNaN bool, percentiles *[]float64) (*Description, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Describe generates descriptive statistics about a provided dataset, similar to python's pandas.describe()
// Takes in a function to use for percentile calculation
func DescribePercentileFunc(input Float64Data, allowNaN bool, percentiles *[]float64, percentileFunc func(Float64Data, float64) (float64, error)) (*Description, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Disregard error, since it cannot be thrown if Count is > 0 and allowNaN is false, else NaN is accepted

/*
Represents the Description instance in a string format with specified number of decimals

	count   3
	mean    2.00
	std     0.82
	max     3.00
	min     1.00
	25.00%  NaN
	50.00%  1.50
	75.00%  2.50
	NaN OK  true
*/
func (d *Description) String(decimals int) string { _ = "STUB: not implemented"; return "" }

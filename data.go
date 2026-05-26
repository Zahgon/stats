package stats

// Float64Data is a named type for []float64 with helper methods
type Float64Data []float64

// Get item in slice
func (f Float64Data) Get(i int) float64 {
	_ = "STUB: not implemented"

	// Len returns length of slice
	return 0
}

func (f Float64Data) Len() int {
	_ = "STUB: not implemented"

	// Less returns if one number is less than another
	return 0
}

func (f Float64Data) Less(i, j int) bool {
	_ = "STUB: not implemented"

	// Swap switches out two numbers in slice
	return false
}

func (f Float64Data) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Min returns the minimum number in the data
func (f Float64Data) Min() (float64, error) {
	_ = "STUB: not implemented"

	// Max returns the maximum number in the data
	return 0, nil
}

func (f Float64Data) Max() (float64, error) {
	_ = "STUB: not implemented"

	// Sum returns the total of all the numbers in the data
	return 0, nil
}

func (f Float64Data) Sum() (float64, error) {
	_ = "STUB: not implemented"

	// CumulativeSum returns the cumulative sum of the data
	return 0, nil
}

func (f Float64Data) CumulativeSum() ([]float64, error) {
	_ = "STUB: not implemented"
	return nil,

		// Mean returns the mean of the data
		nil
}

func (f Float64Data) Mean() (float64, error) {
	_ = "STUB: not implemented"

	// Median returns the median of the data
	return 0, nil
}

func (f Float64Data) Median() (float64, error) {
	_ = "STUB: not implemented"

	// Mode returns the mode of the data
	return 0, nil
}

func (f Float64Data) Mode() ([]float64, error) {
	_ = "STUB: not implemented"

	// GeometricMean returns the geometric mean of the data
	return nil, nil
}

func (f Float64Data) GeometricMean() (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// HarmonicMean returns the harmonic mean of the data
		nil
}

func (f Float64Data) HarmonicMean() (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// MedianAbsoluteDeviation the median of the absolute deviations from the dataset median
		nil
}

func (f Float64Data) MedianAbsoluteDeviation() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MedianAbsoluteDeviationPopulation finds the median of the absolute deviations from the population median
func (f Float64Data) MedianAbsoluteDeviationPopulation() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// StandardDeviation the amount of variation in the dataset
func (f Float64Data) StandardDeviation() (float64, error) { _ = "STUB: not implemented"; return 0, nil }

// StandardDeviationPopulation finds the amount of variation from the population
func (f Float64Data) StandardDeviationPopulation() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// StandardDeviationSample finds the amount of variation from a sample
func (f Float64Data) StandardDeviationSample() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// QuartileOutliers finds the mild and extreme outliers
func (f Float64Data) QuartileOutliers() (Outliers, error) {
	_ = "STUB: not implemented"
	return *new(Outliers), nil
}

// Percentile finds the relative standing in a slice of floats
func (f Float64Data) Percentile(p float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// PercentileNearestRank finds the relative standing using the Nearest Rank method
		nil
}

func (f Float64Data) PercentileNearestRank(p float64) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Correlation describes the degree of relationship between two sets of data
func (f Float64Data) Correlation(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// AutoCorrelation is the correlation of a signal with a delayed copy of itself as a function of delay
		nil
}

func (f Float64Data) AutoCorrelation(lags int) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Pearson calculates the Pearson product-moment correlation coefficient between two variables.
func (f Float64Data) Pearson(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Spearman calculates the Spearman rank correlation coefficient between two variables.
		nil
}

func (f Float64Data) Spearman(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Quartile returns the three quartile points from a slice of data
		nil
}

func (f Float64Data) Quartile(d Float64Data) (Quartiles, error) {
	_ = "STUB: not implemented"
	return *

	// InterQuartileRange finds the range between Q1 and Q3
	new(Quartiles), nil
}

func (f Float64Data) InterQuartileRange() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil

	// Midhinge finds the average of the first and third quartiles
}

func (f Float64Data) Midhinge(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Trimean finds the average of the median and the midhinge
		nil
}

func (f Float64Data) Trimean(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"

	// Sample returns sample from input with replacement or without
	return 0, nil
}

func (f Float64Data) Sample(n int, r bool) ([]float64, error) {
	_ = "STUB: not implemented"
	return nil,

		// Variance the amount of variation in the dataset
		nil
}

func (f Float64Data) Variance() (float64, error) {
	_ = "STUB: not implemented"

	// PopulationVariance finds the amount of variance within a population
	return 0, nil
}

func (f Float64Data) PopulationVariance() (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil

	// SampleVariance finds the amount of variance within a sample
}

func (f Float64Data) SampleVariance() (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Covariance is a measure of how much two sets of data change
		nil
}

func (f Float64Data) Covariance(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0,

		// CovariancePopulation computes covariance for entire population between two variables
		nil
}

func (f Float64Data) CovariancePopulation(d Float64Data) (float64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Sigmoid returns the input values along the sigmoid or s-shaped curve
func (f Float64Data) Sigmoid() ([]float64, error) {
	_ = "STUB: not implemented"

	// SoftMax returns the input values in the range of 0 to 1
	// with sum of all the probabilities being equal to one.
	return nil, nil
}

func (f Float64Data) SoftMax() ([]float64, error) {
	_ = "STUB: not implemented"

	// Entropy provides calculation of the entropy
	return nil, nil
}

func (f Float64Data) Entropy() (float64, error) {
	_ = "STUB: not implemented"

	// Quartiles returns the three quartile points from instance of Float64Data
	return 0, nil
}

func (f Float64Data) Quartiles() (Quartiles, error) {
	_ = "STUB: not implemented"
	return *new(Quartiles), nil
}

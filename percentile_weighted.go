package stats

// PercentileWeighted finds the weighted percentile of a slice of floats
// using the weighted empirical CDF (inverse CDF / nearest-rank method).
//
// For a given percent p, it returns the smallest data value x such that
// the cumulative weight of all values <= x is at least p% of the total
// weight. This matches the behavior of Python's statsmodels
// DescrStatsW.quantile.
//
// The data and weights slices must be the same length. Weights must be
// non-negative and at least one weight must be positive. The percent
// parameter must be between 0 and 100 (exclusive).
func PercentileWeighted(data, weights Float64Data, percent float64) (percentile float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Build sorted pairs by data value

// Find the smallest value where cumulative weight >= target

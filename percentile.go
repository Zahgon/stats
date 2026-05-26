package stats

// Percentile finds the relative standing in a slice of floats.
//
// The function uses the Linear Interpolation Between Closest Ranks method
// as recommended by NIST [1] and used by Excel (PERCENTILE), Google Sheets,
// NumPy (default), and other standard tools.
//
// Algorithm (for percent p and sorted data of length n):
//
//  1. Compute the rank: rank = (p / 100) * (n - 1)
//  2. Split into integer part k and fractional part f
//  3. Result = data[k] + f * (data[k+1] - data[k])
//
// [1] https://www.itl.nist.gov/div898/handbook/prc/section2/prc262.htm
func Percentile(input Float64Data, percent float64) (percentile float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Start by sorting a copy of the slice

// Use the standard linear interpolation method:
// rank = (percent / 100) * (n - 1)
// result = c[k] + f * (c[k+1] - c[k])

// PercentileNearestRank finds the relative standing in a slice of floats using the Nearest Rank method
func PercentileNearestRank(input Float64Data, percent float64) (percentile float64, err error) {
	_ = "STUB: not implemented"

	// Find the length of items in the slice
	return 0, nil
}

// Return an error for empty slices

// Return error for less than 0 or greater than 100 percentages

// Start by sorting a copy of the slice

// Return the last item

// Find ordinal ranking

// Return the item that is in the place of the ordinal rank

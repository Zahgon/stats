package stats

// Validate data for distance calculation
func validateData(dataPointX, dataPointY Float64Data) error { _ = "STUB: not implemented"; return nil }

// ChebyshevDistance computes the Chebyshev distance between two data sets
func ChebyshevDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// EuclideanDistance computes the Euclidean distance between two data sets
func EuclideanDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ManhattanDistance computes the Manhattan distance between two data sets
func ManhattanDistance(dataPointX, dataPointY Float64Data) (distance float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MinkowskiDistance computes the Minkowski distance between two data sets
//
// Arguments:
//
//	dataPointX: First set of data points
//	dataPointY: Second set of data points. Length of both data
//	            sets must be equal.
//	lambda:     aka p or city blocks; With lambda = 1
//	            returned distance is manhattan distance and
//	            lambda = 2; it is euclidean distance. Lambda
//	            reaching to infinite - distance would be chebysev
//	            distance.
//
// Return:
//
//	Distance or error
func MinkowskiDistance(dataPointX, dataPointY Float64Data, lambda float64) (distance float64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

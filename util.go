package stats

// float64ToInt rounds a float64 to an int
func float64ToInt(input float64) (output int) { _ = "STUB: not implemented"; return 0 }

// unixnano returns nanoseconds from UTC epoch
func unixnano() int64 { _ = "STUB: not implemented"; return 0 }

// copyslice copies a slice of float64s
func copyslice(input Float64Data) Float64Data { _ = "STUB: not implemented"; return *new(Float64Data) }

// sortedCopy returns a sorted copy of float64s
func sortedCopy(input Float64Data) (copy Float64Data) {
	_ = "STUB: not implemented"
	return *new(Float64Data)
}

// sortedCopyDif returns a sorted copy of float64s
// only if the original data isn't sorted.
// Only use this if returned slice won't be manipulated!
func sortedCopyDif(input Float64Data) (copy Float64Data) {
	_ = "STUB: not implemented"
	return *new(Float64Data)
}

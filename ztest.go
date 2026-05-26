package stats

// ZTest performs a one-sample or two-sample Z-test.
//
// For a one-sample Z-test, pass the sample data as data1, nil for data2,
// the known population mean as populationMean, and the known population
// standard deviation as populationStdDev.
//
// For a two-sample Z-test, pass both sample datasets and the known population
// standard deviations. The populationMean parameter is ignored in this case.
//
// Returns the Z statistic and the two-tailed p-value.
//
// https://en.wikipedia.org/wiki/Z-test
func ZTest(data1, data2 Float64Data, populationMean, populationStdDev float64) (z float64, pvalue float64, err error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}

// Two-sample Z-test

// One-sample Z-test

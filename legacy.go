package stats

// VarP is a shortcut to PopulationVariance
func VarP(input Float64Data) (sdev float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// VarS is a shortcut to SampleVariance
func VarS(input Float64Data) (sdev float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// StdDevP is a shortcut to StandardDeviationPopulation
func StdDevP(input Float64Data) (sdev float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// StdDevS is a shortcut to StandardDeviationSample
func StdDevS(input Float64Data) (sdev float64, err error) { _ = "STUB: not implemented"; return 0, nil }

// LinReg is a shortcut to LinearRegression
func LinReg(s []Coordinate) (regressions []Coordinate, err error) {
	_ = "STUB: not implemented"
	return nil,

		// ExpReg is a shortcut to ExponentialRegression
		nil
}

func ExpReg(s []Coordinate) (regressions []Coordinate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// LogReg is a shortcut to LogarithmicRegression
func LogReg(s []Coordinate) (regressions []Coordinate, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Legacy error names that didn't start with Err
var (
	EmptyInputErr = ErrEmptyInput
	NaNErr        = ErrNaN
	NegativeErr   = ErrNegative
	ZeroErr       = ErrZero
	BoundsErr     = ErrBounds
	SizeErr       = ErrSize
	InfValue      = ErrInfValue
	YCoordErr     = ErrYCoord
	EmptyInput    = ErrEmptyInput
)

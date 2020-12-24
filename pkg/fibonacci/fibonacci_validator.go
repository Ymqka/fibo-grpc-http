package fibo

// StopLesserThanStart stop < start
type stopLesserThanStart struct{}

func (s *stopLesserThanStart) Error() string { return "stop is higher than start" }

// TooSmallNumber uint64 can't hold fibonacci number higher than 90
type tooSmallNumber struct{}

func (s *tooSmallNumber) Error() string { return "cannot handle number higher than 90" }

func validateFiboRange(start, stop uint32) error {
	if stop < start {
		return new(stopLesserThanStart)
	}

	if stop > 90 {
		return new(tooSmallNumber)
	}

	return nil
}

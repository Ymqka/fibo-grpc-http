package fibo

// StopLesserThanStart stop < start
type stopLesserThanStart struct{}

func (s *stopLesserThanStart) Error() string { return "stop is higher than start" }

// TooSmallNumber uint64 can't hold fibonacci number higher than 90
type tooSmallNumber struct{}

func (s *tooSmallNumber) Error() string { return "cannot handle number higher than 90" }

type zeroStartIsNotAllowed struct {
}

func (s *zeroStartIsNotAllowed) Error() string {
	return "start should be >= 1 (first sequence idx = 1)"
}

func validateFiboRange(start, stop uint32) error {
	if stop < start {
		return new(stopLesserThanStart)
	}

	if stop > 90 {
		return new(tooSmallNumber)
	}

	if start == 0 {
		return new(zeroStartIsNotAllowed)
	}

	return nil
}

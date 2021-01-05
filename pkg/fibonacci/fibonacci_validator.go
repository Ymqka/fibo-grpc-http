package fibo

// StopLesserThanStart stop < start
type startHigherThanStop struct {
	start, stop uint32
}

func (s *startHigherThanStop) Error() string { return "start is higher than stop" }

// tooHigh prevent user asking for too big number
type tooHigh struct{}

func (s *tooHigh) Error() string {
	return "failed to make request, stop is higher than 10 000 and force flag is not used"
}

type lessThanZeroIsNotAllowed struct {
}

func (s *lessThanZeroIsNotAllowed) Error() string {
	return "less than zero is not allowed"
}

func validateFiboParams(p Params) error {

	if p.Stop < p.Start {
		return new(startHigherThanStop)
	}

	if p.Stop > 10000 && p.Force == false {
		return new(tooHigh)
	}

	if p.Start < 0 {
		return new(lessThanZeroIsNotAllowed)
	}

	return nil
}

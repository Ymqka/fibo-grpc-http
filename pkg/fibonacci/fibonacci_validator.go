package fibo

// StopLesserThanStart stop < start
type stopLesserThanStart struct{}

func (s *stopLesserThanStart) Error() string { return "stop is higher than start" }

// tooHigh prevent user asking for too big number
type tooHigh struct{}

func (s *tooHigh) Error() string {
	return `stop is higher than 10 000, to force request use force flag
http example: 
curl "http://localhost:10000/fibonacci?start=10009&stop=10010&force=1"
grpc client example:
&pb.FiboRangeRequest{Start: 1, Stop: 10, Force: true}`
}

type lessThanZeroIsNotAllowed struct {
}

func (s *lessThanZeroIsNotAllowed) Error() string {
	return "less than zero is not allowed"
}

func validateFiboParams(p Params) error {

	if p.Stop < p.Start {
		return new(stopLesserThanStart)
	}

	if p.Stop > 10000 && p.Force == false {
		return new(tooHigh)
	}

	if p.Start < 0 {
		return new(lessThanZeroIsNotAllowed)
	}

	return nil
}

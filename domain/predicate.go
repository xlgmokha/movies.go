package domain

type Predicate func(Movie) bool

func (self Predicate) Or(other Predicate) Predicate {
	return func(m Movie) bool {
		return self(m) || other(m)
	}
}

func (self Predicate) Not() Predicate {
	return func(m Movie) bool {
		return !self(m)
	}
}

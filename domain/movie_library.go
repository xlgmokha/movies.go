package domain

type MovieLibrary struct {
	Movies []Movie
}

//type Specification interface {
//IsSatisfiedBy(m Movie) bool
//}

func (self *MovieLibrary) Find(fn Predicate) *Movie {
	for i := range self.Movies {
		movie := self.Movies[i]
		if fn(movie) {
			return &movie
		}
	}

	return nil
}

func (self *MovieLibrary) FindAll(fn Predicate) []Movie {
	items := []Movie{}

	for i := range self.Movies {
		movie := self.Movies[i]

		if fn(movie) {
			items = append(items, movie)
		}
	}

	return items
}

func NewSpecification(p Predicate) Predicate {
	return p
}

func MovieProducedBy(studio string) Predicate {
	return func(m Movie) bool {
		return m.Studio.Name == studio
	}
}

func MovieReleasedAfter(year int) Predicate {
	return func(m Movie) bool {
		return m.Year > year
	}
}
func MovieReleasedBefore(year int) Predicate {
	return func(m Movie) bool {
		return m.Year < year
	}
}

func (self *MovieLibrary) FindAllMoviesByPixar() []Movie {
	return self.FindAll(MovieProducedBy("Pixar"))
}

func (self *MovieLibrary) FindAllMoviesByPixarOrDisney() []Movie {
	return self.FindAll(MovieProducedBy("Pixar").Or(MovieProducedBy("Disney")))
}

func (self *MovieLibrary) FindAllMoviesNotByPixar() []Movie {
	return self.FindAll(MovieProducedBy("Pixar").Not())
}

func (self *MovieLibrary) FindAllMoviesPublishedAfter2004() []Movie {
	return self.FindAll(MovieReleasedAfter(2004))
}

func (self *MovieLibrary) FindAllMoviesPublishedBetween1982And2003() []Movie {
	return self.FindAll(MovieReleasedAfter(1982).And(MovieReleasedBefore(2003)))
}

func (self *MovieLibrary) Add(movie Movie) {
	found := self.Find(func(x Movie) bool {
		return x.Equals(movie)
	})

	if found != nil {
		return
	}

	self.Movies = append(self.Movies, movie)
}

func (self *MovieLibrary) Count() int {
	return len(self.Movies)
}

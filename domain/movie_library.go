package domain

type MovieLibrary struct {
	Movies []Movie
}

func (self *MovieLibrary) Find(fn Predicate) *Movie {
	for i := range self.Movies {
		movie := self.Movies[i]
		if fn(movie) {
			return &movie
		}
	}

	return nil
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

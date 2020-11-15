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

func (self *MovieLibrary) FindAllMoviesByPixar() []Movie {
	return self.FindAll(func(x Movie) bool {
		return x.Studio.Name == "Pixar"
	})
}

func (self *MovieLibrary) FindAllMoviesByPixarOrDisney() []Movie {
	return self.FindAll(func(x Movie) bool {
		return x.Studio.Name == "Pixar" || x.Studio.Name == "Disney"
	})
}

func (self *MovieLibrary) FindAllMoviesNotByPixar() []Movie {
	return self.FindAll(func(x Movie) bool {
		return x.Studio.Name != "Pixar"
	})
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

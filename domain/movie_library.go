package domain

type MovieLibrary struct {
	Movies []Movie
}

func (self *MovieLibrary) Add(movie Movie) {
	self.Movies = append(self.Movies, movie)
}

func (self *MovieLibrary) Count() int {
	return len(self.Movies)
}

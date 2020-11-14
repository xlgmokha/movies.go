package domain

type Movie struct {
	Title  string
	Year   int
	Studio Studio
}

func (self *Movie) Equals(movie Movie) bool {
	return self.Title == movie.Title && self.Year == movie.Year && self.Studio.Equals(movie.Studio)
}

package testing

import "github.com/stretchr/testify/assert"
import "movies/domain"
import "testing"

func TestMovieLibrary(t *testing.T) {
	subject := domain.MovieLibrary{}

	pixar := domain.Studio{Name: "Pixar"}
	disney := domain.Studio{Name: "Disney"}
	castle_rock := domain.Studio{Name: "Castle Rock"}
	miramax_films := domain.Studio{Name: "Miramax"}
	regency_enterprises := domain.Studio{Name: "Regency"}

	shawshank_redemption := domain.Movie{Title: "The Shawshank Redemption", Year: 1994, Studio: castle_rock}
	chasing_amy := domain.Movie{Title: "Chasing Amy", Studio: miramax_films, Year: 1997}
	man_on_fire := domain.Movie{Title: "Man on Fire", Studio: regency_enterprises, Year: 2004}
	toy_story := domain.Movie{Title: "Toy Story", Studio: pixar, Year: 1995}
	up := domain.Movie{Title: "Up", Studio: pixar, Year: 2006}
	cars := domain.Movie{Title: "Cars", Studio: pixar, Year: 2009}
	monsters_inc := domain.Movie{Title: "Monsters Inc.", Studio: pixar, Year: 2001}
	fantasia := domain.Movie{Title: "Fantasia", Studio: disney, Year: 1940}
	dumbo := domain.Movie{Title: "Dumbo", Studio: disney, Year: 1941}
	pinocchio := domain.Movie{Title: "Pinocchio", Studio: disney, Year: 1940}

	t.Run("Add", func(t *testing.T) {
		t.Run("when adding a movie to the library", func(t *testing.T) {
			subject.Add(shawshank_redemption)
			subject.Add(chasing_amy)
			subject.Add(man_on_fire)
			subject.Add(toy_story)
			subject.Add(up)
			subject.Add(cars)
			subject.Add(monsters_inc)
			subject.Add(fantasia)
			subject.Add(dumbo)
			subject.Add(pinocchio)

			t.Run("increases the total number of movies in the library", func(t *testing.T) {
				assert.Equal(t, 10, subject.Count())
			})

			t.Run("does not allow duplicates", func(t *testing.T) {
				subject.Add(man_on_fire)

				assert.Equal(t, 10, subject.Count())
			})
		})
	})

	t.Run("Find", func(t *testing.T) {
		subject.Add(shawshank_redemption)
		subject.Add(chasing_amy)
		subject.Add(man_on_fire)
		subject.Add(toy_story)
		subject.Add(up)
		subject.Add(cars)
		subject.Add(monsters_inc)
		subject.Add(fantasia)
		subject.Add(dumbo)
		subject.Add(pinocchio)

		t.Run("returns all Pixar movies", func(t *testing.T) {
			movies := subject.FindAllMoviesByPixar()
			expected := [...]domain.Movie{toy_story, cars, up, monsters_inc}

			assert.ElementsMatch(t, expected, movies)
		})

		t.Run("returns all movies published by Pixar or Disney", func(t *testing.T) {
			movies := subject.FindAllMoviesByPixarOrDisney()
			expected := [...]domain.Movie{toy_story, cars, up, monsters_inc, fantasia, dumbo, pinocchio}

			assert.ElementsMatch(t, expected, movies)
		})

		t.Run("returns all movies not published by Pixar", func(t *testing.T) {
			movies := subject.FindAllMoviesNotByPixar()
			expected := [...]domain.Movie{
				chasing_amy,
				dumbo,
				fantasia,
				man_on_fire,
				pinocchio,
				shawshank_redemption,
			}

			assert.ElementsMatch(t, expected, movies)
		})

		t.Run("returns all movies released after 2004", func(t *testing.T) {
			movies := subject.FindAllMoviesPublishedAfter2004()
			expected := [...]domain.Movie{cars, up}

			assert.ElementsMatch(t, expected, movies)
		})
	})
}

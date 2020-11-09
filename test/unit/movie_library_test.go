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

    t.Run("Length", func(t *testing.T) {
      assert.Equal(t, 10, subject.Count())
    })
  })
}

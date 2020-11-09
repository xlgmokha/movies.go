package testing

import (
  "movies/domain"
  "testing"
)

func TestMovie(t *testing.T) {
  pixar := domain.Studio{ Name: "Pixar" }
  movie := domain.Movie{
    Title: "A Bugs Life",
    Year: 1998,
    Studio: pixar,
  }

  t.Run("Title", func(t *testing.T) {
    if movie.Title != "A Bugs Life" { t.Fatal(movie.Title) }
  })

  t.Run("Year", func(t *testing.T) {
    if movie.Year != 1998 { t.Fatal(movie.Year) }
  })

  t.Run("Studio", func(t *testing.T) {
    if movie.Studio.Name != "Pixar" { t.Fatal(movie.Studio.Name) }
  })
}

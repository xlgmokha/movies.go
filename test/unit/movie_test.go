package testing

import (
	"github.com/stretchr/testify/assert"
	"movies/domain"
	"testing"
)

func TestMovie(t *testing.T) {
	pixar := domain.Studio{Name: "Pixar"}
	movie := domain.Movie{
		Title:  "A Bugs Life",
		Year:   1998,
		Studio: pixar,
	}

	t.Run("Title", func(t *testing.T) {
		assert.Equal(t, "A Bugs Life", movie.Title)
	})

	t.Run("Year", func(t *testing.T) {
		assert.Equal(t, 1998, movie.Year)
	})

	t.Run("Studio", func(t *testing.T) {
		assert.Equal(t, "Pixar", movie.Studio.Name)
	})
}

package main

import (
	"netflix-go-htmx/components"
	"strconv"

	"github.com/labstack/echo/v4"
)

func MoviesPage(c echo.Context) error {

	popular, err := tmdbAPI.GetMoviePopular(nil)

	if err != nil {
		return err
	}

	trending, err := tmdbAPI.GetTrending("movie", "week")

	if err != nil {
		return err
	}

	return render(c, components.MoviesPage(popular, trending))

}

func MoviePage(c echo.Context) error {
	movieID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	movie, err := tmdbAPI.GetMovieInfo(movieID, nil)
	if err != nil {
		return err
	}

	return render(c, components.MoviePage(movie))

}

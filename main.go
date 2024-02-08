package main

import (
	"github.com/labstack/echo/v4"
	"github.com/ryanbradynd05/go-tmdb"
)

var tmdbAPI *tmdb.TMDb

func main() {
	config := tmdb.Config{
		APIKey:   "5e0d47f31b2b6168493f8a009d8a82be",
		Proxies:  nil,
		UseProxy: false,
	}

	tmdbAPI = tmdb.Init(config)

	e := echo.New()

	e.GET("/movies/", MoviesPage)
	e.GET("/movies/:id", MoviePage)

	e.Logger.Fatal(e.Start(":1323"))
}

package models

type TrendingResult struct {
	Adult              bool     `json:"adult,omitempty"`
	Gender             int      `json:"gender,omitempty"`
	BackdropPath       string   `json:"backdrop_path,omitempty"`
	GenreIDs           []int64  `json:"genre_ids,omitempty"`
	ID                 int64    `json:"id"`
	OriginalLanguage   string   `json:"original_language"`
	OriginalTitle      string   `json:"original_title,omitempty"`
	Overview           string   `json:"overview,omitempty"`
	PosterPath         string   `json:"poster_path,omitempty"`
	ReleaseDate        string   `json:"release_date,omitempty"`
	Title              string   `json:"title,omitempty"`
	Video              bool     `json:"video,omitempty"`
	VoteAverage        float32  `json:"vote_average,omitempty"`
	VoteCount          int64    `json:"vote_count,omitempty"`
	Popularity         float32  `json:"popularity,omitempty"`
	FirstAirDate       string   `json:"first_air_date,omitempty"`
	Name               string   `json:"name,omitempty"`
	OriginCountry      []string `json:"origin_country,omitempty"`
	OriginalName       string   `json:"original_name,omitempty"`
	KnownForDepartment string   `json:"known_for_department,omitempty"`
	ProfilePath        string   `json:"profile_path,omitempty"`
	KnownFor           []struct {
		Adult            bool    `json:"adult"`
		BackdropPath     string  `json:"backdrop_path"`
		GenreIds         []int   `json:"genre_ids"`
		ID               int     `json:"id"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		Overview         string  `json:"overview"`
		PosterPath       string  `json:"poster_path"`
		ReleaseDate      string  `json:"release_date"`
		Title            string  `json:"title"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		VoteCount        int     `json:"vote_count"`
		Popularity       float64 `json:"popularity"`
		MediaType        string  `json:"media_type"`
	} `json:"known_for,omitempty"`
}

type Movie struct {
	PosterPath string `json:"poster_path,omitempty"`
}

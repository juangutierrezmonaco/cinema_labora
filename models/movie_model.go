package models

type Genres struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type TMDBMovie struct {
	ID               int      `json:"id,omitempty"`
	Title            string   `json:"title,omitempty"`
	Adult            bool     `json:"adult,omitempty"`
	Genres           []Genres `json:"genres,omitempty"`
	ReleaseDate      string   `json:"release_date,omitempty"`
	PosterPath       string   `json:"poster_path,omitempty"`
	IMDbID           string   `json:"imdb_id,omitempty"`
	OriginalLanguage string   `json:"original_language,omitempty"`
	OriginalTitle    string   `json:"original_title,omitempty"`
	Overview         string   `json:"overview,omitempty"`
	Popularity       float64  `json:"popularity,omitempty"`
	Runtime          float64  `json:"runtime,omitempty"`
	Status           string   `json:"status,omitempty"`
	Tagline          string   `json:"tagline,omitempty"`
	ViewsCount       int      `json:"views_count"`
}

type TMDBResponse struct {
	Results []TMDBMovie `json:"results"`
}

package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/labora/labora-golang/cinema_labora/models"
	"github.com/labora/labora-golang/cinema_labora/util"
)

func GetTMDBMovieDetails(movieID int) (*models.TMDBMovie, error) {
	url := fmt.Sprintf("https://api.themoviedb.org/3/movie/%d?api_key=%s", movieID, util.EnvData.MovieDbData.ApiKey)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var tmdbMovie models.TMDBMovie
	err = json.NewDecoder(response.Body).Decode(&tmdbMovie)
	if err != nil {
		return nil, err
	}

	return &tmdbMovie, nil
}

func GetMoviesByTitleAndOverview(title, overview string) ([]models.TMDBMovie, error) {
	baseURL := "https://api.themoviedb.org/3/search/movie"
	apiKey := util.EnvData.MovieDbData.ApiKey
	queryParams := url.Values{}
	queryParams.Set("api_key", apiKey)
	queryParams.Set("query", title)
	queryParams.Set("overview", overview)

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryParams.Encode())

	// Realizar la solicitud GET a la API de TMDb
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON
	var tmdbResponse models.TMDBResponse
	err = json.NewDecoder(resp.Body).Decode(&tmdbResponse)
	if err != nil {
		return nil, err
	}

	// Mapear los resultados de TMDb a modelos de películas personalizados
	var movies []models.TMDBMovie
	for _, result := range tmdbResponse.Results {
		movie := models.TMDBMovie{
			Title:    result.Title,
			Overview: result.Overview,
			// Otros campos que desees incluir...
		}
		movies = append(movies, movie)
	}

	return movies, nil
}

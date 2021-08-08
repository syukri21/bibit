package fixtures

import "bibit-test/src/models/model"

var (
	MovieSearchResponse = model.MoviesSearchResponse{
		Search: []struct {
			Title  string `json:"Title"`
			Year   string `json:"Year"`
			ImdbID string `json:"imdbID"`
			Type   string `json:"Type"`
			Poster string `json:"Poster"`
		}{
			{
				Year:   string("2000"),
				Title:  string("Title"),
				ImdbID: string("tt1111"),
				Type:   string("Good"),
				Poster: string("poster"),
			},
		},
		TotalResults: 1,
		Response:     true,
	}
)

var (
	MoviesGetDetailResponse = model.MovieDetailResponse{
		Title:    "Mock",
		Year:     "Mock",
		Rated:    "Mock",
		Released: "Mock",
		Runtime:  "Mock",
		Genre:    "Mock",
		Director: "Mock",
		Writer:   "Mock",
		Actors:   "Mock",
		Plot:     "Mock",
		Language: "Mock",
		Country:  "Mock",
		Awards:   "Mock",
		Poster:   "Mock",
		Ratings: []struct {
			Source string `json:"Source"`
			Value  string `json:"Value"`
		}{
			{
				Source: "Mock",
				Value:  "Mock",
			},
		},
		Metascore:  "Mock",
		ImdbRating: "Mock",
		ImdbVotes:  "Mock",
		ImdbID:     "Mock",
		Type:       "Mock",
		DVD:        "Mock",
		BoxOffice:  "Mock",
		Production: "Mock",
		Website:    "Mock",
		Response:   true,
	}
)

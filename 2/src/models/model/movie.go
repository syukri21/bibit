package model

type MoviesSearchParams struct {
	Search string `json:"search" validate:"required"`
	Page   int    `json:"page" validate:"required"`
}

type MoviesGetDetailParams struct {
	ID string `json:"id" validate:"required"`
}

type MoviesSearchResponse struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults int64              `json:"totalResults,string"`
	Response     ConvertibleBoolean `json:"Response,string"`
}

type MovieDetailResponse struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string             `json:"Metascore"`
	ImdbRating string             `json:"imdbRating"`
	ImdbVotes  string             `json:"imdbVotes"`
	ImdbID     string             `json:"imdbID"`
	Type       string             `json:"Type"`
	DVD        string             `json:"DVD"`
	BoxOffice  string             `json:"BoxOffice"`
	Production string             `json:"Production"`
	Website    string             `json:"Website"`
	Response   ConvertibleBoolean `json:"Response"`
}

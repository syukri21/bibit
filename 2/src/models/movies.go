package models

type MoviesSearchParams struct {
	Search string
	Page   int
}

type MoviesGetDetailParams struct {
	ID string
}

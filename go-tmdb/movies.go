package tmdb



// Movie struct

// MovieShort struct
type MovieShort struct {
	Adult         bool
	BackdropPath  string `json:"backdrop_path"`
	ID            int
	OriginalTitle string `json:"original_title"`
	Popularity    float32
	PosterPath    string `json:"poster_path"`
	ReleaseDate   string `json:"release_date"`
	Title         string
	Video         bool
	VoteAverage   float32 `json:"vote_average"`
	VoteCount     uint32  `json:"vote_count"`
}

// MovieDatedResults struct


// MoviePagedResults struct
type MoviePagedResults struct {
	ID                int
	Page              int
	Results           []MovieShort
	TotalPages        int                     `json:"total_pages"`
	TotalResults      int                     `json:"total_results"`
	
}


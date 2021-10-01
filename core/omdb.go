package core

type OMDBSearchResult struct {
	Search       []OMDBResultCompact `json:"Search"`
	TotalResults string              `json:"totalResults"`
	Response     string              `json:"Response"`
}

type OMDBResultCompact struct {
	Title  string `json:"Title"`
	Year   string `json:"Year"`
	IMDBID string `json:"imdbID"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type OMDBResultSingle struct {
	Rated      string       `json:"Rated"`
	Released   string       `json:"Released"`
	Runtime    string       `json:"Runtime"`
	Genre      string       `json:"Genre"`
	Director   string       `json:"Director"`
	Writer     string       `json:"Writer"`
	Actors     string       `json:"Actors"`
	Plot       string       `json:"Plot"`
	Language   string       `json:"Language"`
	Country    string       `json:"Country"`
	Awards     string       `json:"Awards"`
	Ratings    []OMDBRating `json:"Ratings"`
	Metascore  string       `json:"Metascore"`
	IMDBRating string       `json:"imdbRating"`
	IMDBVotes  string       `json:"imdbVotes"`
	DVD        string       `json:"DVD"`
	BoxOffice  string       `json:"BoxOffice"`
	Production string       `json:"Production"`
	Website    string       `json:"Website"`
	Response   string       `json:"Response"`
	OMDBResultCompact
}

type OMDBRating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

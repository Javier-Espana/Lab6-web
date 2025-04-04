package api

type Series struct {
	ID               int    `json:"id"`
	Title            string `json:"title"`
	Status           string `json:"status"`
	LastEpisodeWatched int  `json:"lastEpisodeWatched"`
	TotalEpisodes    int    `json:"totalEpisodes"`
	Ranking          int    `json:"ranking"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
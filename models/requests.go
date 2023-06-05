package models

const MaxListLimit = 20

// GetGamesRequest for retrieving signle game
type GetGameRequest struct {
	ID string `json:"id"`
}

type ListGamesRequest struct {
	Limit int `json:"limit"`

	// Optional filters
	Opponent Opponent   `json:"opponent"`
	Status   GameStatus `json:"status"`
}

type CreateGameRequest struct {
	Game *Game `json:"game"`
}

type UpdateGameRequest struct {
	ID     string     `json:"id"`
	Status GameStatus `json:"status"`
}

package models

import "time"

type GameStatus string

// Some default game statuses
const (
	GameStatusCreated  GameStatus = "created"
	GameStatusStarted  GameStatus = "started"
	GameStatusFinished GameStatus = "finished"
)

type Game struct {
	// Identifier
	ID     string     `gorm:"primary_key" json:"id"`
	Status GameStatus `json:"status"`

	// Players
	Opponent Opponent `json:"opponent"`

	// Meta
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Opponent string

const (
	OpponentHuman Opponent = "human"
	OpponentModel Opponent = "model"
)

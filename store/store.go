package store

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go-tic-tac-toe-api/models"
)

type IGameStore interface {
	Get(ctx context.Context, req *models.GetGameRequest) (*models.Game, error)
	List(ctx context.Context, req *models.ListGamesRequest) ([]*models.Game, error)
	Create(ctx context.Context, req *models.CreateGameRequest) error
	Update(ctx context.Context, req *models.UpdateGameRequest) error
}

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// GenerateUniqueID will returns a time based sortable unique id
func GenerateUniqueID() string {
	word := []byte("0987654321")
	rand.Shuffle(len(word), func(i, j int) {
		word[i], word[j] = word[j], word[i]
	})
	now := time.Now().UTC()
	return fmt.Sprintf("%010v-%010v-%s", now.Unix(), now.Nanosecond(), string(word))
}

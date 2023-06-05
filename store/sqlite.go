package store

import (
	"context"
	"go-tic-tac-toe-api/errors"
	"go-tic-tac-toe-api/models"
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlt struct {
	db *gorm.DB
}

func NewSqliteGameStore(conn string) IGameStore {
	db, err := gorm.Open(sqlite.Open(conn),
		&gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "", log.LstdFlags),
				logger.Config{
					LogLevel: logger.Info,
					Colorful: true,
				},
			),
		},
	)

	if err != nil {
		panic("Unable to connect to database: " + err.Error())
	}

	if err := db.AutoMigrate(&models.Game{}); err != nil {
		panic("Unable to migrate database: " + err.Error())
	}

	return &sqlt{db: db}
}

func (s *sqlt) Create(ctx context.Context, in *models.CreateGameRequest) error {
	if in.Game == nil {
		return errors.ErrObjectIsRequired
	}

	in.Game.ID = GenerateUniqueID()
	// in.Game.Opponent = in.
	in.Game.Status = models.GameStatusCreated
	in.Game.CreatedAt = s.db.NowFunc()

	return s.db.WithContext(ctx).
		Create(in.Game).
		Error
}

func (s *sqlt) Get(ctx context.Context, in *models.GetGameRequest) (*models.Game, error) {
	game := &models.Game{}

	// take game where id == uid from database
	err := s.db.WithContext(ctx).Take(game, "id = ?", in.ID).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.ErrGameNotFoundRequest
	}

	return game, err
}

func (s *sqlt) List(ctx context.Context, in *models.ListGamesRequest) ([]*models.Game, error) {
	if in.Limit == 0 || in.Limit > models.MaxListLimit {
		in.Limit = models.MaxListLimit
	}
	query := s.db.WithContext(ctx).Limit(in.Limit)
	list := make([]*models.Game, 0, in.Limit)
	err := query.Order("id").Find(&list).Error

	return list, err
}

func (s *sqlt) Update(ctx context.Context, in *models.UpdateGameRequest) error {
	game := &models.Game{
		ID:        in.ID,
		Status:    in.Status,
		UpdatedAt: s.db.NowFunc(),
	}

	return s.db.WithContext(ctx).Model(game).
		Select("status", "updated_at").
		Updates(game).
		Error
}

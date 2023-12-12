package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/max-bazarov/chat/internal/database/postgres"
	"github.com/max-bazarov/chat/internal/database/redis"
	"github.com/max-bazarov/chat/internal/models"
	r "github.com/redis/go-redis/v9"
)

type Authorization interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type Chat interface {
	WriteMessage()
	ReadMessage(hub *models.Hub)
}

type Repository struct {
	Authorization
	Chat
}

func NewRepository(db *sqlx.DB, rdb *r.Client) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthRepo(db),
		Chat:          redis.NewChatRepo(rdb),
	}
}

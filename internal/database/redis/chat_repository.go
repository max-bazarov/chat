package redis

import (
	"github.com/max-bazarov/chat/internal/models"
	"github.com/redis/go-redis/v9"
)

type ChatRepo struct {
	redis *redis.Client
}

func NewChatRepo(db *redis.Client) *ChatRepo {
	return &ChatRepo{redis: db}
}

func (r *ChatRepo) WriteMessage() {

}

func (r *ChatRepo) ReadMessage(hub *models.Hub) {
}

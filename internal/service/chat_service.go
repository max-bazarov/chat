package service

import (
	"github.com/max-bazarov/chat/internal/database/redis"
	"github.com/max-bazarov/chat/internal/models"
)

type ChatService struct {
	repo redis.ChatRepo
}

func NewChatService(repo redis.ChatRepo) *ChatService {
	return &ChatService{
		repo: repo,
	}
}

func (s *ChatService) WriteMessage() {
	// defer func() {
	// 	c.Conn.Close()
	// }()

	// for {
	// 	message, ok := <-c.Message
	// 	if !ok {
	// 		return
	// 	}

	// 	c.Conn.WriteJSON(message)
	// }

}

func (s *ChatService) ReadMessage(hub *models.Hub) {

	// for {
	// 	_, m, err := c.Conn.ReadMessage()
	// 	if err != nil {
	// 		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
	// 			log.Printf("error: %v", err)
	// 		}
	// 		break
	// 	}
	// 	defer func() {
	// 		hub.Unregister <- c
	// 		c.Conn.Close()
	// 	}()

	// 	msg := &models.Message{
	// 		Content:  string(m),
	// 		RoomID:   c.RoomID,
	// 		Username: c.Username,
	// 	}

	// 	hub.Broadcast <- msg
	// }
}

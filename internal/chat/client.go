package chat

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/max-bazarov/chat/internal/models"
)

func (c *models.Client) WriteMessage() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		message, ok := <-c.Message
		if !ok {
			return
		}

		c.Conn.WriteJSON(message)
	}
}

func (c *models.Client) ReadMessage(hub *Hub) {

	for {
		_, m, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		defer func() {
			hub.Unregister <- c
			c.Conn.Close()
		}()

		msg := &models.Message{
			Content:  string(m),
			RoomID:   c.RoomID,
			Username: c.Username,
		}

		hub.Broadcast <- msg
	}
}

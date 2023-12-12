package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/max-bazarov/chat/internal/chat"
	"github.com/max-bazarov/chat/internal/models"
)

type Handler struct {
	hub *models.Hub
}

func NewHandler(h *models.Hub) *Handler {
	return &Handler{
		h,
	}
}

type CreateRoomReq struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c echo.Context) error {
	req := new(CreateRoomReq)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	h.hub.Rooms[req.ID] = &models.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*models.Client),
	}

	return c.JSON(http.StatusOK, req)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c echo.Context) error {
	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	roomID := c.Param("roomId")
	clientID := c.QueryParam("userId")
	username := c.QueryParam("username")

	cl := &models.Client{
		Conn:     conn,
		Message:  make(chan *models.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &models.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- cl
	h.hub.Broadcast <- m

	go cl.WriteMessage()
	cl.ReadMessage(h.hub)
	return nil
}

type RoomRes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c echo.Context) error {
	rooms := make([]RoomRes, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	return c.JSON(http.StatusOK, rooms)
}

type ClientRes struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c echo.Context) error {
	var clients []ClientRes
	roomId := c.Param("roomId")

	if _, ok := h.hub.Rooms[roomId]; !ok {
		clients = make([]ClientRes, 0)
		c.JSON(http.StatusOK, clients)
	}

	for _, c := range h.hub.Rooms[roomId].Clients {
		clients = append(clients, ClientRes{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	return c.JSON(http.StatusOK, clients)
}

package rest

import (
	"github.com/labstack/echo/v4"
	"github.com/max-bazarov/chat/internal/transport/ws"
)

func InitRoutes(authHandler *Handler, wsHandler *ws.Handler) (e *echo.Echo) {
	e = echo.New()

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.GET("/logout", authHandler.Logout)

	e.POST("/rooms", wsHandler.CreateRoom)
	e.GET("/rooms", wsHandler.GetRooms)
	e.GET("/chat/:roomId", wsHandler.JoinRoom)
	e.GET("/chat/:roomId/clients", wsHandler.GetClients)

	return
}

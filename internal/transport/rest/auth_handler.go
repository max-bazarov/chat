package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/max-bazarov/chat/internal/models"
	"github.com/max-bazarov/chat/internal/service"
)

type Handler struct {
	*service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) Register(c echo.Context) error {
	var u models.CreateUserReq
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := h.Service.Register(c.Request().Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) Login(c echo.Context) error {
	var user models.LoginUserReq
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	u, err := h.Service.Login(c.Request().Context(), &user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	cookie := http.Cookie{
		Name:     "jwt",
		Value:    u.AccessToken,
		Expires:  time.Now().Add(time.Second * 60 * 60 * 24),
		Path:     "/",
		Domain:   "localhost",
		Secure:   false,
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) Logout(c echo.Context) error {
	cookie := http.Cookie{
		Name:     "jwt",
		Value:    "",
		MaxAge:   -1,
		Path:     "",
		Domain:   "",
		Secure:   false,
		HttpOnly: true,
	}
	c.SetCookie(&cookie)
	return c.JSON(http.StatusOK, gin.H{"message": "logout successful"})
}

package service

import (
	"context"

	"github.com/max-bazarov/chat/internal/database"
	"github.com/max-bazarov/chat/internal/models"
)

type Authorization interface {
	Register(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error)
	Login(c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error)
}

type Service struct {
	Authorization
}

func NewService(repo *database.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}

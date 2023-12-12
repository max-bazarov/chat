package service

import (
	"context"

	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/max-bazarov/chat/internal/database"
	"github.com/max-bazarov/chat/internal/models"
	"github.com/max-bazarov/chat/internal/utils"
)

const (
	secretKey = "secret"
)

type AuthService struct {
	repo    database.Authorization
	timeout time.Duration
}

func NewAuthService(repo database.Authorization) *AuthService {
	return &AuthService{
		repo:    repo,
		timeout: time.Duration(2) * time.Second,
	}
}

func (s *AuthService) Register(c context.Context, req *models.CreateUserReq) (*models.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &models.CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *AuthService) Login(c context.Context, req *models.LoginUserReq) (*models.LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	u, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &models.LoginUserRes{}, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &models.LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &models.LoginUserRes{}, err
	}

	return &models.LoginUserRes{AccessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

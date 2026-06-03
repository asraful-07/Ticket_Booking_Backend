package user

import (
	"tickets/internal/auth"
	"tickets/internal/user/dto"
)

type service struct {
	repo Repository
	jwtService  auth.JWTService
}

func NewService(repo Repository, jwtService auth.JWTService) *service {
	return &service{repo: repo, jwtService: jwtService}
}

func (s *service) CreateUser(req dto.CreateUserRequest) (*dto.Response, error) {
	user := User{
		Email: req.Email,
		Name:  req.Name,
	}

	// hash password and set to user.Password
	err := user.hashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	err = s.repo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

    response := dto.Response{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.String(),
	}

	return &response, nil
}

func (s *service) LoginUser(req dto.LoginRequest) (*dto.Response, error) {
user, err := s.repo.GetUserByEmail(req.Email)
if err != nil {
	return nil, err
}

if user == nil || !user.CheckPassword(req.Password) {
	return nil, err
}

// token generation logic can be added here
token, err := s.jwtService.GenerateToken(user.ID, user.Name, user.Email)
if err != nil {
	return nil, err
}

response := dto.Response{
	Id:        user.ID,
	Name:      user.Name,
	Email:     user.Email,
	Token:     token,
	CreatedAt: user.CreatedAt.String(),
}

return &response, nil
}

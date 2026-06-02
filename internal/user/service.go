package user

import "tickets/internal/user/dto"

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(res dto.CreateUserRequest) (*dto.Response, error) {
	user := User{
		Email: res.Email,
		Name:  res.Name,
	}

	err := s.repo.CreateUser(&user)
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
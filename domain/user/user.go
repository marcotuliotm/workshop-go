package user

import "workshop-go/domain"

// User usario
type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (*Service) IsValid(user *domain.User) bool {
	return user.Age > 18
}

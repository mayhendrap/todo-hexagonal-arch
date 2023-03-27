package service

import (
	"golang.org/x/crypto/bcrypt"
	"todo-hexagonal-arch/internal/core/domain"
	"todo-hexagonal-arch/internal/core/interfaces"
)

type authService struct {
	userRepository interfaces.UserRepository
	tokenUtil      interfaces.TokenUtil
}

func NewAuthService(userRepository interfaces.UserRepository, tokenUtil interfaces.TokenUtil) interfaces.AuthService {
	return &authService{
		userRepository: userRepository,
		tokenUtil:      tokenUtil,
	}
}

func (as *authService) Register(user domain.User) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user.Password = string(hashedPassword)

	userID, err := as.userRepository.Create(user)
	if err != nil {
		return "", err
	}

	token, err := as.tokenUtil.GenerateToken(userID, user.Email)
	return token, nil
}

func (as *authService) Login(email, password string) (string, error) {
	user, err := as.userRepository.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	token, err := as.tokenUtil.GenerateToken(user.ID, user.Email)
	return token, err
}

package user

import (
	"errors"
	"log/slog"

	"gorm.io/gorm"
)

var logger = slog.Default()

type Service struct {
	repo UserRepository
}

type UserService interface {
	FinalizeSerie(seriesID string) error
}

func NewService(repo UserRepository) Service {
	return Service{repo: repo}
}

func (us Service) CreateUser(entity User) (User, error) {
	_, err := us.repo.CreateUser(entity)
	if err != nil {
		logger.Error("error creating user", err)
		return User{}, err
	}
	return User{}, nil
}

func (us Service) GetUserById(ChatID uint64) (User, error) {
	user, err := us.repo.GetUserByID(ChatID)
	if err != nil {
		logger.Error("error while getting user by id: ", err)
		return User{}, err
	}
	return user, nil
}

func (us Service) GetByUsername(username string) (User, error) {
	user, err := us.repo.GetByUsername(username)
	if err != nil {
		logger.Error("error while getting user by username: ", err)
		return User{}, err
	}
	return user, nil
}

func (us Service) IsNew(chatID int64) bool {
	user, err := us.repo.GetByChatID(chatID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return true
		}
		return false
	}
	return IsNil(user)
}

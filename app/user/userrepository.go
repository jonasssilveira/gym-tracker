package user

import "gorm.io/gorm"

type UserRepository struct {
	db gorm.DB
}

type Repository interface {
	GetUserByID(id uint64) (User, error)
	CreateUser(user User) (User, error)
	GetByUsername(username string) (User, error)
	GetByChatID(chatID int64) (User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: *db}
}

func (r UserRepository) GetUserByID(userID uint64) (User, error) {
	var user User
	if err := r.db.First(&user, userID).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r UserRepository) GetByUsername(username string) (User, error) {
	var user User
	if err := r.db.First(&user, username).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r UserRepository) GetByChatID(chatID int64) (User, error) {
	var user User
	if err := r.db.First(&user, chatID).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r UserRepository) CreateUser(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

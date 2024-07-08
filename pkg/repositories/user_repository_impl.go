package repositories

import (
	"gocrud/pkg/cache"
	"gocrud/pkg/models"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUser(userID int) (*models.User, error) {
	user := &models.User{UserID: userID}
	repoFunc := func() (interface{}, error) {
		var fetchedUser models.User
		if err := r.db.First(&fetchedUser, userID).Error; err != nil {
			return nil, err
		}
		return &fetchedUser, nil
	}

	result, err := cache.ReadFromCache("user:%user_id%:age:%age%", user, repoFunc)
	if err != nil {
		return nil, err
	}

	return result.(*models.User), nil
}

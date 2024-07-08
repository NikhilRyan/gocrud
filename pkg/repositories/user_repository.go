package repositories

import "gocrud/pkg/models"

type UserRepository interface {
	GetUser(userID int) (*models.User, error)
}

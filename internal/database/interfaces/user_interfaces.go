package interfaces

import (
	"context"
	"shafra-task-1/internal/models"
)

type IUserInterface interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUserById(ctx context.Context, id int64) (*models.User, error)
	UpdateUser(ctx context.Context, id int64, nameSurname string, age int) (*models.User, error)
	DeleteUser(ctx context.Context, id int64) error
}

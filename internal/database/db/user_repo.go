package db

import (
	"context"
	"errors"
	"shafra-task-1/internal/database/interfaces"
	"shafra-task-1/internal/models"
	"shafra-task-1/utils"

	"gorm.io/gorm"
)

type PostgresDB struct {
	DB *gorm.DB
}

func NewUserRepo(conn *gorm.DB) interfaces.IUserInterface {
	return &PostgresDB{DB: conn}
}

func (p *PostgresDB) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	if user == nil {
		return &models.User{}, errors.New("user details empty")
	}

	err := p.DB.WithContext(ctx).Model(&models.User{}).Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return user, nil
}

func (p *PostgresDB) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	user := &models.User{}
	err := p.DB.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return &models.User{}, err
	}

	userFromDb := &models.User{
		ID:          utils.RandomInt(1, 1000),
		NameSurname: user.NameSurname,
		Age:         user.Age,
	}

	return userFromDb, nil
}

func (p *PostgresDB) UpdateUser(ctx context.Context, id int64, nameSurname string, age int) (*models.User, error) {
	var userForUpdate = &models.User{}

	if err := p.DB.WithContext(ctx).Model(userForUpdate).Where("id = ?", id).
		Updates(map[string]interface{}{"name_surname": nameSurname, "age": age}).Error; err != nil {
		return &models.User{}, err
	}
	return userForUpdate, nil
}

func (p *PostgresDB) DeleteUser(ctx context.Context, id int64) error {
	var err error

	tx := p.DB.Begin()

	delTx := tx.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Delete(&models.User{})

	if err = delTx.Error; err != nil {
		return err
	} else {
		tx.Commit()
	}

	return nil
}

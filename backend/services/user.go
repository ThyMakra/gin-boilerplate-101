package services

import (
	"github.com/ThyMakra/gin-boilerplate/backend/models"
	"github.com/ThyMakra/gin-boilerplate/backend/schemas"
)

type UserEntity interface {
	RegisterEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError)
	LoginEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError)
}

type userService struct {
	user UserEntity
}

func NewUserService(user UserEntity) *userService {
	return &userService{user: user}
}

func (s *userService) RegisterEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError) {
	var schema schemas.UserSchema
	schema.FirstName = input.FirstName
	schema.LastName = input.LastName
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.RegisterEntity(&schema)
	return res, err
}

func (s *userService) LoginEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError) {
	var schema schemas.UserSchema
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.user.LoginEntity(&schema)
	return res, err
}

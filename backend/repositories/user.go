package repositories

import (
	"net/http"

	"github.com/ThyMakra/gin-boilerplate/backend/constant"
	"github.com/ThyMakra/gin-boilerplate/backend/models"
	"github.com/ThyMakra/gin-boilerplate/backend/schemas"
	"github.com/ThyMakra/gin-boilerplate/pkg/utils"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError) {
	var user models.UserModel
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)
	if checkEmailExist.RowsAffected > 0 {
		err <- schemas.SchemaDatabaseError{
			Code:    http.StatusConflict,
			Message: constant.ErrorRegisterUserConflict,
		}
		return &user, <-err
	}

	addNewUser := db.Debug().Create(&user).Commit()

	if addNewUser.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code:    http.StatusForbidden,
			Message: constant.ErrorRegisterUserFailed,
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

func (r *userRepository) LoginEntity(input *schemas.UserSchema) (*models.UserModel, schemas.SchemaDatabaseError) {
	var user models.UserModel
	user.Email = input.Email
	user.Password = input.Password

	err := make(chan schemas.SchemaDatabaseError, 1)

	db := r.db.Model(&user)

	checkEmailExist := db.Debug().First(&user, "email = ?", input.Email)
	if checkEmailExist.RowsAffected < 1 {
		err <- schemas.SchemaDatabaseError{
			Code:    http.StatusNotFound,
			Message: constant.ErrorLoginNotFound,
		}
		return &user, <-err
	}

	checkPasswordMatch := utils.ComparePassword(user.Password, input.Password)
	if checkPasswordMatch != nil {
		err <- schemas.SchemaDatabaseError{
			Code:    http.StatusBadRequest,
			Message: constant.ErrorLoginIncorrect,
		}
		return &user, <-err
	}

	err <- schemas.SchemaDatabaseError{}
	return &user, <-err
}

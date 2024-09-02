package models

import (
	"time"

	"github.com/ThyMakra/gin-boilerplate/pkg/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID        string    `json:"id" gorm:"primary_key"`
	FirstName string    `json:"first_name" gorm:"type:varchar;  not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar; not null"`
	Email     string    `json:"email" gorm:"type:varchar; unique; not null"`
	Password  string    `json:"password" gorm:"type:varchar; not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *UserModel) BeforeCreate(db *gorm.DB) error {
	m.ID = uuid.NewString()
	m.Password = utils.HashPassword(m.Password)
	m.CreatedAt = time.Now()
	return nil
}

func (m *UserModel) BeforeUpdate(db *gorm.DB) error {
	// m.ID = uuid.NewString
	m.Password = utils.HashPassword(m.Password)
	m.UpdatedAt = time.Now()
	return nil
}

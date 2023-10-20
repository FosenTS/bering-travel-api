package dto

import (
	"github.com/achillescres/pkg/object/oid"
	"saina.gitlab.yandexcloud.net/saina/backend/saina-api/internal/domain/entity"
	"time"
)

type UserCreate struct {
	Create `json:"-" db:"-" binding:"-"`

	Login          string    `json:"login" db:"login" binding:"required,max=25,min=5"`
	Password       string    `json:"password" db:"password" binding:"required,min=8,max=50"`
	RegisteredTime time.Time `json:"registeredTime" db:"registered_time" binding:"required"`
	FullName       string    `json:"fullName" db:"full_name" binding:"required,max=100"`
	Position       string    `json:"position" db:"position" binding:"required,max=50"`
	Email          string    `json:"email" db:"email" binding:"required,email,max=200"`
	PhoneNumber    string    `json:"phoneNumber" db:"phone_number" binding:"required,e164"`
	Role           string    `json:"role" db:"role" binding:"required,max=50"`
	AirlineCode    string    `json:"airlineCode" db:"airline_code" binding:"required,max=25"`
	Approved       bool      `json:"approved" db:"approved" binding:"required"`
}

func NewUserCreate(login string, password string, registeredTime time.Time, fullName string, position string, email string, phoneNumber string, role string, airlineCode string, approved bool) *UserCreate {
	return &UserCreate{Login: login, Password: password, RegisteredTime: registeredTime, FullName: fullName, Position: position, Email: email, PhoneNumber: phoneNumber, Role: role, AirlineCode: airlineCode, Approved: approved}
}

func (uC *UserCreate) ToEntity(id oid.ID, hashedPassword string) *entity.User {
	return entity.NewUser(
		id,
		uC.Login,
		hashedPassword,
		uC.RegisteredTime,
		uC.FullName,
		uC.Position,
		uC.Email,
		uC.PhoneNumber,
		uC.Role,
		uC.AirlineCode,
		uC.Approved,
	)
}

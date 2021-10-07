package user

import (
	"miniproject/business/user"
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey"`
	Username  string `gorm:"unique"`
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	DOB       string
	Gender    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func toDomain(users User) user.Domain {
	return user.Domain{
		ID:        users.ID,
		Username:  users.Username,
		FirstName: users.FirstName,
		LastName:  users.LastName,
		Email:     users.Email,
		Password:  users.Password,
		DOB:       users.DOB,
		Gender:    users.Gender,
		CreatedAt: users.CreatedAt,
		UpdatedAt: users.UpdatedAt,
	}
}

func fromDomain(domain user.Domain) User {
	return User{
		ID:        domain.ID,
		Username:  domain.Username,
		FirstName: domain.FirstName,
		LastName:  domain.LastName,
		Email:     domain.Email,
		Password:  domain.Password,
		DOB:       domain.DOB,
		Gender:    domain.Gender,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

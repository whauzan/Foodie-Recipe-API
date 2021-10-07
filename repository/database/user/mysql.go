package user

import (
	"miniproject/business/user"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) user.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (rep *MysqlUserRepository) Login(username string, password string) (user.Domain, error) {
	var users User
	result := rep.Conn.First(&users, "username = ?", username)

	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	return toDomain(users), nil
}

func (rep *MysqlUserRepository) Register(domain *user.Domain) (user.Domain, error) {
	users := fromDomain(*domain)
	result := rep.Conn.Create(&users)

	if result.Error != nil {
		return user.Domain{}, result.Error
	}

	return toDomain(users), nil
}

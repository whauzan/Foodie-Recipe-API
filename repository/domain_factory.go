package repository

import (
	"gorm.io/gorm"
	userDomain "miniproject/business/user"
	userDB "miniproject/repository/database/user"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMysqlUserRepository(conn)
}
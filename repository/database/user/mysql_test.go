package user_test

import (
	userBusiness "miniproject/business/user"
	"miniproject/repository/database/user"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var users = userBusiness.Domain{
	ID          : 1,
    Username    : "testUser",
	FirstName	: "Wahyu",
	LastName	: "Rafi",
	Email       : "test@gmail.com",
    Password    : "testPassword",
    DOB			: "08 Agustus 2000",
	Gender		: "Male",
    CreatedAt   : time.Now(),
    UpdatedAt   : time.Now(),
}

func TestInsert(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gdb, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	userRepo := user.NewMysqlUserRepository(gdb)
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`email`,`firstname`, `lastname`, `date_of_birth`, `gender` ,`id`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?)").
		WithArgs(users.CreatedAt, users.UpdatedAt, nil, users.Username, users.Password, users.Email, users.FirstName, users.LastName, users.DOB, users.Gender ,users.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	_, err = userRepo.Register(&users)
	require.NoError(t, err)
}
package user_test

import (
	"miniproject/app/middleware"
	"miniproject/business"
	"miniproject/business/user"
	userMock "miniproject/business/user/mock"
	"miniproject/helper/encrypt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userMockRepo userMock.Repository
	jwtAuth      *middleware.ConfigJWT
	userService  user.Service
)

func TestMain(m *testing.M) {
	jwtAuth := &middleware.ConfigJWT{
		SecretJWT:       "testmock123",
		ExpiresDuration: 1,
	}
	userService = user.NewUserService(&userMockRepo, 2, jwtAuth)

	m.Run()
}

// go test -v ./...
func TestLogin(t *testing.T) {
	t.Run("Test case 1 | Valid login", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "user123",
			Password: "wahyu123",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		resp, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)
	})
	t.Run("Test case 2 | Password Invalid", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "nama123",
			Password: "wahyuwrong",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrUser)
	})

	t.Run("Test case 3 | Empty Password", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "nama123",
			Password: "",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})
	t.Run("Test case 4 | Empty Username", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "",
			Password: encryptedPass,
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Invalid Req", func(t *testing.T) {
		// encryptedPass, _ := encrypt.HashPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{}

		expectedReturnService := user.Domain{
			Username: "nama",
			Password: "pass123",
		}

		userMockRepo.On("Login", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, business.ErrUser).Once()

		_, err := userService.Login(expectedReturnService.Username, expectedReturnService.Password)
		assert.NotNil(t, err)
		assert.NotEqual(t, UserDomain.Username, expectedReturnService.Username)
	})
}

func TestRegister(t *testing.T) {
	t.Run("Test case 1 | Valid Register", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "user123",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.Anything).Return(UserDomain, nil).Once()

		resp, err := userService.Register(&expectedReturnService)
		assert.Nil(t, err)
		assert.NotEmpty(t, resp)

	})

	t.Run("Test case 2 | Empty Username", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "",
			Email:    "nama@email123.com",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
		// err == business.ErrPassword == true; ok pass

	})

	t.Run("Test case 3 | Empty Email", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "nama123",
			Email:    "",
			Password: encryptedPass,
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})

	t.Run("Test case 4 | Empty Password", func(t *testing.T) {
		encryptedPass, _ := encrypt.HashAndPassword("wahyu123")
		// anggap si user domain database
		UserDomain := user.Domain{
			ID:        1,
			Username:  "nama123",
			Email:     "nama@email123.com",
			Password:  encryptedPass,
			Token:     "token123",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		expectedReturnService := user.Domain{
			Username: "nama123",
			Email:    "nama@email123.com",
			Password: "",
		}

		userMockRepo.On("Register", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(UserDomain, nil).Once()

		_, err := userService.Register(&expectedReturnService)
		assert.Equal(t, err, business.ErrEmptyForm)
	})
}
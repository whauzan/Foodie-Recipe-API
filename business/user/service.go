package user

import (
	"miniproject/app/middleware"
	"miniproject/business"
	"miniproject/helper/encrypt"
	"time"
)

type UserService struct {
	repository     Repository
	contextTimeout time.Duration
	jwtAuth        *middleware.ConfigJWT
}

func NewUserService(repo Repository, timeout time.Duration, jwtauth *middleware.ConfigJWT) Service {
	return &UserService{
		repository:     repo,
		contextTimeout: timeout,
		jwtAuth:        jwtauth,
	}
}

func (service *UserService) Register(domain *Domain) (Domain, error) {
	if domain.Username == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Email == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if domain.Password == "" {
		return Domain{}, business.ErrEmptyForm
	}
	encryptPass, _ := encrypt.HashAndPassword(domain.Password)
	domain.Password = encryptPass
	usr, err := service.repository.Register(domain)

	if err != nil {
		return Domain{}, err
	}
	return usr, nil
}

func (service *UserService) Login(usrname, passwd string) (Domain, error) {
	if usrname == "" {
		return Domain{}, business.ErrEmptyForm
	}
	if passwd == "" {
		return Domain{}, business.ErrEmptyForm
	}

	usr, err := service.repository.Login(usrname, passwd)
	if err != nil {
		return Domain{}, err
	}
	valid := encrypt.CheckPasswordHash(passwd, usr.Password)
	if !valid {
		return Domain{}, business.ErrUser
	}
	usr.Token = service.jwtAuth.GenerateToken(usr.ID, "user")
	return usr, nil
}

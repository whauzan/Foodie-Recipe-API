package user

import "time"

type Domain struct {
	ID        int
	Username  string
	FirstName string
	LastName  string
	Email     string
	Password  string
	DOB       string
	Gender    string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username string, password string) (Domain, error)
}

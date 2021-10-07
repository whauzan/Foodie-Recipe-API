package business

import "errors"

var (
	ErrInternalServer = errors.New("Our service can't reach you :( Please contact our administrator")
	ErrNotFound       = errors.New("Data not found :(")
	ErrIDNotFound     = errors.New("ID not found :(")
	ErrDuplicateData  = errors.New("There is something duplicated")
	ErrUser           = errors.New("Ara, username or password wrong")
	ErrEmailNotFound  = errors.New("Email not found")
	ErrPasswdNotFound = errors.New("Ara ara, password not found")
	ErrwrongPasswd    = errors.New("Ara ara ara, password incorrect")
	ErrEmptyForm      = errors.New("please, fill all required form")
)

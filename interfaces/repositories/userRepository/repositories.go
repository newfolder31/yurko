package userRepository

import (
	"github.com/jmoiron/sqlx"
	"github.com/newfolder31/yurko/usecases/userUsecases"
)

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *userUsecases.User
}

type SqlxHandler interface {
	Execute(query string, domain interface{}) error
	NamedExec(query string, data map[string]interface{}) error
	Get(domain interface{}, query string, args ...interface{}) error
	Select(slice []interface{}, query string) error
	QueryRowx(domain interface{}, query string, args ...interface{}) *sqlx.Row
}

package userInterfaces

import (
	"usecases/user"
)

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *userUsecases.User
}

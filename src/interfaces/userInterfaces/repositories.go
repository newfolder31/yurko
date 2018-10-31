package userInterfaces

import (
	"usecases/userUsecases"
)

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *userUsecases.User
}

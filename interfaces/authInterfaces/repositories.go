package authInterfaces

import (
	"yurko/usecases/authUsecases"
)

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *authUsecases.User
}

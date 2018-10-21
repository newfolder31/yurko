package authInterfaces

import (
	"usecases/authUsecases"
)

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *authUsecases.User
}

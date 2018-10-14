package interfaces

import "yurko/usecases"

type UserRepo interface {
	FindByEmailAndPassword(email, pass string) *usecases.User
}

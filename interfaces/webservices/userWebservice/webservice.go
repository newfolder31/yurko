package userWebservice

import (
	"github.com/newfolder31/yurko/usecases/userUsecases"
)

type RegistrationInteractor interface {
	Registration(form *userUsecases.RegistrationForm)
	ValidateRegistrationRequest(form *userUsecases.RegistrationForm) error
	ValidateFastRegistrationRequest(form *userUsecases.RegistrationForm) error
}

type AuthorizationInteractor interface {
	ValidateCredentials(form *userUsecases.LoginForm) error
}

type ProfileInteractor interface {
	GetUser(email string) (*userUsecases.User, error)
	ValidateUser() error //todo
	UpdateUser() error   //todo
}

type UserWebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
	ProfileInteractor       ProfileInteractor
}

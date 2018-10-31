package userInterfaces

import (
	"usecases/userUsecases"
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
	UpdateUser() error //todo
}

type WebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
	ProfileInteractor       ProfileInteractor
}

package authInterfaces

import (
	"yurko/usecases/authUsecases"
)

type RegistrationInteractor interface {
	Registration(form *authUsecases.RegistrationForm)
	ValidateRegistrationRequest(form *authUsecases.RegistrationForm) error
	ValidateFastRegistrationRequest(form *authUsecases.RegistrationForm) error
}

type AuthorizationInteractor interface {
	ValidateCredentials(form *authUsecases.LoginForm) error
}

type WebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
}

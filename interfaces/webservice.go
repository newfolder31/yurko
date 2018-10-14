package interfaces

import (
	"yurko/usecases"
)

type RegistrationInteractor interface {
	Registration(form *usecases.RegistrationForm)
	ValidateRegistrationRequest(form *usecases.RegistrationForm) error
	ValidateFastRegistrationRequest(form *usecases.RegistrationForm) error
}

type AuthorizationInteractor interface {
	ValidateCredentials(form *usecases.LoginForm) error
}

type WebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
}

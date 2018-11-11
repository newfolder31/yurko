package handlers

import (
	userUsecases "github.com/newfolder31/yurko/usecases/user"
)

type RegistrationInteractor interface {
	Registration(form *userUsecases.RegistrationForm)
	FastRegistration(form *userUsecases.FastRegistrationForm)
}

type AuthorizationInteractor interface {
	ValidateCredentials(form *userUsecases.LoginForm) error
}

type ProfileInteractor interface {
	GetUser(email string) (*userUsecases.User, error)
	GetProfileResponse(email string) (map[string]interface{}, error)
	ValidateUser(email string, form *userUsecases.ProfileForm) error
	UpdateUser(form *userUsecases.ProfileForm) (*userUsecases.User, error)
}

type UserWebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
	ProfileInteractor       ProfileInteractor
}

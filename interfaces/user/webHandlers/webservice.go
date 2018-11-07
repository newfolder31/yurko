package webHandlers

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
	GetProfileResponse(email string) (map[string]interface{}, error)
	ValidateUser(email string, form *userUsecases.ProfileForm) error
	UpdateUser(form *userUsecases.ProfileForm) (*userUsecases.User, error)
}

type UserWebserviceHandler struct {
	RegistrationInteractor  RegistrationInteractor
	AuthorizationInteractor AuthorizationInteractor
	ProfileInteractor       ProfileInteractor
}

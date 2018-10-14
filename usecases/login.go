package usecases

import (
	"errors"
)

type LoginForm struct {
	Email, Password string
}

type AuthorizationInteractor struct {
	UserRepository UserRepository
}

func (interactor *AuthorizationInteractor) ValidateCredentials(form *LoginForm) error {
	//todo: encode password
	user, _ := interactor.UserRepository.FindByEmailAndPassword(form.Email, form.Password)
	if user == nil {
		return errors.New("credentials are invalid")
	} else if user.IsActive == false {
		return errors.New("user is inactive")
	}

	return nil
}

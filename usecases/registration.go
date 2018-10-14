package usecases

import (
	"errors"
	"fmt"
)

type RegistrationForm struct {
	Email, Password, ConfirmPassword string

	FirstName, LastName, FathersName string
}

type RegistrationInteractor struct {
	UserRepository UserRepository
}

func (interactor *RegistrationInteractor) Registration(form *RegistrationForm) {
	interactor.createUser(form)
	//todo: send confirm email
}

func (interactor *RegistrationInteractor) createUser(form *RegistrationForm) User {
	user := User{}

	user.Email = form.Email
	user.Password = form.Password

	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.FathersName = form.FathersName

	user.IsActive = false

	interactor.UserRepository.Store(user)

	fmt.Println("user created", user)

	return user
}

func (interactor *RegistrationInteractor) ValidateRegistrationRequest(form *RegistrationForm) error {
	if err := interactor.validateEmail(form.Email); err != nil {
		return err
	}
	if err := interactor.validatePasswords(form.Password, form.ConfirmPassword); err != nil {
		return err
	}
	return nil
}

func (interactor *RegistrationInteractor) ValidateFastRegistrationRequest(form *RegistrationForm) error {
	if err := interactor.validateEmail(form.Email); err != nil {
		return err
	}
	return nil
}

func (interactor *RegistrationInteractor) validateEmail(email string) error {
	//todo: find user by email
	//todo: validate on regexp
	return nil
}

func (interactor *RegistrationInteractor) validatePasswords(pass, confirmPass string) error {
	if pass == "" {
		return errors.New("password is empty")
	}
	if confirmPass == "" {
		return errors.New("confirm password is empty")
	}
	if pass != confirmPass {
		return errors.New("passwords are not equals")
	}
	return nil
}

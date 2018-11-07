package userUsecases

import (
	"errors"
	"fmt"
)

type RegistrationForm struct {
	Email, Password string

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

	user.IsActive = true //todo: activate user by email

	interactor.UserRepository.Store(&user)

	fmt.Println("user created", user)

	return user
}

func (interactor *RegistrationInteractor) ValidateRegistrationRequest(form *RegistrationForm) error {
	if err := interactor.validateEmail(form.Email); err != nil {
		return err
	}
	if err := interactor.validatePasswords(form.Password); err != nil {
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
	if email == "" {
		return errors.New("email can't be empty")
	}
	user, _ := interactor.UserRepository.FindByEmail(email)
	if user != nil {
		return errors.New("user with specified email exists")
	}
	//todo: validate on regexp
	return nil
}

func (interactor *RegistrationInteractor) validatePasswords(pass string) error {
	if pass == "" {
		return errors.New("password is empty")
	}
	return nil
}

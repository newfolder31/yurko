package user

import (
	"errors"
	"fmt"
	"github.com/go-ozzo/ozzo-validation/is"
	"regexp"

	"github.com/go-ozzo/ozzo-validation"
)

type FastRegistrationForm struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FathersName string `json:"fathersName"`
}

//todo: add fields  Address, Phone to domain user
type RegistrationForm struct {
	FastRegistrationForm
	Address Address
	Phone   string `json:"phone"`
}
type Address struct {
	Building string `json:"building"`
	Street   string `json:"street"`
	City     string `json:"city"`
}

func (form FastRegistrationForm) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.Email, validation.Required, is.Email),
		validation.Field(&form.Password, validation.Required, validation.Length(8, 50)),
		validation.Field(&form.FirstName, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Zа-яА-Я '-]+$"))),
		validation.Field(&form.LastName, validation.Required, validation.Match(regexp.MustCompile("^[a-zA-Zа-яА-Я '-]+$"))),
		validation.Field(&form.FathersName, validation.Match(regexp.MustCompile("^[a-zA-Zа-яА-Я '-]+$"))),
	)
}
func (form RegistrationForm) Validate() error {
	return validation.ValidateStruct(&form,
		validation.Field(&form.FastRegistrationForm),
		validation.Field(&form.Phone, validation.Required, validation.Match(regexp.MustCompile("^[+]{0,1}[0-9 ]{6,14}$"))), //todo: optimize regexp
		validation.Field(&form.Address),
	)
}
func (address Address) Validate() error {
	return validation.ValidateStruct(&address,
		validation.Field(&address.Building, validation.Length(1, 10)),
		validation.Field(&address.Street, validation.Length(2, 100)),
		validation.Field(&address.City, validation.Length(2, 100)),
	)
}

type RegistrationInteractor struct {
	UserRepository UserRepository
}

func (interactor *RegistrationInteractor) Registration(form *RegistrationForm) {
	interactor.createUser(form)
	//todo: send confirm email
}

func (interactor *RegistrationInteractor) FastRegistration(form *FastRegistrationForm) {
	interactor.fastCreateUser(form)
	//todo: send confirm email
}

func (interactor *RegistrationInteractor) createUser(form *RegistrationForm) (*User, error) {
	if existedUser, _ := interactor.UserRepository.FindByEmail(form.Email); existedUser != nil {
		return nil, errors.New("user with specified email exists")
	}

	user := User{}

	user.Email = form.Email
	user.Password = form.Password

	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.FathersName = form.FathersName

	user.IsActive = true //todo: activate user by email

	interactor.UserRepository.Store(&user)

	fmt.Println("user created", user)

	return &user, nil
}

//todo: analyze and fix structure of fast and regular registration
func (interactor *RegistrationInteractor) fastCreateUser(form *FastRegistrationForm) (*User, error) {
	user := User{}

	user.Email = form.Email
	user.Password = form.Password

	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.FathersName = form.FathersName

	user.IsActive = true //todo: activate user by email

	interactor.UserRepository.Store(&user)

	fmt.Println("user created", user)

	return &user, nil
}

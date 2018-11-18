package user

import "errors"

type ProfileForm struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FathersName string `json:"fathersName"`
	Phone       string `json:"phone"`
	AddressForm
}

type ProfileInteractor struct {
	UserRepository    UserRepository
	AddressRepository AddressRepository
}

func (interactor *ProfileInteractor) GetUser(email string) (*User, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	return user, err
}

func (interactor *ProfileInteractor) GetProfileResponse(email string) (map[string]interface{}, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil || user == nil {
		return nil, err
	}

	userMap := make(map[string]interface{})

	userMap["id"] = user.Id
	userMap["email"] = user.Email
	userMap["firstName"] = user.FirstName
	userMap["lastName"] = user.LastName
	userMap["fathersName"] = user.FathersName
	userMap["phone"] = user.Phone

	if user.AddressId != 0 {
		address, err := interactor.AddressRepository.FindById(user.AddressId)
		if err != nil {
			return nil, err
		}
		userMap["building"] = address.Building
		userMap["city"] = address.City
		userMap["street"] = address.Street
	}

	return userMap, nil
}

func (interactor *ProfileInteractor) ValidateUser(email string, form *ProfileForm) error {
	currentUser, err := interactor.UserRepository.FindByEmail(email)

	if err != nil {
		return err
	} else if currentUser.Id != form.Id {
		return errors.New("permission denied")
	}

	//todo: validate other form fields

	return nil
}

func (interactor *ProfileInteractor) UpdateUser(form *ProfileForm) (*User, error) {
	user, err := interactor.UserRepository.FindById(form.Id)
	if err != nil {
		return nil, err
	}

	user.Email = form.Email
	user.FirstName = form.FirstName
	user.LastName = form.LastName
	user.FathersName = form.FathersName
	user.Phone = form.Phone

	if user.AddressId == 0 {
		address := interactor.createAddress(form)
		user.AddressId = address.Id
	} else {
		interactor.updateAddress(user.AddressId, form)
	}

	interactor.UserRepository.Store(user)

	return user, nil
}

func (interactor *ProfileInteractor) updateAddress(id int, form *ProfileForm) *Address {
	address, _ := interactor.AddressRepository.FindById(id)
	address.Building = form.Building
	address.City = form.City
	address.Street = form.Street
	interactor.AddressRepository.Store(address)
	return address
}

func (interactor *ProfileInteractor) createAddress(form *ProfileForm) *Address {
	address := Address{}
	address.Building = form.Building
	address.City = form.City
	address.Street = form.Street
	interactor.AddressRepository.Store(&address)
	return &address
}

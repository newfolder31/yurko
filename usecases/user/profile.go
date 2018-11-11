package user

import "errors"

type ProfileForm struct {
	Id          int    `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	FathersName string `json:"fathersName"`
}

type ProfileInteractor struct {
	UserRepository UserRepository
}

func (interactor *ProfileInteractor) GetUser(email string) (*User, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	return user, err
}

func (interactor *ProfileInteractor) GetProfileResponse(email string) (map[string]interface{}, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	userMap := make(map[string]interface{})

	userMap["id"] = user.Id
	userMap["email"] = user.Email
	userMap["first_name"] = user.FirstName
	userMap["last_name"] = user.LastName
	userMap["fathers_name"] = user.FathersName

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

	interactor.UserRepository.Store(user)

	return user, nil
}

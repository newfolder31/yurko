package userUsecases

type ProfileForm struct {
	Email, FirstName, LastName, FathersName string
}

type ProfileInteractor struct {
	UserRepository UserRepository
}

func (interactor *ProfileInteractor) GetUser(email string) (*User, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (interactor *ProfileInteractor) ValidateUser(form *ProfileForm) error {
	//validate form

	return nil
}

func (interactor *ProfileInteractor) UpdateUser(email string, form *ProfileForm) (*User, error) {
	user, err := interactor.UserRepository.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	//todo :move params from form to user

	return user, nil
}

package userUsecases

type UserForm struct {
	Email, Password string
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

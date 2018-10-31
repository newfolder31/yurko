package professionUsecases

import (
	"domains/profession"
	"errors"
	"fmt"
)

type ClientForm struct {
	Id int
	//todo: some parameters of client that will be stored in repository
}

type ClientInteractor struct {
	ClientRepository professionDomain.ClientRepository
}

func (interactor *ClientInteractor) CreateClient() (professionDomain.Client, error) {
	client := professionDomain.Client{}
	err := interactor.ClientRepository.Store(&client)

	return client, err
}

func (interactor *ClientInteractor) ValidateUpdateClient(form *ClientForm) error {
	client, _ := interactor.ClientRepository.FindById(form.Id)
	if client == nil {
		return errors.New(fmt.Sprintf("client doesn't exist with id [%v]", form.Id))
	}

	return nil
}

func (interactor *ClientInteractor) UpdateClient(form *ClientForm) error {
	//client, _ := interactor.ClientRepository.FindById(form.Id)

	//todo: update some parameters of client

	return nil
}

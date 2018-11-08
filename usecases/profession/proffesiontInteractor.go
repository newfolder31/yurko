package professionUsecases

import (
	"errors"
	"fmt"
	"github.com/newfolder31/yurko/domains/profession"
)

type ProfessionForm struct {
	Id int
	//todo: some parameters of Profession that will be stored in repository
}

type ProfessionInteractor struct {
	ProfessionRepository professionDomain.ProfessionRepository
}

func (interactor *ProfessionInteractor) CreateProfession() (professionDomain.Profession, error) {
	profession := professionDomain.Profession{}
	err := interactor.ProfessionRepository.Store(&profession)

	return profession, err
}

func (interactor *ProfessionInteractor) ValidateUpdateProfession(form *ProfessionForm) error {
	profession, _ := interactor.ProfessionRepository.FindById(form.Id)
	if profession == nil {
		return errors.New(fmt.Sprintf("profession doesn't exist with id [%v]", form.Id))
	}

	return nil
}

func (interactor *ProfessionInteractor) UpdateProfession(form *ProfessionForm) error {
	//profession, _ := interactor.ProfessionRepository.FindById(form.Id)

	//todo: update some parameters of Profession

	return nil
}

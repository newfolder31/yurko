package interfaces

import (
	"math/rand"
	"yurko/usecases"
)

type UserInMemoryRepo struct {
	data map[int]usecases.User
}

func (repo UserInMemoryRepo) Store(user usecases.User) error {

	r := rand.New(rand.NewSource(999999999))

	repo.data[r.Int()] = user

	return nil
}

func (repo UserInMemoryRepo) FindById(id int) (usecases.User, error) {
	return repo.data[id], nil
}

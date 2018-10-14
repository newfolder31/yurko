package interfaces

import (
	"math/rand"
	"time"
	"yurko/usecases"
)

type UserInMemoryRepo struct {
	data map[int]usecases.User
}

func NewUserInMemoryRepo() *UserInMemoryRepo {
	userRepo := new(UserInMemoryRepo)
	userRepo.data = make(map[int]usecases.User)
	return userRepo
}

func (repo UserInMemoryRepo) Store(user *usecases.User) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int()
	user.Id = id
	repo.data[id] = *user

	return nil
}

func (repo UserInMemoryRepo) FindById(id int) (*usecases.User, error) {
	user, _ := repo.data[id]
	return &user, nil
}

func (repo UserInMemoryRepo) FindByEmail(email string) (*usecases.User, error) {
	for _, value := range repo.data {
		if value.Email == email {
			return &value, nil
		}
	}
	return nil, nil
}

func (repo UserInMemoryRepo) FindByEmailAndPassword(email, password string) (*usecases.User, error) {
	for _, value := range repo.data {
		if value.Email == email && value.Password == password {
			return &value, nil
		}
	}
	return nil, nil
}

package userRepositories

import (
	"math/rand"
	"time"
	"usecases/user"
)

type UserInMemoryRepo struct {
	data map[int]userUsecases.User
}

func NewUserInMemoryRepo() *UserInMemoryRepo {
	userRepo := new(UserInMemoryRepo)
	userRepo.data = make(map[int]userUsecases.User)
	return userRepo
}

func (repo UserInMemoryRepo) Store(user *userUsecases.User) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int()
	user.Id = id
	repo.data[id] = *user

	return nil
}

func (repo UserInMemoryRepo) FindById(id int) (*userUsecases.User, error) {
	user, _ := repo.data[id]
	return &user, nil
}

func (repo UserInMemoryRepo) FindByEmail(email string) (*userUsecases.User, error) {
	for _, value := range repo.data {
		if value.Email == email {
			return &value, nil
		}
	}
	return nil, nil
}

func (repo UserInMemoryRepo) FindByEmailAndPassword(email, password string) (*userUsecases.User, error) {
	for _, value := range repo.data {
		if value.Email == email && value.Password == password {
			return &value, nil
		}
	}
	return nil, nil
}

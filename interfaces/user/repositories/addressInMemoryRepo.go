package repositories

import (
	addressUsecases "github.com/newfolder31/yurko/usecases/user"
	"math/rand"
	"time"
)

type AddressInMemoryRepo struct {
	data map[int]addressUsecases.Address
}

func NewAddressInMemoryRepo() *AddressInMemoryRepo {
	addressRepo := new(AddressInMemoryRepo)
	addressRepo.data = make(map[int]addressUsecases.Address)
	return addressRepo
}

func (repo AddressInMemoryRepo) Store(address *addressUsecases.Address) error {
	if address.Id != 0 {
		repo.data[address.Id] = *address
	} else {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		id := r.Int()
		address.Id = id
		repo.data[id] = *address
	}

	return nil
}

func (repo AddressInMemoryRepo) FindById(id int) (*addressUsecases.Address, error) {
	address, _ := repo.data[id]

	return &address, nil
}

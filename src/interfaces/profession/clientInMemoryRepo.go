package professionInterfaces

import (
	"domains/profession"
	"math/rand"
	"time"
)

type ClientInMemoryRepo struct {
	data map[int]professionDomain.Client
}

func NewClientInMemoryRepo() *ClientInMemoryRepo {
	clientInMemoryRepo := new(ClientInMemoryRepo)
	clientInMemoryRepo.data = make(map[int]professionDomain.Client)
	return clientInMemoryRepo
}

func (repo ClientInMemoryRepo) Store(client *professionDomain.Client) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int()
	client.Id = id
	repo.data[id] = *client

	return nil
}

func (repo ClientInMemoryRepo) FindById(id int) (*professionDomain.Client, error) {
	client, _ := repo.data[id]
	return &client, nil
}

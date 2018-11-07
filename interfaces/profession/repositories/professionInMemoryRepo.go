package repositories

import (
	"github.com/newfolder31/yurko/domains/profession"
	"math/rand"
	"time"
)

type ProfessionInMemoryRepo struct {
	data map[int]professionDomain.Profession
}

func NewProfessionInMemoryRepo() *ProfessionInMemoryRepo {
	professionInMemoryRepo := new(ProfessionInMemoryRepo)
	professionInMemoryRepo.data = make(map[int]professionDomain.Profession)
	return professionInMemoryRepo
}

func (repo ProfessionInMemoryRepo) Store(profession *professionDomain.Profession) error {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := r.Int()
	profession.Id = id
	repo.data[id] = *profession

	return nil
}

func (repo ProfessionInMemoryRepo) FindByUser(userId int) (*professionDomain.Profession, error) {
	profession, _ := repo.data[userId]
	return &profession, nil
}

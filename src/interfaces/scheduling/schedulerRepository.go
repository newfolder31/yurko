package scheduling

import (
	domains "domains/scheduling"
	"errors"
)

type TestSchedulerRepository struct {
	storage []*domains.Scheduler
}

func initTestSchedulerRepository() *TestSchedulerRepository {
	repo := new(TestSchedulerRepository)
	repo.storage = make([]*domains.Scheduler, 0)
	return repo
}

func (repository *TestSchedulerRepository) FindById(id uint64) (*domains.Scheduler, error) {
	for _, item := range repository.storage {
		if item.Id == id {
			resultItem := *item
			return &resultItem, nil
		}
	}
	return nil, nil
}

func (repository *TestSchedulerRepository) FindAllByUserId(userId uint64) (*[]*domains.Scheduler, error) {
	resultSlice := make([]*domains.Scheduler, 0)
	for _, item := range repository.storage {
		if item.UserId == userId {
			resultItem := *item
			resultSlice = append(resultSlice, &resultItem)
		}
	}
	return &resultSlice, nil
}

func (repository *TestSchedulerRepository) Store(scheduler *domains.Scheduler) error {
	if scheduler.Id == 0 {
		scheduler.Id = repository.generateNextId(1)
	}
	repository.storage = append(repository.storage, scheduler)
	return nil
}

func (repository *TestSchedulerRepository) Update(scheduler *domains.Scheduler) error {
	if item, _ := repository.FindById(scheduler.Id); item.Id == 0 {
		return errors.New("Object with current id not found!") // TODO discus!!!???
	}
	repository.Delete(scheduler.Id)
	return repository.Store(scheduler)
}

func (repository *TestSchedulerRepository) Delete(id uint64) error {
	for i, item := range repository.storage {
		if item.Id == id {
			repository.storage = append(repository.storage[:i], repository.storage[i+1:]...)
			return nil
		}
	}
	return nil
}

func (repository *TestSchedulerRepository) generateNextId(step int) uint64 {
	val := uint64(len(repository.storage) + step)
	if i, _ := repository.FindById(val); i != nil {
		return repository.generateNextId(step + 1)
	}
	return val
}

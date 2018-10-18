package scheduling

import (
	domains "domains/scheduling"
	"errors"
)

type TestIntervalRepository struct {
	storage []*domains.Interval
}

func initTestIntervalRepository() *TestIntervalRepository {
	repo := new(TestIntervalRepository)
	repo.storage = make([]*domains.Interval, 0)
	return repo
}

func (repository *TestIntervalRepository) FindById(id uint64) (*domains.Interval, error) {
	for _, item := range repository.storage {
		if item.Id == id {
			resultItem := *item
			return &resultItem, nil
		}
	}
	return nil, nil
}

func (repository *TestIntervalRepository) FindAllBySchedulerId(schedulerId uint64, sortBy string) (*[]*domains.Interval, error) {
	resultSlice := make([]*domains.Interval, 0)
	for _, item := range repository.storage {
		if item.SchedulerId == schedulerId {
			resultItem := *item
			resultSlice = append(resultSlice, &resultItem)
		}
	}
	return &resultSlice, nil
}

func (repository *TestIntervalRepository) FindAllBySchedulerIdAndDay(schedulerId uint64, day uint8) (*[]*domains.Interval, error) {
	resultSlice := make([]*domains.Interval, 0)
	for _, item := range repository.storage {
		if item.SchedulerId == schedulerId && item.GetWeekDay() == day {
			resultItem := *item
			resultSlice = append(resultSlice, &resultItem)
		}
	}
	return &resultSlice, nil
}

func (repository *TestIntervalRepository) FindAllBySchedulerIdAndDate(schedulerId uint64, date int64) (*[]*domains.Interval, error) {
	resultSlice := make([]*domains.Interval, 0)
	for _, item := range repository.storage {
		if item.SchedulerId == schedulerId && item.Date == date {
			resultItem := *item
			resultSlice = append(resultSlice, &resultItem)
		}
	}
	return &resultSlice, nil
}

func (repository *TestIntervalRepository) Store(interval *domains.Interval) error {
	if interval.Id == 0 {
		interval.Id = repository.generateNextId(1)
	}
	repository.storage = append(repository.storage, interval)
	return nil
}

func (repository *TestIntervalRepository) Update(interval *domains.Interval) error {
	if item, _ := repository.FindById(interval.Id); item.Id == 0 {
		return errors.New("Object with current id not found!") // TODO discus!!!???
	}
	repository.Delete(interval.Id)
	return repository.Store(interval)
}

func (repository *TestIntervalRepository) Delete(id uint64) error {
	for i, item := range repository.storage {
		if item.Id == id {
			repository.storage = append(repository.storage[:i], repository.storage[i+1:]...)
			return nil
		}
	}
	return nil
}

func (repository *TestIntervalRepository) DeleteAllBySchedulerIdAndDay(id uint64, weekDay uint8) error {
	for i, item := range repository.storage {
		if item.SchedulerId == id && item.GetWeekDay() == weekDay {
			repository.storage = append(repository.storage[:i], repository.storage[i+1:]...)
		}
	}
	return nil
}

func (repository *TestIntervalRepository) DeleteAllBySchedulerIdAndDate(id uint64, date int64) error {
	for i, item := range repository.storage {
		if item.SchedulerId == id && item.Date == date {
			repository.storage = append(repository.storage[:i], repository.storage[i+1:]...)
		}
	}
	return nil
}

func (repository *TestIntervalRepository) generateNextId(step int) uint64 {
	val := uint64(len(repository.storage) + step)
	if i, _ := repository.FindById(val); i != nil {
		return repository.generateNextId(step + 1)
	}
	return val
}

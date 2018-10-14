package main

import (
	"errors"
	"fmt"
	"time"
	"yurko/domains"
	"yurko/usecases"
)

func main() {
	fmt.Println("Init application")
	fmt.Println(time.Now())

	//TODO start case #1
	testUserId := uint64(1)

	intervalRepository := initTestIntervalRepository()
	schedulerRepository := initTestSchedulerRepository()
	a := usecases.SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]usecases.Day, 0, 5)
	for i := 1; i < 6; i++ {
		start, _ := usecases.InitTime(uint16(i), 0)
		end, _ := usecases.InitTime(uint16(i), 30)
		timeRange, _ := usecases.InitTimeRange(start, end)
		day, _ := usecases.InitDay(uint8(i), []usecases.TimeRange{timeRange})
		days = append(days, day)
	}

	fmt.Println(days)
	a.CreateNewScheduler(testUserId, "Jurist", days)
	fmt.Println("--------------------------------------")

	for _, i := range intervalRepository.storage {
		fmt.Println(i)
	}
	for _, i := range schedulerRepository.storage {
		fmt.Println(i)
	}

	//TODO end case #1

	//TODO test code
	//b := domains.Interval{Id: 1, SchedulerId: 1, From: 230, To: 300, Date: 232113313}
	//bPointer := &b
	//c := *bPointer
	//cPointer := &c
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(cPointer)
	//
	//bPointer.Id = 2
	//
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Println(cPointer)
}

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

func (repository *TestIntervalRepository) generateNextId(step int) uint64 {
	val := uint64(len(repository.storage) + step)
	if i, _ := repository.FindById(val); i != nil {
		return repository.generateNextId(step + 1)
	}
	return val
}

//-----------------------------------------------------------------------------

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

package main

import (
	domains "domains/scheduling"
	"errors"
	"fmt"
	"sort"
	"time"
)

func main() {
	fmt.Println("Init application")
	fmt.Println(time.Now())

	//TODO start case #1
	//testUserId := uint64(1)
	//
	//intervalRepository := initTestIntervalRepository()
	//schedulerRepository := initTestSchedulerRepository()
	//a := usecases.SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}
	//
	//days := make([]usecases.Day, 0, 5)
	//for i := 1; i < 6; i++ {
	//	start, _ := usecases.InitTime(uint16(i), 0)
	//	end, _ := usecases.InitTime(uint16(i), 30)
	//	timeRange, _ := usecases.InitTimeRange(start, end)
	//	day, _ := usecases.InitDay(uint8(i), []usecases.TimeRange{timeRange})
	//	days = append(days, day)
	//}
	//
	//fmt.Println(days)
	//a.CreateNewScheduler(testUserId, "Jurist", days)
	//fmt.Println("--------------------------------------")
	//
	//for _, i := range intervalRepository.storage {
	//	fmt.Println(i)
	//}
	//for _, i := range schedulerRepository.storage {
	//	fmt.Println(i)
	//}
	//TODO end case #1

	//TODO test code
	var a []*Test
	a1 := Test{a: 2}
	a2 := Test{a: 1}
	a3 := Test{a: 4}
	a4 := Test{a: 3}
	a5 := Test{a: 5}
	a = append(a, &a1)
	a = append(a, &a2)
	a = append(a, &a3)
	a = append(a, &a4)
	a = append(a, &a5)
	p := &a
	sort.Slice(*p, func(i, j int) bool {
		return (*p)[i].a < (*p)[j].a
	})

	for _, i := range a {
		fmt.Println(i)
	}

}

type Test struct {
	a int
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

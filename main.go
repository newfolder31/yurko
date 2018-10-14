package main

import (
	"fmt"
	"time"
	"yurko/domains"
	"yurko/usecases"
)

func main() {
	fmt.Println("Init application")
	fmt.Println(time.Now())

	testUserId := uint64(1)

	var IntervalRepository domains.IntervalRepository
	var SchedulerRepository domains.SchedulerRepository
	a := usecases.SchedulingInteractor{IntervalRepository: IntervalRepository, SchedulerRepository: SchedulerRepository}

	days := make([]usecases.Day, 0, 5)
	for i := 1; i < 6; i++ {
		start, _ := usecases.InitTime(uint16(i), 0)
		end, _ := usecases.InitTime(uint16(i), 30)
		timeRange, _ := usecases.InitTimeRange(start, end)
		day, _ := usecases.InitDay(uint8(i), []usecases.TimeRange{timeRange})
		days = append(days, day)
	}

	//fmt.Println(days)
	a.CreateNewScheduler(testUserId, "Jurist", days)

}

type TestIntervalRepository struct {
	storage []domains.Interval
}

func initTestIntervalRepository() *TestIntervalRepository {
	repo := new(TestIntervalRepository)
	repo.storage = make([]domains.Interval, 0)
	return repo
}

func (repository *TestIntervalRepository) FindById(id uint64) domains.Interval {
	for _, item := range repository.storage {
		if item.Id == id {
			return item
		}
	}
	return domains.Interval{}
}

func (repository *TestIntervalRepository) FindAllBySchedulerId(schedulerId uint64, sortBy string) []domains.Interval {
	result := make([]domains.Interval, 0)
	for _, item := range repository.storage {
		if item.SchedulerId == schedulerId {
			result = append(result, item)
		}
	}
	return result
}

func (repository *TestIntervalRepository) Store(interval domains.Interval) domains.Interval {
	if interval.Id == 0 {
		interval.Id = repository.generateNextId(1)
	}
	repository.storage = append(repository.storage, interval)
	return interval
}

func (repository *TestIntervalRepository) Update(interval domains.Interval) domains.Interval {
	if item := repository.FindById(interval.Id); item.Id == 0 {
		return item
	}
	repository.Delete(interval.Id)
	return repository.Store(interval)
}

func (repository *TestIntervalRepository) Delete(id uint64) {

}

func (repository *TestIntervalRepository) generateNextId(step int) uint64 {
	val := uint64(len(repository.storage) + step)
	if i := repository.FindById(val); i.Id != 0 || val == i.Id {
		return repository.generateNextId(step + 1)
	}
	return val
}

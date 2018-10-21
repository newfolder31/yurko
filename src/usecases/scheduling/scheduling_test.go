package scheduling

import (
	domains "domains/scheduling"
	"interfaces/scheduling"
	"sort"
	"testing"
	time2 "time"
)

func TestCreatingScheduler(t *testing.T) {
	testUserId := uint64(1)

	intervalRepository := scheduling.InitTestIntervalRepository()
	schedulerRepository := scheduling.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]Day, 0, 5)
	for i := 0; i < 5; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		timeRange, _ := InitTimeRange(start, end)
		day, _ := InitDay(uint8(i), []TimeRange{timeRange})
		days = append(days, day)
	}

	if err, _ := a.CreateNewScheduler(testUserId, "Jurist", &days); err != nil {
		t.Error("Error during creating new scheduler!")
	}

	if resultScheduler, err := schedulerRepository.FindAllByUserId(testUserId); err != nil {
		t.Error("Error during finding scheduler by user id")
	} else if len(*resultScheduler) != 1 {
		t.Error("Expected schedulers size: 1 , resulted: ", len(*resultScheduler))
	} else {
		schedulerId := (*resultScheduler)[0].Id
		intervals, err := intervalRepository.FindAllBySchedulerId(schedulerId, "")
		if err != nil {
			t.Error("Error during finding all intervals by scheduler id")
		}
		if len(*intervals) != 5 {
			t.Error("Expected intervals size: 5, resulted: ", len(*intervals))
		}

		sortIntervalSliceByDay(intervals)
		for i, interval := range *intervals {
			from := days[i].ranges[0].start.timeInMinutes
			to := days[i].ranges[0].end.timeInMinutes
			if interval.From != from || interval.To != to {
				t.Error("Unexpected interval #", i,
					" From = ", interval.From,
					" Expected = ", from,
					" To = ", interval.To,
					" Expected = ", to)
			}
		}
	}
}

func sortIntervalSliceByDay(slice *[]*domains.Interval) {
	sort.Slice(*slice, func(i, j int) bool {
		return (*slice)[i].GetWeekDay() < (*slice)[j].GetWeekDay()
	})
}

func TestGetAllSchedulerByUserId(t *testing.T) {
	testUserId := uint64(1)

	intervalRepository := scheduling.InitTestIntervalRepository()
	schedulerRepository := scheduling.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	if err, _ := a.CreateNewScheduler(testUserId, "Jurist", nil); err != nil {
		t.Error("Error during creating scheduler for Jurist", err)
	}

	if err, _ := a.CreateNewScheduler(testUserId, "Test", nil); err != nil {
		t.Error("Error during creating scheduler for Test", err)
	}

	if err, _ := a.CreateNewScheduler(testUserId, "Test2", nil); err != nil {
		t.Error("Error during creating scheduler for Test2", err)
	}

	if schedulers, err := a.getAllSchedulersByUserId(testUserId); err == nil {
		if schedulers == nil {
			t.Error("Result of getting all schedulers by user id is nil")
		}
		if len(*schedulers) != 3 {
			t.Error("Invalid result of getting all schedulers by user id",
				"schedulers slice result:", len(*schedulers))
		}
	} else {
		t.Error("Error during getting all schedulers by user id", err)
	}
}

func Test_BuildSchedulerForDate_forRegularDay(t *testing.T) {
	testUserId := uint64(1)

	intervalRepository := scheduling.InitTestIntervalRepository()
	schedulerRepository := scheduling.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]Day, 0, 1)

	rangesDate := [][]uint16{{9, 0, 9, 30}, {8, 0, 8, 30}, {10, 0, 10, 30}}
	ranges := make([]TimeRange, 0, 3)

	for _, day := range rangesDate {
		start, _ := InitTime(day[0], day[1])
		end, _ := InitTime(day[2], day[3])
		timeRange, _ := InitTimeRange(start, end)
		ranges = append(ranges, timeRange)
	}
	day, _ := InitDay(5, ranges)
	days = append(days, day)

	scheduler, err := a.CreateNewScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	time := time2.Date(2018, 10, 20, 0, 0, 0, 0, nil) //it's Saturday
	if result, err := a.BuildSchedulerForDate(scheduler.Id, time); err != nil {
		if len(*result) != 3 {

		}
	} else {
		t.Error("Error during building scheduler: ", err)
	}
}

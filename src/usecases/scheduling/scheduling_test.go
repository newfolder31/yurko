package scheduling

import (
	domains "domains/scheduling"
	"interfaces/scheduling"
	"sort"
	"testing"
	"time"
)

var preparedDate time.Time = time.Date(2018, time.October, 20, 0, 0, 0, 0, time.UTC)

func TestCreatingScheduler(t *testing.T) {
	testUserId := uint64(1)

	intervalRepository := scheduling.InitTestIntervalRepository()
	schedulerRepository := scheduling.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	//preparing test data
	// 5 - days with single interval for each day
	days := make([]Day, 0, 5)
	for i := 0; i < 5; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		timeRange, _ := InitTimeRange(start, end)
		day, _ := InitDay(uint8(i), []TimeRange{timeRange})
		days = append(days, day)
	}

	//creating scheduler by userId, days and profession type - "Jurist",
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
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(intervals)

		//comparing resulted intervals with appropriate days
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

	if _, err := a.CreateNewScheduler(testUserId, "Jurist", nil); err != nil {
		t.Error("Error during creating scheduler for Jurist", err)
	}

	if _, err := a.CreateNewScheduler(testUserId, "Test", nil); err != nil {
		t.Error("Error during creating scheduler for Test", err)
	}

	if _, err := a.CreateNewScheduler(testUserId, "Test2", nil); err != nil {
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

func TestAddExceptionsInScheduler(t *testing.T) {
	testUserId := uint64(1)

	intervalRepository := scheduling.InitTestIntervalRepository()
	schedulerRepository := scheduling.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	//creating scheduler by userId, days and profession type - "Jurist",
	scheduler, err := a.CreateNewScheduler(testUserId, "Jurist", nil)
	if err != nil {
		t.Error("Error during creating new scheduler!")
	}

	//-----------Part 1-----------
	//create three time ranges
	timeRanges := make([]TimeRange, 0, 3)
	for i := 9; i < 12; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception := ExceptionalDate{date: preparedDate.Unix(), ranges: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.IntervalRepository.FindAllBySchedulerId(scheduler.Id, ""); err != nil {
		t.Error("Error during finding all by scheduler id", err)
		return
	} else {
		if len(*result) != 3 {
			t.Error("Invalid result of getting all intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < 3; i++ {
			if (*result)[i].From != timeRanges[i].start.timeInMinutes {
				t.Error("Invalid interval. Expected: ",
					timeRanges[i].start.timeInMinutes, " - ", timeRanges[i].end.timeInMinutes,
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//-----------Part 2-----------
	//Test deleting old exceptions for the same date
	timeRanges = make([]TimeRange, 0, 2)
	for i := 13; i < 15; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception = ExceptionalDate{date: preparedDate.Unix(), ranges: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.IntervalRepository.FindAllBySchedulerId(scheduler.Id, ""); err != nil {
		t.Error("Error during finding all by scheduler id", err)
		return
	} else {
		if len(*result) != 2 {
			t.Error("Invalid result of getting all intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < 2; i++ {
			if (*result)[i].From != timeRanges[i].start.timeInMinutes {
				t.Error("Invalid interval. Expected: ",
					timeRanges[i].start.timeInMinutes, " - ", timeRanges[i].end.timeInMinutes,
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
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
	day, _ := InitDay(6, ranges) // 6 - it's Saturday
	days = append(days, day)

	scheduler, err := a.CreateNewScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedDate); err == nil {
		if len(*result) != 3 {
			t.Error("Expected slice size: 3, resulted: ", len(*result))
		} else {
			//range[1] - 08:00 - 08:30
			if (*result)[0].From != ranges[1].start.timeInMinutes ||
				(*result)[0].To != ranges[1].end.timeInMinutes {
				t.Error("Invalid interval. Expected: ",
					ranges[1].start.timeInMinutes, " - ", ranges[1].end.timeInMinutes,
					"resulted: ", (*result)[0].From, " - ", (*result)[0].To)
			}
			//range[1] - 09:00 - 09:30
			if (*result)[1].From != ranges[0].start.timeInMinutes ||
				(*result)[1].To != ranges[0].end.timeInMinutes {
				t.Error("Invalid interval. Expected: ",
					ranges[0].start.timeInMinutes, " - ", ranges[0].end.timeInMinutes,
					"resulted: ", (*result)[1].From, " - ", (*result)[1].To)
			}
			//range[1] - 10:00 - 10:30
			if (*result)[2].From != ranges[2].start.timeInMinutes ||
				(*result)[2].To != ranges[2].end.timeInMinutes {
				t.Error("Invalid interval. Expected: ",
					ranges[2].start.timeInMinutes, " - ", ranges[2].end.timeInMinutes,
					"resulted: ", (*result)[2].From, " - ", (*result)[2].To)
			}
		}
	} else {
		t.Error("Error during building scheduler: ", err)
	}
}

func Test_BuildSchedulerForDate_forExceptionalDate(t *testing.T) {
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
	day, _ := InitDay(6, ranges) // 6 - it's Saturday
	days = append(days, day)

	scheduler, err := a.CreateNewScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	timeRanges := make([]TimeRange, 0, 2)
	for i := 13; i < 15; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception := ExceptionalDate{date: preparedDate.Unix(), ranges: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedDate); err == nil {
		if len(*result) != 2 {
			t.Error("Expected slice size: 2, resulted: ", len(*result))
		} else {
			for i := 0; i < 2; i++ {
				if (*result)[i].From != timeRanges[i].start.timeInMinutes {
					t.Error("Invalid interval. Expected: ",
						timeRanges[i].start.timeInMinutes, " - ", timeRanges[i].end.timeInMinutes,
						"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
				}
			}
		}
	} else {
		t.Error("Error during building scheduler: ", err)
	}
}

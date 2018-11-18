package scheduling

import (
	domains "github.com/newfolder31/yurko/domains/scheduling"
	"github.com/newfolder31/yurko/interfaces/scheduling/repositories"
	"sort"
	"testing"
	"time"
)

var preparedFridayDate time.Time = time.Date(2018, time.October, 19, 0, 0, 0, 0, time.UTC)   // Friday
var preparedSaturdayDate time.Time = time.Date(2018, time.October, 20, 0, 0, 0, 0, time.UTC) // Saturday

var testUserId = uint64(1)

func TestCreatingScheduler(t *testing.T) {
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	//preparing test data
	// 5 - days with single interval for each day
	days := make([]Day, 0, 5)
	for i := 0; i < 5; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		timeRange, _ := InitTimeRange(start, end)
		day, _ := InitDay(uint8(i), []Interval{timeRange})
		days = append(days, day)
	}

	//creating scheduler by userId, days and profession type - "Jurist",
	if _, err := a.CreateScheduler(testUserId, "Jurist", &days); err != nil {
		t.Error("Error during creating new scheduler!", err)
	}

	if resultScheduler, err := schedulerRepository.FindAllByUserId(testUserId); err != nil {
		t.Error("Error during finding scheduler by user id")
	} else if len(*resultScheduler) != 1 {
		t.Error("Expected schedulers size: 1 , resulted: ", len(*resultScheduler))
	} else {
		schedulerId := (*resultScheduler)[0].Id
		intervals, err := intervalRepository.FindAllBySchedulerId(schedulerId, "")
		if err != nil {
			t.Error("Error during finding all Intervals by scheduler id")
		}
		if len(*intervals) != 5 {
			t.Error("Expected Intervals size: 5, resulted: ", len(*intervals))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(intervals)

		//comparing resulted Intervals with appropriate days
		for i, interval := range *intervals {
			from := days[i].Intervals[0].Start.GetTimeInMinutes()
			to := days[i].Intervals[0].End.GetTimeInMinutes()
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
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	if _, err := a.CreateScheduler(testUserId, "Jurist", nil); err != nil {
		t.Error("Error during creating scheduler for Jurist", err)
	}

	if _, err := a.CreateScheduler(testUserId, "Test", nil); err != nil {
		t.Error("Error during creating scheduler for Test", err)
	}

	if _, err := a.CreateScheduler(testUserId, "Test2", nil); err != nil {
		t.Error("Error during creating scheduler for Test2", err)
	}

	if schedulers, err := a.GetAllSchedulersByUserId(testUserId); err == nil {
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
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	//creating scheduler by userId, days and profession type - "Jurist",
	scheduler, err := a.CreateScheduler(testUserId, "Jurist", nil)
	if err != nil {
		t.Error("Error during creating new scheduler!", err)
	}

	//-----------Part 1-----------
	//create three time Intervals
	timeRanges := make([]Interval, 0, 3)
	for i := 9; i < 12; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception := ExceptionalDate{Date: preparedSaturdayDate.Unix(), Intervals: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.IntervalRepository.FindAllBySchedulerId(scheduler.Id, ""); err != nil {
		t.Error("Error during finding all by scheduler id", err)
		return
	} else {
		if len(*result) != 3 {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < 3; i++ {
			if (*result)[i].From != timeRanges[i].Start.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					timeRanges[i].Start.GetTimeInMinutes(), " - ", timeRanges[i].End.GetTimeInMinutes(),
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//-----------Part 2-----------
	//Test overriding old exceptions for the same date
	timeRanges = make([]Interval, 0, 2) //timeRanges - Intervals [[13:00-13:30], [14:00-14:30]]
	for i := 13; i < 15; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception = ExceptionalDate{Date: preparedSaturdayDate.Unix(), Intervals: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.IntervalRepository.FindAllBySchedulerId(scheduler.Id, ""); err != nil {
		t.Error("Error during finding all by scheduler id", err)
		return
	} else {
		if len(*result) != 2 {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < 2; i++ {
			if (*result)[i].From != timeRanges[i].Start.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					timeRanges[i].Start.GetTimeInMinutes(), " - ", timeRanges[i].End.GetTimeInMinutes(),
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//-----------Part 3-----------
	//Test deleting exceptions for specified day
	exception = ExceptionalDate{Date: preparedSaturdayDate.Unix()}

	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.IntervalRepository.FindAllBySchedulerId(scheduler.Id, ""); err != nil {
		t.Error("Error during finding all by scheduler id", err)
		return
	} else {
		if len(*result) != 0 {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result), " - expected result: 0")
		}
	}
}

func Test_BuildSchedulerForDate_forRegularDay(t *testing.T) {
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]Day, 0, 1)

	rangesDate := [][]uint16{{9, 0, 9, 30}, {8, 0, 8, 30}, {10, 0, 10, 30}}
	ranges := make([]Interval, 0, 3)

	for _, day := range rangesDate {
		start, _ := InitTime(day[0], day[1])
		end, _ := InitTime(day[2], day[3])
		timeRange, _ := InitTimeRange(start, end)
		ranges = append(ranges, timeRange)
	}
	day, _ := InitDay(6, ranges) // 6 - it's Saturday
	days = append(days, day)

	scheduler, err := a.CreateScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedSaturdayDate); err == nil {
		if len(*result) != 3 {
			t.Error("Expected slice size: 3, resulted: ", len(*result))
		} else {
			//range[1] - 08:00 - 08:30
			if (*result)[0].From != ranges[1].Start.GetTimeInMinutes() ||
				(*result)[0].To != ranges[1].End.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[1].Start.GetTimeInMinutes(), " - ", ranges[1].End.GetTimeInMinutes(),
					"resulted: ", (*result)[0].From, " - ", (*result)[0].To)
			}
			//range[1] - 09:00 - 09:30
			if (*result)[1].From != ranges[0].Start.GetTimeInMinutes() ||
				(*result)[1].To != ranges[0].End.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[0].Start.GetTimeInMinutes(), " - ", ranges[0].End.GetTimeInMinutes(),
					"resulted: ", (*result)[1].From, " - ", (*result)[1].To)
			}
			//range[1] - 10:00 - 10:30
			if (*result)[2].From != ranges[2].Start.GetTimeInMinutes() ||
				(*result)[2].To != ranges[2].End.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[2].Start.GetTimeInMinutes(), " - ", ranges[2].End.GetTimeInMinutes(),
					"resulted: ", (*result)[2].From, " - ", (*result)[2].To)
			}
		}
	} else {
		t.Error("Error during building scheduler: ", err)
	}
}

func Test_BuildSchedulerForDate_forExceptionalDate(t *testing.T) {
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]Day, 0, 1)

	rangesDate := [][]uint16{{9, 0, 9, 30}, {8, 0, 8, 30}, {10, 0, 10, 30}}
	ranges := make([]Interval, 0, 3)

	for _, day := range rangesDate {
		start, _ := InitTime(day[0], day[1])
		end, _ := InitTime(day[2], day[3])
		timeRange, _ := InitTimeRange(start, end)
		ranges = append(ranges, timeRange)
	}
	day, _ := InitDay(6, ranges) // 6 - it's Saturday
	days = append(days, day)

	scheduler, err := a.CreateScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	timeRanges := make([]Interval, 0, 2)
	for i := 13; i < 15; i++ {
		start, _ := InitTime(uint16(i), 0)
		end, _ := InitTime(uint16(i), 30)
		if timeRange, err := InitTimeRange(start, end); err != nil {
			t.Error("Invalid time range creating!", err)
		} else {
			timeRanges = append(timeRanges, timeRange)
		}
	}

	exception := ExceptionalDate{Date: preparedSaturdayDate.Unix(), Intervals: timeRanges}
	if err := a.AddExceptionInScheduler(scheduler.Id, &exception); err != nil {
		t.Error("Error during adding exceptions!")
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedSaturdayDate); err == nil {
		if len(*result) != 2 {
			t.Error("Expected slice size: 2, resulted: ", len(*result))
		} else {
			for i := 0; i < 2; i++ {
				if (*result)[i].From != timeRanges[i].Start.GetTimeInMinutes() {
					t.Error("Invalid interval. Expected: ",
						timeRanges[i].Start.GetTimeInMinutes(), " - ", timeRanges[i].End.GetTimeInMinutes(),
						"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
				}
			}
		}
	} else {
		t.Error("Error during building scheduler: ", err)
	}
}

func Test_UpdateDayIntervals(t *testing.T) {
	intervalRepository := repositories.InitTestIntervalRepository()
	schedulerRepository := repositories.InitTestSchedulerRepository()
	a := SchedulingInteractor{IntervalRepository: intervalRepository, SchedulerRepository: schedulerRepository}

	days := make([]Day, 0, 1)

	rangesDate := [][]uint16{{8, 0, 8, 30}, {9, 0, 9, 30}, {10, 0, 10, 30}}
	ranges := make([]Interval, 0, 3)

	for _, day := range rangesDate {
		start, _ := InitTime(day[0], day[1])
		end, _ := InitTime(day[2], day[3])
		timeRange, _ := InitTimeRange(start, end)
		ranges = append(ranges, timeRange)
	}
	day, _ := InitDay(6, ranges) // 6 - it's Saturday
	days = append(days, day)

	scheduler, err := a.CreateScheduler(testUserId, "Test", &days)
	if err != nil {
		t.Error("Error during creating scheduler: ", err)
	}

	//-----------Part 1-----------
	//Simple Intervals update
	day, _ = InitDay(uint8(time.Friday), ranges) // 5 - it's Friday
	if err := a.UpdateDayIntervals(scheduler.Id, day); err != nil {
		t.Error("Error during update day Intervals: ", err)
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedFridayDate); err != nil {
		t.Error("Error during building scheduler: ", err)
	} else {
		if len(*result) != len(ranges) {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < len(ranges); i++ {
			if (*result)[i].From != ranges[i].Start.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[i].Start.GetTimeInMinutes(), " - ", ranges[i].End.GetTimeInMinutes(),
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//check for latest Intervals is presents
	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedSaturdayDate); err != nil {
		t.Error("Error during building scheduler: ", err)
	} else {
		if len(*result) != len(ranges) {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < len(ranges); i++ {
			if (*result)[i].From != ranges[i].Start.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[i].Start.GetTimeInMinutes(), " - ", ranges[i].End.GetTimeInMinutes(),
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//-----------Part 2-----------
	//Overriding Intervals
	rangesDate = [][]uint16{{10, 0, 11, 30}, {12, 30, 14, 30}}
	ranges = make([]Interval, 0, 2)

	for _, day := range rangesDate {
		start, _ := InitTime(day[0], day[1])
		end, _ := InitTime(day[2], day[3])
		timeRange, _ := InitTimeRange(start, end)
		ranges = append(ranges, timeRange)
	}
	day, _ = InitDay(uint8(time.Saturday), ranges)

	if err := a.UpdateDayIntervals(scheduler.Id, day); err != nil {
		t.Error("Error during update day Intervals: ", err)
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedSaturdayDate); err != nil {
		t.Error("Error during building scheduler: ", err)
	} else {
		if len(*result) != len(ranges) {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result))
		}
		//resulted interval not save sequence after FindAllBySchedulerId
		sortIntervalSliceByDay(result)
		for i := 0; i < len(ranges); i++ {
			if (*result)[i].From != ranges[i].Start.GetTimeInMinutes() {
				t.Error("Invalid interval. Expected: ",
					ranges[i].Start.GetTimeInMinutes(), " - ", ranges[i].End.GetTimeInMinutes(),
					"resulted: ", (*result)[i].From, " - ", (*result)[i].To)
			}
		}
	}

	//-----------Part 2-----------
	//Deleting Intervals
	ranges = make([]Interval, 0, 0)
	day, _ = InitDay(uint8(time.Saturday), ranges)

	if err := a.UpdateDayIntervals(scheduler.Id, day); err != nil {
		t.Error("Error during update day Intervals: ", err)
	}

	if result, err := a.BuildSchedulerForDate(scheduler.Id, preparedSaturdayDate); err != nil {
		t.Error("Error during building scheduler: ", err)
	} else {
		if len(*result) != len(ranges) {
			t.Error("Invalid result of getting all Intervals by scheduler id",
				"schedulers slice result:", len(*result), " - expected result: ", len(ranges))
		}
	}
}

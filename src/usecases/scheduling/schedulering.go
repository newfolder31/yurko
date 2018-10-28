package scheduling

import (
	domains "domains/scheduling"
	"errors"
	"sort"
	"time"
)

type SchedulingInteractor struct {
	SchedulerRepository domains.SchedulerRepository
	IntervalRepository  domains.IntervalRepository
}

type Day struct {
	WeekDay   uint8      `json:"weekDay"`
	Intervals []Interval `json:"intervals"`
}

type ExceptionalDate struct {
	date      int64
	Intervals []Interval
}

type Interval struct {
	Start Time `json:"start"`
	End   Time `json:"end"`
}

type Time struct {
	Hours   uint16 `json:"hours"`
	Minutes uint16 `json:"minutes"`
}

func (t *Time) GetTimeInMinutes() uint16 {
	return uint16((t.Hours)*60 + t.Minutes)
}

func InitTime(hours, minutes uint16) (Time, error) {
	if 0 > hours || hours > 23 {
		return Time{}, errors.New("Invalid hours value [" + string(hours) + "]")
	} else if 0 > minutes || minutes > 59 {
		return Time{}, errors.New("Invalid minutes value [" + string(minutes) + "]")
	}
	return Time{Hours: hours, Minutes: minutes}, nil
}

func InitTimeRange(start, end Time) (Interval, error) {
	if end.Compare(start) < 1 {
		return Interval{}, errors.New("Invalid time range. End time not later then Start time!")
	}
	return Interval{Start: start, End: end}, nil
}

func InitDay(day uint8, ranges []Interval) (Day, error) {
	if day < 0 || day > 6 {
		return Day{}, errors.New("Invalid WeekDay number [" + string(day) + "]")
	}
	return Day{WeekDay: day, Intervals: ranges}, nil
}

func (current *Time) Compare(compared Time) int8 {
	if current.GetTimeInMinutes() > compared.GetTimeInMinutes() {
		return 1
	} else if current.GetTimeInMinutes() < compared.GetTimeInMinutes() {
		return -1
	}
	return 0
}

func (interactor *SchedulingInteractor) CreateScheduler(userId uint64, professionType string, days *[]Day) (*domains.Scheduler, error) {
	scheduler := domains.Scheduler{UserId: userId, ProfessionType: professionType}
	err := interactor.SchedulerRepository.Store(&scheduler) //TODO handle error!
	if err != nil {
		return nil, errors.New("Function CreateScheduler! Unable create new Scheduler!")
	}
	if days != nil {
		for _, day := range *days {
			if err := interactor.storeIntervalsFromDay(scheduler.Id, day); err != nil {
				return nil, err //TODO Rewrap error
			}
		}
	}
	return &scheduler, nil
}

func (interactor *SchedulingInteractor) GetAllSchedulersByUserId(userId uint64) (*[]*domains.Scheduler, error) {
	if schedulers, err := interactor.SchedulerRepository.FindAllByUserId(userId); err != nil {
		return nil, err ////TODO Rewrap error
	} else {
		return schedulers, nil
	}
}

func (interactor *SchedulingInteractor) BuildSchedulerForDateRange(schedulerId uint64, dates *[]time.Time) (map[time.Time]*[]*domains.Interval, error) {
	result := make(map[time.Time]*[]*domains.Interval)
	sort.Slice(*dates, func(i, j int) bool {
		return (*dates)[i].Unix() < (*dates)[j].Unix()
	})
	for _, item := range *dates {
		var err error
		result[item], err = interactor.BuildSchedulerForDate(schedulerId, item)
		if err != nil {
			return result, err
		}
	}
	return result, nil
}

func (interactor *SchedulingInteractor) BuildSchedulerForDate(schedulerId uint64, date time.Time) (*[]*domains.Interval, error) {
	date = cleanUpDate(&date)
	exceptionsSlice, err := interactor.IntervalRepository.FindAllBySchedulerIdAndDate(schedulerId, date.Unix())
	if err != nil {
		return nil, err ////TODO Rewrap error
	}
	if len(*exceptionsSlice) != 0 {
		sort.Slice(*exceptionsSlice, func(i, j int) bool {
			return (*exceptionsSlice)[i].From < (*exceptionsSlice)[j].From
		})
		sortIntervalSliceByDate(exceptionsSlice)
		return exceptionsSlice, nil
	}

	regularSlice, err := interactor.IntervalRepository.FindAllBySchedulerIdAndDay(schedulerId, uint8(date.Weekday()))
	if err != nil {
		return nil, err ////TODO Rewrap error
	}

	sortIntervalSliceByDate(regularSlice)
	return regularSlice, nil
}

func sortIntervalSliceByDate(slice *[]*domains.Interval) {
	sort.Slice(*slice, func(i, j int) bool {
		return (*slice)[i].From < (*slice)[j].From
	})
}

func (interactor *SchedulingInteractor) AddExceptionInScheduler(schedulerId uint64, exception *ExceptionalDate) error {
	scheduler, _ := interactor.SchedulerRepository.FindById(schedulerId) // TODO handle error!
	if scheduler != nil && exception != nil {
		interactor.IntervalRepository.DeleteAllBySchedulerIdAndDate(schedulerId, exception.date)
		if exception.Intervals != nil {
			interactor.storeIntervalsFromExceptions(schedulerId, exception)
		}
		return nil
	}
	errorMessage := "Scheduler with current id[" + string(schedulerId) + "] not found!"
	return errors.New(errorMessage)
}

func (interactor *SchedulingInteractor) AddExceptionSliceInScheduler(schedulerId uint64, exceptions *[]ExceptionalDate) error {
	scheduler, _ := interactor.SchedulerRepository.FindById(schedulerId) // TODO handle error!
	if scheduler != nil && exceptions != nil {
		for _, oneException := range *exceptions {
			interactor.IntervalRepository.DeleteAllBySchedulerIdAndDate(schedulerId, oneException.date) // DeleteDateExceptions
			if oneException.Intervals != nil {
				interactor.storeIntervalsFromExceptions(schedulerId, &oneException)
			}
		}
		return nil
	}
	errorMessage := "Scheduler with current id[" + string(schedulerId) + "] not found!"
	return errors.New(errorMessage)
}

func (interactor *SchedulingInteractor) UpdateDayIntervals(schedulerId uint64, day Day) error {
	scheduler, _ := interactor.SchedulerRepository.FindById(schedulerId)
	if scheduler == nil {
		return errors.New("Scheduler with current id-" + string(schedulerId) + " is not found!")
	}
	if err := interactor.DeleteDayIntervals(schedulerId, day.WeekDay); err != nil {
		return err //TODO Rewrap error
	}
	if err := interactor.storeIntervalsFromDay(schedulerId, day); err != nil {
		return err //TODO Rewrap error
	}
	return nil
}

func (interactor *SchedulingInteractor) DeleteScheduler(schedulerId uint64) error {
	if err := interactor.IntervalRepository.DeleteAllBySchedulerId(schedulerId); err != nil {
		return err
	}
	if err := interactor.SchedulerRepository.Delete(schedulerId); err != nil {
		return err
	}
	return nil
}

func (interactor *SchedulingInteractor) DeleteDayIntervals(schedulerId uint64, day uint8) error {
	return interactor.IntervalRepository.DeleteAllBySchedulerIdAndDay(schedulerId, day)
}

func (interactor *SchedulingInteractor) DeleteDateExceptions(schedulerId uint64, date time.Time) error {
	return interactor.IntervalRepository.DeleteAllBySchedulerIdAndDate(schedulerId, date.Unix())
}

func (interactor *SchedulingInteractor) storeIntervalsFromDay(schedulerId uint64, day Day) error {
	for _, timeRange := range day.Intervals {
		start := timeRange.Start.GetTimeInMinutes()
		end := timeRange.End.GetTimeInMinutes()
		interval, err := domains.InitInterval(schedulerId, int64(0), start, end, day.WeekDay)
		if err != nil {
			return err
		}
		interactor.IntervalRepository.Store(&interval)
	}
	return nil
}

func (interactor *SchedulingInteractor) storeIntervalsFromExceptions(schedulerId uint64, exception *ExceptionalDate) error {
	for _, timeRange := range exception.Intervals {
		start := timeRange.Start.GetTimeInMinutes()
		end := timeRange.End.GetTimeInMinutes()
		interval, err := domains.InitInterval(schedulerId, exception.date, start, end, 1)
		if err != nil {
			return err
		}
		interactor.IntervalRepository.Store(&interval)
	}
	return nil
}

//remove hours, minutes, seconds
func cleanUpDate(date *time.Time) time.Time {
	year, month, day := date.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

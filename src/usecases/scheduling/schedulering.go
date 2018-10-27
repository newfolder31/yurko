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

type Time struct {
	timeInMinutes uint16
}

type TimeRange struct {
	start Time
	end   Time
}

type Day struct {
	weekDay uint8
	ranges  []TimeRange
}

type ExceptionalDate struct {
	date   int64
	ranges []TimeRange
}

func InitTime(hours, minutes uint16) (Time, error) {
	if 0 > hours || hours > 23 {
		return Time{}, errors.New("Invalid hours value [" + string(hours) + "]")
	} else if 0 > minutes || minutes > 59 {
		return Time{}, errors.New("Invalid minutes value [" + string(minutes) + "]")
	}
	return Time{timeInMinutes: (hours)*60 + minutes}, nil
}

func InitTimeRange(start, end Time) (TimeRange, error) {
	if end.Compare(start) < 1 {
		return TimeRange{}, errors.New("Invalid time range. End time not later then start time!")
	}
	return TimeRange{start: start, end: end}, nil
}

func InitDay(day uint8, ranges []TimeRange) (Day, error) {
	if day < 0 || day > 6 {
		return Day{}, errors.New("Invalid weekDay number [" + string(day) + "]")
	}
	return Day{weekDay: day, ranges: ranges}, nil
}

func (current *Time) Compare(compared Time) int8 {
	if current.timeInMinutes > compared.timeInMinutes {
		return 1
	} else if current.timeInMinutes < compared.timeInMinutes {
		return -1
	}
	return 0
}

//*добавление времени работы по умолчанию в профайле
// а занчит, нужно созать для пользователя (userId) рассписание по
// значениям [weekDay: 1, ranges: [[start:12:00, end:13:00],[],[]]] ----> CreateNewScheduler
//*быстрая модификация указанного времени
// schedulerId + [date: 21.12.1996, ranges: [[start:12:00, end:13:00],[],[]]]
//*получение рассписания на по дате
// schedulerId, date
//- получение рассписания на неделю, а лучше на слайс дат

func (interactor *SchedulingInteractor) CreateNewScheduler(userId uint64, professionType string, days *[]Day) (*domains.Scheduler, error) {
	scheduler := domains.Scheduler{UserId: userId, ProfessionType: professionType}
	err := interactor.SchedulerRepository.Store(&scheduler) //TODO handle error!
	if err != nil {
		return nil, errors.New("Function CreateNewScheduler! Unable create new Scheduler!")
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

func (interactor *SchedulingInteractor) getAllSchedulersByUserId(userId uint64) (*[]*domains.Scheduler, error) {
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

//*в случае если был хоть один тайм рейнж был модифицирован
// для исключительных ситуаций передаються все рейжи этого дня и сохраняються как исключительные тайм ренжи
// сохранять тайм ренжи можно за один раз,
// про попытке сохранить в несколько раз ренжи, прошлые ренжи на этот день должны удаляться
func (interactor *SchedulingInteractor) AddExceptionInScheduler(schedulerId uint64, exception *ExceptionalDate) error {
	scheduler, _ := interactor.SchedulerRepository.FindById(schedulerId) // TODO handle error!
	//TODO добавить проверку на то, что количество, интервалов не равно нулю,
	//TODO в случае если всётаки равно нулю, то удалить все интервалы за указанную дату и допилить на это тесты
	if scheduler != nil && exception != nil {
		interactor.IntervalRepository.DeleteAllBySchedulerIdAndDate(schedulerId, exception.date)
		if exception.ranges != nil {
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
			interactor.IntervalRepository.DeleteAllBySchedulerIdAndDate(schedulerId, oneException.date)
			if oneException.ranges != nil {
				interactor.storeIntervalsFromExceptions(schedulerId, &oneException)
			}
		}
		return nil
	}
	errorMessage := "Scheduler with current id[" + string(schedulerId) + "] not found!"
	return errors.New(errorMessage)
}

//with all interval
func RemoveScheduler(schedulerId uint64) error {
	return nil
}

//update all ranges in weekDay in one action
func (interactor *SchedulingInteractor) UpdateDayIntervals(schedulerId uint64, day Day) error {
	//удалить прошлые за день, вставить новые
	//TODO добавить проверку на то, что количество, интервалов не равно нулю,
	//TODO в случае если всётаки равно нулю, то удалить все интервалы за указанную дату и допилить на это тесты
	scheduler, _ := interactor.SchedulerRepository.FindById(schedulerId)
	if scheduler == nil {
		return errors.New("Scheduler with current id-" + string(schedulerId) + " is not found!")
	}
	if err := interactor.DeleteDayIntervals(schedulerId, day.weekDay); err != nil {
		return err //TODO Rewrap error
	}
	if err := interactor.storeIntervalsFromDay(schedulerId, day); err != nil {
		return err //TODO Rewrap error
	}
	return nil
}

func (interactor *SchedulingInteractor) DeleteDayIntervals(schedulerId uint64, day uint8) error {
	if err := interactor.IntervalRepository.DeleteAllBySchedulerIdAndDay(schedulerId, day); err != nil {
		return err //TODO Rewrap error
	}
	return nil
}

func (interactor *SchedulingInteractor) DeleteDayExceptions(schedulerId uint64, date time.Time) error {
	return nil
}

func (interactor *SchedulingInteractor) storeIntervalsFromDay(schedulerId uint64, day Day) error {
	for _, timeRange := range day.ranges {
		start := timeRange.start.timeInMinutes
		end := timeRange.end.timeInMinutes
		interval, err := domains.InitInterval(schedulerId, int64(0), start, end, day.weekDay)
		if err != nil {
			return err
		}
		interactor.IntervalRepository.Store(&interval)
	}
	return nil
}

func (interactor *SchedulingInteractor) storeIntervalsFromExceptions(schedulerId uint64, exception *ExceptionalDate) error {
	for _, timeRange := range exception.ranges {
		start := timeRange.start.timeInMinutes
		end := timeRange.end.timeInMinutes
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

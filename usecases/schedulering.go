package usecases

import (
	"errors"
	"fmt"
	"yurko/domains"
)

type SchedulingInteractor struct {
	SchedulerRepository domains.SchedulerRepository
	IntervalRepository  domains.IntervalRepository
	//Logger
}

type Time struct {
	timeInMinutes uint16
}

type TimeRange struct {
	start Time
	end   Time
}

type Day struct {
	day    uint8
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
	if 0 > day || day > 7 {
		return Day{}, errors.New("Invalid day number [" + string(day) + "]")
	}
	return Day{day: day, ranges: ranges}, nil
}

func (current *Time) Compare(compared Time) int8 {
	if current.timeInMinutes > compared.timeInMinutes {
		return 1
	} else if current.timeInMinutes < compared.timeInMinutes {
		return -1
	}
	return 0
}

//
//*добавление времени работы по умолчанию в профайле
// а занчит, нужно созать для пользователя (userId) рассписание по
// значениям [day: 1, ranges: [[start:12:00, end:13:00],[],[]]]
//*быстрая модификация указанного времени
// schedulerId + [date: 21.12.1996, ranges: [[start:12:00, end:13:00],[],[]]]
//*получение рассписания на по дате
// schedulerId, date
//- получение рассписания на неделю, а лучше на слайс дат

func (interactor *SchedulingInteractor) CreateNewScheduler(userId uint64, professionType string, days []Day) error {
	scheduler := interactor.SchedulerRepository.Store(domains.Scheduler{UserId: userId, ProfessionType: professionType})
	intervals := make([]domains.Interval, 0, 14) //to test var
	for _, day := range days {
		for _, timeRange := range day.ranges {
			start := timeRange.start.timeInMinutes
			end := timeRange.end.timeInMinutes
			interval, err := domains.InitInterval(scheduler.Id, uint64(0), start, end, day.day)
			if err != nil {
				return err
			}
			interactor.IntervalRepository.Store(interval)
			intervals = append(intervals, interval)
		}
	}
	fmt.Println(intervals)
	return nil
}

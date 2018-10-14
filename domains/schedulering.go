package domains

import "errors"

//todo it's mock for development
const (
	PROFFESSIONAL_TYPE_LAWYERS = iota
)

type SchedulerRepository interface {
	FindById(id uint64) Scheduler //TODO , error ?????
	FindAllByUserId(userId uint64) []Scheduler
	Store(scheduler Scheduler) Scheduler
	Update(scheduler Scheduler) Scheduler
	Delete(id uint64)
}

type IntervalRepository interface {
	FindById(id uint64) Interval
	FindAllBySchedulerId(schedulerId uint64, sortBy string) []Interval
	Store(interval Interval) Interval
	Update(interval Interval) Interval
	Delete(id uint64)
}

type Scheduler struct {
	Id             uint64
	UserId         uint64
	ProfessionType string
}

type Interval struct {
	Id          uint64
	SchedulerId uint64
	From        uint16 // to
	To          uint16 // to
	Date        uint64 // date for exceptional situation: overriding default interval
	weekDay     uint8
}

func InitInterval(schedulerId, Date uint64, from, to uint16, weekDay uint8) (Interval, error) {
	item := Interval{SchedulerId: schedulerId, From: from, To: to, Date: Date}
	if err := item.SetWeekDay(weekDay); err != nil {
		return item, err
	}

	return item, nil
}

func (item *Interval) SetWeekDay(weekDay uint8) error {
	if weekDay < 1 || weekDay > 7 {
		return errors.New("Illegal week weekDay number!")
	}
	item.weekDay = weekDay
	return nil
}

func (item *Interval) GetWeekDay() uint8 {
	return item.weekDay
}

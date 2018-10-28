package scheduling

import "errors"

const (
	PROFFESSIONAL_TYPE_LAWYERS = iota
)

type SchedulerRepository interface {
	FindById(id uint64) (*Scheduler, error) //TODO , error ?????
	FindAllByUserId(userId uint64) (*[]*Scheduler, error)
	Store(scheduler *Scheduler) error
	Update(scheduler *Scheduler) error
	Delete(id uint64) error
}

type IntervalRepository interface {
	FindById(id uint64) (*Interval, error)
	FindAllBySchedulerId(schedulerId uint64, sortBy string) (*[]*Interval, error)
	FindAllBySchedulerIdAndDay(schedulerId uint64, day uint8) (*[]*Interval, error)
	FindAllBySchedulerIdAndDate(schedulerId uint64, date int64) (*[]*Interval, error)
	Store(interval *Interval) error
	Update(interval *Interval) error
	Delete(id uint64) error
	DeleteAllBySchedulerId(schedulerId uint64) error
	DeleteAllBySchedulerIdAndDay(schedulerId uint64, weekDay uint8) error
	DeleteAllBySchedulerIdAndDate(schedulerId uint64, date int64) error
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
	Date        int64  // date for exceptional situation: overriding default interval
	weekDay     uint8
}

func InitInterval(schedulerId uint64, Date int64, from, to uint16, weekDay uint8) (Interval, error) {
	item := Interval{SchedulerId: schedulerId, From: from, To: to, Date: Date}
	if err := item.SetWeekDay(weekDay); err != nil {
		return item, err
	}

	return item, nil
}

func (item *Interval) SetWeekDay(weekDay uint8) error {
	if weekDay < 0 || weekDay > 6 {
		return errors.New("Illegal week weekDay number!")
	}
	item.weekDay = weekDay
	return nil
}

func (item *Interval) GetWeekDay() uint8 {
	return item.weekDay
}

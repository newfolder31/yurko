package handlers

import (
	"encoding/json"
	domains "github.com/newfolder31/yurko/domains/scheduling"
	usecases "github.com/newfolder31/yurko/usecases/scheduling"
	"net/http"
	"time"
)

type SchedulingInteractor interface {
	CreateScheduler(userId uint64, professionType string, days *[]usecases.Day) (*domains.Scheduler, error)
	GetAllSchedulersByUserId(userId uint64) (*[]*domains.Scheduler, error)
	BuildSchedulerForDateRange(schedulerId uint64, dates *[]time.Time) (map[time.Time]*usecases.ExceptionalDate, error)
}

type ScheduleWebserviceHandler struct {
	SchedulingInteractor SchedulingInteractor
}

type schedule struct {
	Id             uint64         `json:"scheduleId"`
	ProfessionType string         `json:"professionType"`
	Days           []usecases.Day `json:"days"`
}

type date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

func (d *date) toTime() time.Time {
	return time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
}

func fromTime(t time.Time) date {
	return date{
		Year:  t.Year(),
		Month: int(t.Month()),
		Day:   t.Day(),
	}
}

type createScheduleRequest struct {
	UserId    uint64     `json:"userId"`
	Schedules []schedule `json:"schedules"`
}

type getAllSchedulersRequest struct {
	UserId uint64 `json:"userId"`
}

type buildSchedulesForDatesRequest struct {
	SchedulerId uint64 `json:"schedulerId"`
	Dates       []date `json:"dates"`
}

type buildSchedulesForDatesResponse struct {
	Date      date                `json:"date"`
	Intervals []usecases.Interval `json:"intervals"`
}

func (handler *ScheduleWebserviceHandler) CreateScheduler(w http.ResponseWriter, r *http.Request) {
	interactor := handler.SchedulingInteractor

	decoder := json.NewDecoder(r.Body)
	var t createScheduleRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	for _, it := range t.Schedules {
		if _, err := interactor.CreateScheduler(t.UserId, it.ProfessionType, &it.Days); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (handler *ScheduleWebserviceHandler) GetAllSchedulesByUserId(w http.ResponseWriter, r *http.Request) {
	interactor := handler.SchedulingInteractor

	decoder := json.NewDecoder(r.Body)
	var t getAllSchedulersRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	schedules, err := interactor.GetAllSchedulersByUserId(t.UserId)
	if err != nil {
		errorMap := map[string]string{"Error": err.Error()}
		errorJsonMessage, _ := json.Marshal(errorMap)
		http.Error(w, string(errorJsonMessage), http.StatusBadRequest)
		return
	}
	result, _ := json.Marshal(&schedules)
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}

func (handler *ScheduleWebserviceHandler) BuildSchedulesForDatesRequest(w http.ResponseWriter, r *http.Request) {
	interactor := handler.SchedulingInteractor

	decoder := json.NewDecoder(r.Body)
	var t buildSchedulesForDatesRequest

	if err := decoder.Decode(&t); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	preparedTime := convertDateSliceToTime(t.Dates)

	builtSchedules, err := interactor.BuildSchedulerForDateRange(t.SchedulerId, &preparedTime)
	if err != nil {
		errorMap := map[string]string{"Error": err.Error()}
		errorJsonMessage, _ := json.Marshal(errorMap)
		http.Error(w, string(errorJsonMessage), http.StatusBadRequest)
		return
	}

	result := make([]buildSchedulesForDatesResponse, 0, len(preparedTime))
	for k, v := range builtSchedules {
		item := buildSchedulesForDatesResponse{}
		item.Date = date{Year: k.Year(), Month: int(k.Month()), Day: k.Day()}
		item.Intervals = v.Intervals
		result = append(result, item)
	}

	marshaledResult, _ := json.Marshal(&result)
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshaledResult)
}

func convertDateSliceToTime(dates []date) (result []time.Time) {
	result = make([]time.Time, 0, len(dates))
	for _, i := range dates {
		result = append(result, i.toTime())
	}
	return
}

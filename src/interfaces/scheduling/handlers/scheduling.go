package handlers

import (
	domains "domains/scheduling"
	"encoding/json"
	"net/http"
	usecases "usecases/scheduling"
)

type SchedulingInteractor interface {
	CreateScheduler(userId uint64, professionType string, days *[]usecases.Day) (*domains.Scheduler, error)
	GetAllSchedulersByUserId(userId uint64) (*[]*domains.Scheduler, error)
}

type ScheduleWebserviceHandler struct {
	SchedulingInteractor *SchedulingInteractor
}

type schedule struct {
	Id             uint64         `json:"scheduleId"`
	ProfessionType string         `json:"professionType"`
	Days           []usecases.Day `json:"days"`
}

type createScheduleRequest struct {
	UserId    uint64     `json:"userId"`
	Schedules []schedule `json:"schedules"`
}

type getAllSchedulersRequest struct {
	userId uint64 `json:"userId"`
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
		if _, err := SchedulingInteractor(*interactor).CreateScheduler(t.UserId, it.ProfessionType, &it.Days); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func (handler *ScheduleWebserviceHandler) GetAllSchedulersByUserId(w http.ResponseWriter, r *http.Request) {
	interactor := handler.SchedulingInteractor

	decoder := json.NewDecoder(r.Body)
	var t getAllSchedulersRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	if _, err := SchedulingInteractor(*interactor).GetAllSchedulersByUserId(t.userId); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	json.NewEncoder(w)
}

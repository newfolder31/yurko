package handlers

import (
	"encoding/json"
	domains "github.com/newfolder31/yurko/domains/scheduling"
	usecases "github.com/newfolder31/yurko/usecases/scheduling"
	"net/http"
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
	UserId uint64 `json:"userId"`
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

func (handler *ScheduleWebserviceHandler) GetAllSchedulesByUserId(w http.ResponseWriter, r *http.Request) {
	interactor := handler.SchedulingInteractor

	decoder := json.NewDecoder(r.Body)
	var t getAllSchedulersRequest
	err := decoder.Decode(&t)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	schedules, err := SchedulingInteractor(*interactor).GetAllSchedulersByUserId(t.UserId)
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

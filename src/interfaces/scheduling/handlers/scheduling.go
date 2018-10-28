package handlers

import (
	domains "domains/scheduling"
	"net/http"
	usecases "usecases/scheduling"
)

type SchedulingInteractor interface {
	CreateScheduler(userId uint64, professionType string, days *[]usecases.Day) (*domains.Scheduler, error)
}

type ScheduleWebserviceHandler struct {
	SchedulingInteractor SchedulingInteractor
}

func CreateScheduler(w http.ResponseWriter, r *http.Request) {

}

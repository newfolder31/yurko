package interfaces

import (
	schedulingDomains "domains/scheduling"
	schedulingUsecases "usecases/scheduling"
)

type SchedulingInteractor interface {
	CreateScheduler(userId uint64, professionType string, days *[]schedulingUsecases.Day) (*schedulingDomains.Scheduler, error)
}

type WebserviceHandler struct {
	SchedulingInteractor schedulingUsecases.SchedulingInteractor
}

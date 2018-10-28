package interfaces

import (
	schedulingDomains "domains/scheduling"
	schedulingUsecases "usecases/scheduling"
)

type WebserviceHandler struct {
	SchedulingInteractor SchedulingInteractor
}

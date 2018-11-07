package daemons

import (
	"github.com/newfolder31/yurko/interfaces"
	schedulingHandlers "github.com/newfolder31/yurko/interfaces/scheduling/handlers"
	schedulingRepos "github.com/newfolder31/yurko/interfaces/scheduling/repositories"
	schedulingUsecases "github.com/newfolder31/yurko/usecases/scheduling"
	"net/http"
)

func Run() error {
	intervalRepository := schedulingRepos.InitTestIntervalRepository()
	schedulerRepository := schedulingRepos.InitTestSchedulerRepository()

	schedulingInteractor := schedulingUsecases.SchedulingInteractor{
		IntervalRepository:  intervalRepository,
		SchedulerRepository: schedulerRepository,
	}
	interactor := schedulingHandlers.SchedulingInteractor(&schedulingInteractor)

	scheduleHandler := schedulingHandlers.ScheduleWebserviceHandler{}
	scheduleHandler.SchedulingInteractor = &interactor

	webserviceHandler := interfaces.WebserviceHandler{
		ScheduleWebserviceHandler: scheduleHandler,
	}

	initScheduleHandling(&webserviceHandler)

	http.ListenAndServe(":8081", nil)
	return nil
}

func initScheduleHandling(webserviceHandler *interfaces.WebserviceHandler) {
	scheduleHandler := &webserviceHandler.ScheduleWebserviceHandler

	http.HandleFunc("/scheduling/createSchedule", func(res http.ResponseWriter, req *http.Request) {
		scheduleHandler.CreateScheduler(res, req)
	})

	http.HandleFunc("/scheduling/getAllSchedulers", func(res http.ResponseWriter, req *http.Request) {
		scheduleHandler.GetAllSchedulesByUserId(res, req)
	})
}

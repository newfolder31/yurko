package schedulingDaemon

import (
	"github.com/go-chi/chi"
	schedulingHandlers "github.com/newfolder31/yurko/interfaces/scheduling/handlers"
	schedulingRepos "github.com/newfolder31/yurko/interfaces/scheduling/repositories"
	schedulingUsecases "github.com/newfolder31/yurko/usecases/scheduling"
	"net/http"
)

func InitUserModule(r *chi.Mux) {

	intervalRepository := schedulingRepos.InitTestIntervalRepository()
	schedulerRepository := schedulingRepos.InitTestSchedulerRepository()

	schedulingInteractor := new(schedulingUsecases.SchedulingInteractor)
	schedulingInteractor.IntervalRepository = intervalRepository
	schedulingInteractor.SchedulerRepository = schedulerRepository

	scheduleHandler := new(schedulingHandlers.ScheduleWebserviceHandler)
	scheduleHandler.SchedulingInteractor = schedulingInteractor

	r.Post("/api/v0/scheduling/createSchedule", func(res http.ResponseWriter, req *http.Request) {
		scheduleHandler.CreateScheduler(res, req)
	})

	r.Post("/api/v0/scheduling/getAllSchedulers", func(res http.ResponseWriter, req *http.Request) {
		scheduleHandler.GetAllSchedulesByUserId(res, req)
	})

	r.Post("/api/v0/scheduling/buildSchedules", func(res http.ResponseWriter, req *http.Request) {
		scheduleHandler.BuildSchedulesForDatesRequest(res, req)
	})
}

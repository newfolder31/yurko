package daemons

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/newfolder31/yurko/daemons/userDaemon"
	"github.com/newfolder31/yurko/interfaces"
	schedulingHandlers "github.com/newfolder31/yurko/interfaces/scheduling/handlers"
	schedulingRepos "github.com/newfolder31/yurko/interfaces/scheduling/repositories"
	schedulingUsecases "github.com/newfolder31/yurko/usecases/scheduling"
	"net/http"
	"os"
)

func Run() error {
	r := chi.NewRouter()

	corsRule := cors.New(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Author ization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsRule.Handler)

	userDaemon.InitUserModule(r)

	r.Handle("/", indexHandler())

	a := os.Getenv("PORT")
	if len(a) == 0 {
		a = "8081"
	}

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

	http.ListenAndServe(":"+a, r)

	return nil
}

//todo: delete
func indexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
		cookie, _ := r.Cookie("sessionId")
		if cookie != nil {
			userEmail := interfaces.GetCurrentUserEmail(cookie.Value)
			fmt.Fprintf(w, "Current user is %s\n", userEmail)
		}
	})
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	//if origin == "http://example.com" {
	//	return true
	//}
	return true
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


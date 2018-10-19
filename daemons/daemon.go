package daemons

import (
	"net/http"
	"yurko/interfaces"
	"yurko/usecases/task"
)

func Run() error {
	inMemoStorage := interfaces.NewMemoStorage()
	taskInMemoRepo := &interfaces.TaskInMemoRepo{Storage: inMemoStorage}
	relationInMemoRepo := &interfaces.RelationInMemoRepo{Storage: inMemoStorage}
	communicationInMemoRepo := &interfaces.CommunicationInMemoRepo{Storage: inMemoStorage}

	taskInteractor := new(task.TaskInteractor)
	taskInteractor.TaskRepository = taskInMemoRepo
	taskInteractor.RelationRepository = relationInMemoRepo
	taskInteractor.CommunicationRepository = communicationInMemoRepo

	webServiceHandler := interfaces.WebServiceHandler{}
	webServiceHandler.Interactor = taskInteractor

	http.HandleFunc("/task/announce", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.Announce(res, req)
	})

	http.HandleFunc("/task/request", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.Request(res, req)
	})

	http.HandleFunc("/task/task", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.Task(res, req)
	})

	http.HandleFunc("/task/tasklist", func(res http.ResponseWriter, req *http.Request) {
		webServiceHandler.TaskList(res, req)
	})

	http.ListenAndServe(":8081", nil)

	return nil
}

package interfaces

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"yurko/domains/task"
)

type TaskInteractor interface {
	CreateAnnounce(userId int, description string) (*task.Task, error)
	CreateRequest(userId int, description string, lawyerUserId int) (*task.Task, error)
	//AssignTask(task *task.Task, lawyerUserId int) error
	TaskList(userId int) []*task.Task
	Task(taskId int) *task.Task
}

type WebServiceHandler struct {
	Interactor TaskInteractor
}

func (handler *WebServiceHandler) Announce(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(request.FormValue("userId"))
	description := request.FormValue("description")
	//todo: validation

	_, err := handler.Interactor.CreateAnnounce(userId, description)

	if err != nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func (handler *WebServiceHandler) Request(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(request.FormValue("userId"))
	description := request.FormValue("description")
	lawyerUserId := request.FormValue("lawyerUserId")
	//todo: validation

	_, err := handler.Interactor.CreateRequest(userId, description, lawyerUserId)

	if err != nil {
		writer.WriteHeader(http.StatusOK)
	} else {
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func (handler *WebServiceHandler) TaskList(writer http.ResponseWriter, request *http.Request) {
	//todo
}

func (handler *WebServiceHandler) Task(writer http.ResponseWriter, request *http.Request) {
	//todo
}

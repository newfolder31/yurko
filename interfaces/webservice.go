package interfaces

import (
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
	"yurko/domains/task"
	usecases "yurko/usecases/task"
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
		fmt.Fprint(writer, "Form parsing error : ", err)
		return
	}

	form := new(usecases.AnnounceForm)
	if err := schema.NewDecoder().Decode(form, request.Form); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Parsing error : ", err)
	} else {
		fmt.Println(form.UserId, form.Description)
		_, err = handler.Interactor.CreateAnnounce(form.UserId, form.Description)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Error during creating announce : ", err)
		} else {
			//todo: send id of new announce?
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func (handler *WebServiceHandler) Request(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Form parsing error : ", err)
		return
	}

	form := new(usecases.RequestForm)
	if err := schema.NewDecoder().Decode(form, request.Form); err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Parsing error : ", err)
	} else {
		fmt.Println(form.UserId, form.Description)
		_, err = handler.Interactor.CreateRequest(form.UserId, form.Description, form.LawyerId)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Error during creating request : ", err)
		} else {
			//todo: send id of new request?
			writer.WriteHeader(http.StatusOK)
		}
	}
}

func (handler *WebServiceHandler) TaskList(writer http.ResponseWriter, request *http.Request) {
	//todo
}

func (handler *WebServiceHandler) Task(writer http.ResponseWriter, request *http.Request) {
	//todo
}

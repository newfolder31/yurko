package interfaces

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"net/http"
	"yurko/domains/task"
	usecases "yurko/usecases/task" //todo: resolve name overriding
)

type TaskInteractor interface {
	CreateAnnounce(userId int, description string) (*task.Task, error)
	CreateRequest(userId int, description string, lawyerUserId int) (*task.Task, error)
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
		announce, err := handler.Interactor.CreateAnnounce(form.UserId, form.Description)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Error during creating announce : ", err)
		} else {
			writer.WriteHeader(http.StatusOK)
			jsonedAnnounce, _ := json.Marshal(announce)
			writer.Write(jsonedAnnounce)
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
		request, err := handler.Interactor.CreateRequest(form.UserId, form.Description, form.LawyerId)

		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(writer, "Error during creating request : ", err)
		} else {
			jsonedRequest, _ := json.Marshal(request)
			writer.Write(jsonedRequest)
		}
	}
}

func (handler *WebServiceHandler) TaskList(writer http.ResponseWriter, request *http.Request) {
	//todo
}

func (handler *WebServiceHandler) Task(writer http.ResponseWriter, request *http.Request) {
	//todo
}

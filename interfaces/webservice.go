package interfaces

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yurko/usecases"
)

type RegistrationInteractor interface {
	Registration(form *usecases.RegistrationForm)
	ValidateRegistrationRequest(form *usecases.RegistrationForm) error
	ValidateFastRegistrationRequest(form *usecases.RegistrationForm) error
}

type WebserviceHandler struct {
	RegistrationInteractor RegistrationInteractor
}

func (webservice WebserviceHandler) FastRegistration(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	email := r.FormValue("email")
	fmt.Println("email", email)
	decoder := json.NewDecoder(r.Body)
	form := usecases.RegistrationForm{}
	err := decoder.Decode(&form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		if err := webservice.RegistrationInteractor.ValidateFastRegistrationRequest(&form); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(&form)
			w.WriteHeader(http.StatusOK)
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}

func (webservice WebserviceHandler) Registration(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	decoder := json.NewDecoder(r.Body)
	form := usecases.RegistrationForm{}
	err := decoder.Decode(&form)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.RegistrationInteractor.ValidateRegistrationRequest(&form); err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(&form)
			w.WriteHeader(http.StatusOK)
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}

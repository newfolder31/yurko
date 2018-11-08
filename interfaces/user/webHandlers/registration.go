package webHandlers

import (
	"encoding/json"
	"fmt"
	"github.com/newfolder31/yurko/usecases/userUsecases"
	"net/http"
)

func (webservice UserWebserviceHandler) FastRegistration(w http.ResponseWriter, r *http.Request) {
	form := new(userUsecases.RegistrationForm)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.RegistrationInteractor.ValidateFastRegistrationRequest(form); err != nil {
			fmt.Fprintf(w, "some error in request params: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(form)
			w.WriteHeader(http.StatusOK)
		}
	}
}

func (webservice UserWebserviceHandler) Registration(w http.ResponseWriter, r *http.Request) {
	form := new(userUsecases.RegistrationForm)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := webservice.RegistrationInteractor.ValidateRegistrationRequest(form); err != nil {
			fmt.Fprintf(w, "some error in request params: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(form)
			w.WriteHeader(http.StatusOK)
		}
	}
}

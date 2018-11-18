package handlers

import (
	"encoding/json"
	"fmt"
	userUsecases "github.com/newfolder31/yurko/usecases/user"
	"net/http"
)

func (webservice UserWebserviceHandler) FastRegistration(w http.ResponseWriter, r *http.Request) {
	form := new(userUsecases.FastRegistrationForm)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		if err := form.Validate(); err != nil {
			jsonByte, _ := json.Marshal(err)
			fmt.Fprintf(w, string(jsonByte))
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.FastRegistration(form)
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
		if err := form.Validate(); err != nil {
			jsonByte, _ := json.Marshal(err)
			fmt.Fprintf(w, string(jsonByte))
			w.WriteHeader(http.StatusBadRequest)
		} else {
			webservice.RegistrationInteractor.Registration(form)
			w.WriteHeader(http.StatusOK)
		}
	}
}

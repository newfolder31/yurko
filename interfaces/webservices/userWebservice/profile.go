package userWebservice

import (
	"fmt"
	"github.com/gorilla/schema"
	"github.com/newfolder31/yurko/infrastructures"
	"github.com/newfolder31/yurko/usecases/userUsecases"
	"net/http"
)

func (webservice UserWebserviceHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		cookie, err := r.Cookie(infrastructures.COOKIE_SESSION_NAME)

		if cookie == nil || err != nil || cookie.Value == "" {
			fmt.Fprintf(w, "authorization is failed: %s!", err)
			w.WriteHeader(http.StatusBadRequest)
		} else {

			var email = infrastructures.InMemorySession.Get(cookie.Value)
			var user, _ = webservice.ProfileInteractor.GetUser(email)

			//todo: return user as json
			fmt.Fprint(w, user)

			w.WriteHeader(http.StatusOK)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
func (webservice UserWebserviceHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	//if r.Method == http.MethodPost {
	if err := r.ParseForm(); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	form := new(userUsecases.ProfileForm)
	if err := schema.NewDecoder().Decode(form, r.Form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		cookie, err := r.Cookie(infrastructures.COOKIE_SESSION_NAME)

		if cookie.Value == "" || err != nil {
			fmt.Fprintf(w, "authorization is failed: %s!", err)
			w.WriteHeader(http.StatusUnauthorized)
		} else {

			var email = infrastructures.InMemorySession.Get(cookie.Value)
			if err := webservice.ProfileInteractor.ValidateUser(form); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				/*var user, _ = */ webservice.ProfileInteractor.UpdateUser(email, form)
				w.WriteHeader(http.StatusOK)
			}
		}
	}
	//} else {
	//	w.WriteHeader(http.StatusMethodNotAllowed)
	//}
}

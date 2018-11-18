package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/newfolder31/yurko/infrastructures"
	userUsecases "github.com/newfolder31/yurko/usecases/user"
	"net/http"
)

func (webservice UserWebserviceHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(infrastructures.COOKIE_SESSION_NAME)

	if cookie == nil || err != nil || cookie.Value == "" {
		fmt.Fprintf(w, "authorization is failed: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {

		var email = infrastructures.InMemorySession.Get(cookie.Value)

		response, _ := webservice.ProfileInteractor.GetProfileResponse(email)
		userJson, _ := json.Marshal(response)
		fmt.Fprint(w, string(userJson))

		w.WriteHeader(http.StatusOK)
	}
}
func (webservice UserWebserviceHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	form := new(userUsecases.ProfileForm)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&form); err != nil {
		fmt.Fprintf(w, "some error in parse request params: %s!", err)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		sessionId, err := r.Cookie(infrastructures.COOKIE_SESSION_NAME)

		if sessionId.Value == "" || err != nil {
			fmt.Fprintf(w, "authorization is failed: %s!", err)
			w.WriteHeader(http.StatusUnauthorized)
		} else {

			var email = infrastructures.InMemorySession.Get(sessionId.Value)
			if err := webservice.ProfileInteractor.ValidateUser(email, form); err != nil {
				w.WriteHeader(http.StatusBadRequest)
			} else {
				var user, _ = webservice.ProfileInteractor.UpdateUser(form)
				if email != user.Email {
					infrastructures.InMemorySession.Update(sessionId.Value, user.Email)
				}
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

package daemons

import (
	"fmt"
	"github.com/newfolder31/yurko/daemons/userDaemon"
	"github.com/newfolder31/yurko/interfaces"
	"net/http"
)

func Run() error {
	userDaemon.InitUserModule()

	http.Handle("/", indexHandler())

	http.ListenAndServe(":8081", nil)

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

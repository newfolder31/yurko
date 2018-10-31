package daemons

import (
	"fmt"
	"net/http"
	"daemons/userDaemon"
	"interfaces"
)

func Run() error {
	userDaemon.InitAuthModule()

	http.Handle("/", indexHandler())

	http.ListenAndServe(":8081", nil)

	return nil
}

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

package daemons

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/newfolder31/yurko/daemons/schedulingDaemon"
	"github.com/newfolder31/yurko/daemons/userDaemon"
	"github.com/newfolder31/yurko/interfaces"
	"net/http"
	"os"
)

func Run() error {
	r := chi.NewRouter()
	initCors(r)

	userDaemon.InitUserModule(r)
	schedulingDaemon.InitUserModule(r)

	r.Handle("/", indexHandler())

	a := os.Getenv("PORT")
	if len(a) == 0 {
		a = "8081"
	}

	http.ListenAndServe(":"+a, r)

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

func initCors(r *chi.Mux) {
	corsRule := cors.New(cors.Options{
		AllowOriginFunc:  AllowOriginFunc,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Author ization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(corsRule.Handler)
}

func AllowOriginFunc(r *http.Request, origin string) bool {
	//if origin == "http://example.com" {
	//	return true
	//}
	return true
}

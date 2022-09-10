package api

import (
	"net/http"

	"github.com/claudemuller/fishwife/api/routes"
	"github.com/claudemuller/fishwife/internal/pkg"
)

type handlerFunc func(w http.ResponseWriter, req *http.Request, app pkg.AppState)

func getHandler(app pkg.AppState, f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "GET" {
			f(w, req, app)
			return
		}

		pkg.SendMethodNotSupported(w)
	}
}

func postHandler(app pkg.AppState, f handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			f(w, req, app)
			return
		}

		pkg.SendMethodNotSupported(w)
	}
}

func NewRouter(app pkg.AppState) *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/health", getHandler(app, routes.Health))
	r.HandleFunc("/reminders", getHandler(app, routes.GetReminders))

	return r
}

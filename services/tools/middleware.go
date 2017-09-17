package tools

import (
	"net/http"
	"github.com/pressly/chi/middleware"
	"dev.dohrm.com/git/rpg/portal/tools/errors"
	"runtime/debug"
	log "github.com/sirupsen/logrus"
)

var logCmd = log.WithFields(log.Fields{
	"module": "tools",
})

func ErrorHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logEntry := middleware.GetLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(err, debug.Stack())
				} else {
					debug.PrintStack()
				}
				switch err.(type) {
				case errors.Error:
					e := err.(errors.Error)
					http.Error(w, e.Message, e.Code)
				case error:
					e := err.(error)
					http.Error(w, e.Error(), http.StatusInternalServerError)
				default:
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

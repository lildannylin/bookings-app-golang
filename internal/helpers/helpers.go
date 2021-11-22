package helpers

import (
	"fmt"
	"github.com/lildannylin/bookings-app-golang/internal/config"
	"net/http"
	"runtime/debug"
)

var app *config.AppConfig

// Newhelpers sets up app config for helpers
func NewHelper(a *config.AppConfig) {
	app = a
}

func ClientError(w http.ResponseWriter, status int) {
	app.InforLog.Println("Client error with status of", status)
	http.Error(w, http.StatusText(status), status)
}

func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
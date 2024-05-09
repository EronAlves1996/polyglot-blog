package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (a *app) serverError(err error, w http.ResponseWriter) {
	trace := fmt.Sprintf("%s\n%s", err, debug.Stack())
	a.errorLog.Print(trace)
	http.Error(w, "internalServerError", 500)
}

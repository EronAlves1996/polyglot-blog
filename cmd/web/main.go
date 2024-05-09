package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type app struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Port where server gonna listen. Default: ':4000'")
	flag.Parse()

	app := app{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	app.infoLog.Printf("Listening on %s", *addr)
	err := srv.ListenAndServe()
	app.errorLog.Fatal(err)
}

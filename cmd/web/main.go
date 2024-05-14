package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/EronAlves1996/polyglot-blog/internal/models"
	_ "github.com/lib/pq"
)

type app struct {
	errorLog  *log.Logger
	infoLog   *log.Logger
	postModel *models.PostModel
}

func main() {
	addr := flag.String("addr", ":4000", "Port where server gonna listen. Default: ':4000'")
	dbConnString := flag.String("db-conn-string", "postgres//root:root@localhost:5432/blog_app", "Database connection String")
	flag.Parse()

	app := app{
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
		postModel: &models.PostModel{
			Db: createConnection(*dbConnString),
		},
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: app.errorLog,
	}

	app.infoLog.Printf("Listening on %s", *addr)
	if err := srv.ListenAndServe(); err != nil {
		app.errorLog.Fatal(err)
	}
}

func createConnection(connString string) *sql.DB {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected successfully to database")
	return db
}

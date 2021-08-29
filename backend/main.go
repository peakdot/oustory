package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/peakdot/oustory/config"
	"github.com/peakdot/oustory/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type contextKey string

const (
	contextKeyChosenProject = contextKey("chosenProject")
	contextKeyChosenBacklog = contextKey("chosenBacklog")
)

type application struct {
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	Location *time.Location
	DB       *gorm.DB
	Config   *config.Config
	Mode     string
}

func main() {
	mode := flag.String("mode", "prod", "Service mode. dev or prod.")
	flag.Parse()

	infoLog := log.New(os.Stdout, "[web]\tINFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[web]\tERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	cfgs, err := config.ReadConfig()
	if err != nil {
		errorLog.Fatal("failed to read config", err)
	}

	var cfg *config.Config
	switch *mode {
	case "dev", "prod":
		cfg = cfgs[*mode]
	default:
		errorLog.Fatal("invalid mode:", *mode)
	}

	db, err := gorm.Open(sqlite.Open("oustory.db"), &gorm.Config{})
	if err != nil {
		errorLog.Fatal("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(
		&entity.Project{},
		&entity.Backlog{},
		&entity.AcceptanceCriteria{},
		&entity.BacklogAssignee{},
		&entity.Subtask{},
		&entity.SubtaskAssignee{},
		&entity.Member{},
	)

	loc, _ := time.LoadLocation("Asia/Ulaanbaatar")

	// Web application section
	app := &application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
		Location: loc,
		DB:       db,
		Config:   cfg,
		Mode:     *mode,
	}

	srv := &http.Server{
		Addr:         app.Config.Server.Host + ":" + app.Config.Server.Port,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 5 * time.Minute,
	}

	infoLog.Printf("Starting server on http://%s:%s", app.Config.Server.Host, app.Config.Server.Port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

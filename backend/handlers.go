package main

import (
	"net/http"

	"github.com/peakdot/oustory/common/oapi"
	"github.com/peakdot/oustory/entity"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
	var projects []*entity.Project
	if result := app.DB.Find(&projects); result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, projects)
}

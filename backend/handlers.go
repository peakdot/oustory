package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/peakdot/oustory/common/oapi"
	"github.com/peakdot/oustory/entity"
	"gorm.io/gorm"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
	var projects []*entity.Project
	// Order please
	if result := app.DB.Find(&projects); result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}
	oapi.SendResp(w, projects)
}

func (app *application) getBacklogs(w http.ResponseWriter, r *http.Request) {
	project := r.Context().Value(contextKeyChosenProject).(*entity.Project)

	// Order please
	// Thinking about pagination. But i think it's not necessary right now.
	var backlogs []*entity.Backlog
	result := app.DB.Where("project_id=?", project.ID).Find(&backlogs)
	if result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, backlogs)
}

// Is it okay to be upsert? Argh...
func (app *application) addBacklog(w http.ResponseWriter, r *http.Request) {
	project := r.Context().Value(contextKeyChosenProject).(*entity.Project)

	data := new(entity.Backlog)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		app.InfoLog.Println(err)
		oapi.ClientError(w, http.StatusBadRequest)
		return
	}

	lastBacklog := new(entity.Backlog)
	result := app.DB.Where("project_id=?", project.ID).Last(lastBacklog)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// When not found, we rely on zero value.
		// We are getting last backlog to get backlog number.
		// Check not found only if necessary.
		oapi.ServerError(w, result.Error)
		return
	}

	// No validation. Who wants to break their own space?
	// Order goes up means high & urgent. So default is 0.
	backlog := &entity.Backlog{
		Type:      data.Type,
		Text:      data.Text,
		Number:    lastBacklog.Number + 1, // Relying on zero value here
		Order:     data.Order,
		Point:     data.Point,
		Status:    data.Status,
		ProjectID: uint(project.ID),
	}

	if err := app.DB.Create(&backlog).Error; err != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, backlog)
}

func (app *application) editBacklog(w http.ResponseWriter, r *http.Request) {
	backlog := r.Context().Value(contextKeyChosenBacklog).(*entity.Backlog)

	data := new(entity.Backlog)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		app.InfoLog.Println(err)
		oapi.ClientError(w, http.StatusBadRequest)
		return
	}

	// No validation. Same as addBacklog.
	// I think it only updates non zero fields
	if err := app.DB.Model(&backlog).Updates(entity.Backlog{
		Type:   data.Type,
		Text:   data.Text,
		Order:  data.Order,
		Point:  data.Point,
		Status: data.Status,
	}).Error; err != nil {
		oapi.ServerError(w, err)
		return
	}

	oapi.SendResp(w, backlog)
}

func (app *application) getSubtasks(w http.ResponseWriter, r *http.Request) {
	backlog := r.Context().Value(contextKeyChosenBacklog).(*entity.Backlog)

	var subtasks []*entity.Subtask
	result := app.DB.Where("backlog_id=?", backlog.ID).Find(&subtasks)
	if result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, subtasks)
}

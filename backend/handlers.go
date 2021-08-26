package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
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

func (app *application) getStories(w http.ResponseWriter, r *http.Request) {
	projectID, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if projectID <= 0 {
		oapi.ClientError(w, http.StatusBadRequest)
		return
	}

	var stories []*entity.UserStory
	result := app.DB.Where("project_id=?", projectID).Find(&stories)
	if result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, stories)
}

func (app *application) getSubtasks(w http.ResponseWriter, r *http.Request) {
	storyID, _ := strconv.Atoi(chi.URLParam(r, "storyID"))
	if storyID <= 0 {
		oapi.ClientError(w, http.StatusBadRequest)
		return
	}

	var subtasks []*entity.Subtask
	result := app.DB.Where("user_story_id=?", storyID).Find(&subtasks)
	if result.Error != nil {
		oapi.ServerError(w, result.Error)
		return
	}

	oapi.SendResp(w, subtasks)
}

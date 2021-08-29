package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/peakdot/oustory/common/oapi"
	"github.com/peakdot/oustory/entity"
	"gorm.io/gorm"
)

// No middlewares for now
func (app *application) setProjectCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		projectID, err := strconv.Atoi(chi.URLParam(r, "projectID"))
		if err != nil || projectID <= 0 {
			oapi.NotFound(w)
			return
		}

		project := new(entity.Project)
		if err := app.DB.First(&project, projectID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				oapi.NotFound(w)
			} else {
				oapi.ServerError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyChosenProject, project)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) setBacklogCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Middleware logic before executing given Handler
		backlogID, err := strconv.Atoi(chi.URLParam(r, "backlogID"))
		if err != nil || backlogID <= 0 {
			oapi.NotFound(w)
			return
		}

		backlog := new(entity.Backlog)
		if err := app.DB.First(&backlog, backlogID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				oapi.NotFound(w)
			} else {
				oapi.ServerError(w, err)
			}
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyChosenBacklog, backlog)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

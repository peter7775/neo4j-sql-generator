/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package web

import (
	"net/http"

	"github.com/gorilla/mux"

	"neo4j-sql-generator/internal/infrastructure/middleware"
	"neo4j-sql-generator/internal/interfaces/web/handlers"
)

// Router sets up HTTP routes
type Router struct {
	designHandler     *handlers.DesignHandler
	generationHandler *handlers.GenerationHandler
}

// NewRouter creates a new router
func NewRouter() *Router {
	return &Router{
		designHandler:     handlers.NewDesignHandler(),
		generationHandler: handlers.NewGenerationHandler(),
	}
}

// SetupRoutes configures all HTTP routes
func (ro *Router) SetupRoutes() http.Handler {
	r := mux.NewRouter()

	// Apply middleware
	r.Use(middleware.LoggingMiddleware)
	r.Use(middleware.CORSMiddleware)

	// API routes
	api := r.PathPrefix("/api").Subrouter()

	// Design routes
	api.HandleFunc("/designs", ro.designHandler.GetDesigns).Methods("GET")
	api.HandleFunc("/designs", ro.designHandler.CreateDesign).Methods("POST")
	api.HandleFunc("/designs/{id}", ro.designHandler.GetDesign).Methods("GET")
	api.HandleFunc("/designs/{id}", ro.designHandler.UpdateDesign).Methods("PUT")
	api.HandleFunc("/designs/{id}", ro.designHandler.DeleteDesign).Methods("DELETE")

	// Generation routes
	api.HandleFunc("/generate", ro.generationHandler.GenerateSQL).Methods("POST")
	api.HandleFunc("/analyze", ro.generationHandler.AnalyzeDatabase).Methods("POST")

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/static/")))

	return r
}

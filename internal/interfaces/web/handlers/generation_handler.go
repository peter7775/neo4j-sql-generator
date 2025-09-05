/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import "net/http"

// GenerationHandler handles SQL generation HTTP requests
type GenerationHandler struct {}

// NewGenerationHandler creates a new generation handler
func NewGenerationHandler() *GenerationHandler {
	return &GenerationHandler{}
}

// GenerateSQL handles POST /api/generate
func (h *GenerationHandler) GenerateSQL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// AnalyzeDatabase handles POST /api/analyze
func (h *GenerationHandler) AnalyzeDatabase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package queries

// ListDesignsQuery represents a list designs query
type ListDesignsQuery struct {}

// ListDesignsHandler handles list designs queries
type ListDesignsHandler struct {}

// Handle executes the list designs query
func (h *ListDesignsHandler) Handle(query ListDesignsQuery) error {
	return nil
}

/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package queries

// GetDesignQuery represents a get design query
type GetDesignQuery struct {}

// GetDesignHandler handles get design queries
type GetDesignHandler struct {}

// Handle executes the get design query
func (h *GetDesignHandler) Handle(query GetDesignQuery) error {
	return nil
}

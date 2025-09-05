/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package commands

// CreateDesignCommand represents a create design command
type CreateDesignCommand struct {}

// CreateDesignHandler handles create design commands
type CreateDesignHandler struct {}

// Handle executes the create design command
func (h *CreateDesignHandler) Handle(cmd CreateDesignCommand) error {
	return nil
}

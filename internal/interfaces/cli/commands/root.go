/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package commands

// RootCommand represents the base CLI command
type RootCommand struct {}

// NewRootCommand creates a new root CLI command
func NewRootCommand() *RootCommand {
	return &RootCommand{}
}

// Execute runs the root command
func (c *RootCommand) Execute() error {
	return nil
}

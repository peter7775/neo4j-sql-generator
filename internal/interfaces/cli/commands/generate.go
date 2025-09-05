/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package commands

// GenerateCommand represents the generate CLI command
type GenerateCommand struct {
	DesignFile string
	Target     string
	Output     string
}

// NewGenerateCommand creates a new generate CLI command
func NewGenerateCommand() *GenerateCommand {
	return &GenerateCommand{}
}

// Execute runs the generate command
func (c *GenerateCommand) Execute() error {
	return nil
}

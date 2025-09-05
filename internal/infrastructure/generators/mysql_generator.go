/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package generators

import "context"

// MySQLGenerator generates MySQL DDL scripts
type MySQLGenerator struct {}

// NewMySQLGenerator creates a new MySQL generator
func NewMySQLGenerator() *MySQLGenerator {
	return &MySQLGenerator{}
}

// GenerateSQL generates MySQL DDL from schema
func (g *MySQLGenerator) GenerateSQL(ctx context.Context) (string, error) {
	return "", nil
}

// ValidateGeneration validates MySQL generation parameters
func (g *MySQLGenerator) ValidateGeneration(ctx context.Context) error {
	return nil
}

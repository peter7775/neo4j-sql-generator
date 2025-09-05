/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package generators

import "context"

// PostgreSQLGenerator generates PostgreSQL DDL scripts
type PostgreSQLGenerator struct {}

// NewPostgreSQLGenerator creates a new PostgreSQL generator
func NewPostgreSQLGenerator() *PostgreSQLGenerator {
	return &PostgreSQLGenerator{}
}

// GenerateSQL generates PostgreSQL DDL from schema
func (g *PostgreSQLGenerator) GenerateSQL(ctx context.Context) (string, error) {
	return "", nil
}

// ValidateGeneration validates PostgreSQL generation parameters
func (g *PostgreSQLGenerator) ValidateGeneration(ctx context.Context) error {
	return nil
}

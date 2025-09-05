/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package generators

import "context"

// SQLiteGenerator generates SQLite DDL scripts
type SQLiteGenerator struct {}

// NewSQLiteGenerator creates a new SQLite generator
func NewSQLiteGenerator() *SQLiteGenerator {
	return &SQLiteGenerator{}
}

// GenerateSQL generates SQLite DDL from schema
func (g *SQLiteGenerator) GenerateSQL(ctx context.Context) (string, error) {
	return "", nil
}

// ValidateGeneration validates SQLite generation parameters
func (g *SQLiteGenerator) ValidateGeneration(ctx context.Context) error {
	return nil
}

/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package services

import "context"

// GeneratorService defines the SQL generator service interface
type GeneratorService interface {
	GenerateSQL(ctx context.Context) error
	ValidateGeneration(ctx context.Context) error
}

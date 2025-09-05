/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package generation

import "context"

// GenerationService provides SQL generation application services
type GenerationService struct {}

// GenerateSQL generates SQL from a design
func (gs *GenerationService) GenerateSQL(ctx context.Context) error {
	return nil
}

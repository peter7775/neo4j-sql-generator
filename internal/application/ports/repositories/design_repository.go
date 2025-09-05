/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package repositories

import "context"

// DesignRepository defines the design repository interface
type DesignRepository interface {
	Save(ctx context.Context) error
	FindByID(ctx context.Context, id string) error
	FindAll(ctx context.Context) error
	Delete(ctx context.Context, id string) error
}

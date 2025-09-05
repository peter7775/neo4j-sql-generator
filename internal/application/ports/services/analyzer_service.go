/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package services

import "context"

// AnalyzerService defines the Neo4j analyzer service interface
type AnalyzerService interface {
	AnalyzeDatabase(ctx context.Context) error
	ExtractSchema(ctx context.Context) error
}

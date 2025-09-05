/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package resolvers

import "context"

// GenerationResolver handles GraphQL generation operations
type GenerationResolver struct {}

// NewGenerationResolver creates a new generation resolver
func NewGenerationResolver() *GenerationResolver {
	return &GenerationResolver{}
}

// GenerateSchema generates SQL schema via GraphQL
func (r *GenerationResolver) GenerateSchema(ctx context.Context, designID string, target string) error {
	return nil
}

// AnalyzeDatabase analyzes Neo4j database via GraphQL
func (r *GenerationResolver) AnalyzeDatabase(ctx context.Context, uri string) error {
	return nil
}

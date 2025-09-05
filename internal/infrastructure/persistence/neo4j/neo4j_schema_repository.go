/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package neo4j

import "context"

// Neo4jSchemaRepository implements schema repository using Neo4j
type Neo4jSchemaRepository struct {}

// Save saves a schema to Neo4j
func (r *Neo4jSchemaRepository) Save(ctx context.Context) error {
	return nil
}

// FindByID finds a schema by ID in Neo4j
func (r *Neo4jSchemaRepository) FindByID(ctx context.Context, id string) error {
	return nil
}

// Delete deletes a schema from Neo4j
func (r *Neo4jSchemaRepository) Delete(ctx context.Context, id string) error {
	return nil
}

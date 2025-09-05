/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package neo4j

import "context"

// Neo4jDesignRepository implements design repository using Neo4j
type Neo4jDesignRepository struct {}

// Save saves a design to Neo4j
func (r *Neo4jDesignRepository) Save(ctx context.Context) error {
	return nil
}

// FindByID finds a design by ID in Neo4j
func (r *Neo4jDesignRepository) FindByID(ctx context.Context, id string) error {
	return nil
}

// FindAll finds all designs in Neo4j
func (r *Neo4jDesignRepository) FindAll(ctx context.Context) error {
	return nil
}

// Delete deletes a design from Neo4j
func (r *Neo4jDesignRepository) Delete(ctx context.Context, id string) error {
	return nil
}

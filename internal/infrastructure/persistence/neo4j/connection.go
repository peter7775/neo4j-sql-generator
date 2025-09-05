/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package neo4j

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

// Connection manages Neo4j database connections
type Connection struct {
	driver neo4j.Driver
}

// NewConnection creates a new Neo4j connection
func NewConnection() (*Connection, error) {
	return &Connection{}, nil
}

// Close closes the Neo4j connection
func (c *Connection) Close(ctx context.Context) error {
	return c.driver.Close()
}

// Session creates a new Neo4j session
func (c *Connection) Session() neo4j.Session {
	return c.driver.NewSession(neo4j.SessionConfig{})
}

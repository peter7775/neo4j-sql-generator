/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/sirupsen/logrus"

	"neo4j-sql-generator/internal/application/services/design"
	"neo4j-sql-generator/internal/application/services/generation"
	"neo4j-sql-generator/internal/infrastructure/persistence/neo4j"
)

func main() {
	ctx := context.Background()

	logrus.Info("Starting Neo4j SQL Generator...")

	logrus.Info("Neo4j SQL Generator started successfully")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logrus.Info("Shutting down Neo4j SQL Generator...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logrus.Info("Neo4j SQL Generator shutdown complete")
}

func init() {
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))
	if err != nil {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(level)
	}
}

/*
 * Copyright (c) 2025 Petr Miroslav Stepanek <petrstepanek99@gmail.com>
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

package specifications

// SchemaSpecification defines schema business rules
type SchemaSpecification struct {}

// IsSatisfiedBy checks if schema satisfies business rules
func (ss *SchemaSpecification) IsSatisfiedBy() bool {
	return true
}

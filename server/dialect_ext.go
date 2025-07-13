// Copyright 2012-2025 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

// MySQLDialectExt provides MySQL dialect extensions
type MySQLDialectExt struct {
	// MySQLDialect embedded field
}

// SetTableNameMapper sets the table name mapper for MySQL
func (d *MySQLDialectExt) SetTableNameMapper(mapper interface{}) {
	// Implementation for MySQL
}

// SqlServerDialectExt provides SQL Server dialect extensions
type SqlServerDialectExt struct {
	// SqlServerDialect embedded field
}

// SetTableNameMapper sets the table name mapper for SQL Server
func (d *SqlServerDialectExt) SetTableNameMapper(mapper interface{}) {
	// Implementation for SQL Server
}

// OracleDialectExt provides Oracle dialect extensions
type OracleDialectExt struct {
	// OracleDialect embedded field
}

// SetTableNameMapper sets the table name mapper for Oracle
func (d *OracleDialectExt) SetTableNameMapper(mapper interface{}) {
	// Implementation for Oracle
}

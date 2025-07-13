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

package migration

import (
	"database/sql"

	// Import gorp with the correct module path
	"github.com/go-gorp/gorp"
	// This is what the code was trying to import before
	// "github.com/go-gorp/gorp/v3"
	// "github.com/go-gorp/gorp/v3/dialect"
)

// Import re-exports the necessary types and functions from gorp
// to maintain API compatibility with the expected v3 imports

// DbMap is a re-export of gorp.DbMap
type DbMap struct {
	*gorp.DbMap
}

// NewDbMap creates a new DbMap
func NewDbMap(db *sql.DB, dialect gorp.Dialect) *DbMap {
	return &DbMap{
		DbMap: &gorp.DbMap{
			Db:      db,
			Dialect: dialect,
		},
	}
}

// SqlDialect provides access to gorp dialects
type SqlDialect struct{}

var Dialect SqlDialect

// PostgreSQL returns a PostgreSQL dialect
func (d SqlDialect) PostgreSQL() gorp.Dialect {
	return gorp.PostgresDialect{}
}

// MySQL returns a MySQL dialect
func (d SqlDialect) MySQL() gorp.Dialect {
	return gorp.MySQLDialect{}
}

// SQLite3 returns a SQLite3 dialect
func (d SqlDialect) SQLite3() gorp.Dialect {
	return gorp.SqliteDialect{}
}

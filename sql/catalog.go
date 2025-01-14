// Copyright 2021 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.



package sql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/dolthub/go-mysql-server/sql"
)



type Catalog interface {
	// AllDatabases returns all databases known to this catalog
	AllDatabases(ctx *Context) []Database

	// HasDB returns whether a db with the name given exists, case-insensitive
	HasDB(ctx *Context, db string) bool

	// Database returns the database with the name given, case-insensitive, or an error if it doesn't exist
	Database(ctx *Context, db string) (Database, error)

	// CreateDatabase creates a new database, or returns an error if the operation isn't supported or fails.
	CreateDatabase(ctx *Context, dbName string, collation CollationID) error

	// RemoveDatabase removes the  database named, or returns an error if the operation isn't supported or fails.
	RemoveDatabase(ctx *Context, dbName string) error

	// Table returns the table with the name given in the db with the name given
	Table(ctx *Context, dbName, tableName string) (Table, Database, error)

	// TableAsOf returns the table with the name given in the db with the name given, as of the given marker
	TableAsOf(ctx *Context, dbName, tableName string, asOf interface{}) (Table, Database, error)

	// Function returns the function with the name given, or sql.ErrFunctionNotFound if it doesn't exist
	Function(ctx *Context, name string) (Function, error)

	// RegisterFunction registers the functions given, adding them to the built-in functions.
	// Integrators with custom functions should typically use the FunctionProvider interface to register their functions.
	RegisterFunction(ctx *Context, fns ...Function)

	// LockTable locks the table named
	LockTable(ctx *Context, table string)

	// UnlockTables unlocks all tables locked by the session id given
	UnlockTables(ctx *Context, id uint32) error

	// Statistics returns a StatsReadWriter for saving and updating table statistics
	Statistics(ctx *Context) (StatsReadWriter, error)
}

// CatalogTable is a Table that depends on a Catalog.
type CatalogTable interface {
	Table

	// AssignCatalog assigns a Catalog to the table.
	AssignCatalog(cat Catalog) Table
}

func NewDbTable(db, table string) DbTable {
	return DbTable{Db: strings.ToLower(db), Table: strings.ToLower(table)}
}

type DbTable struct {
	Db    string
	Table string
}

func (dt *DbTable) String() string {
	if dt.Db == "" {
		return dt.Table
	}
	return fmt.Sprintf("%s.%s", dt.Db, dt.Table)
}

func (c *catalog) IndexByName(ctx Context, db, table, index string) (Index, error) {
	tbl, err := c.Table(ctx, db, table)
	if err != nil {
	return nil, err
	}
	index, ok := tbl.Indexes().GetByName(index)
	if !ok {
	return nil, sql.ErrIndexNotFound.New(index)
	}
	if statsIndex, ok := index.(StatisticsIndex); ok {
	return statsIndex, nil
	}
	return index, nil
	}
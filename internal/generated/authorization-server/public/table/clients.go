//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Clients = newClientsTable("public", "clients", "")

type clientsTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnString
	Name        postgres.ColumnString
	URL         postgres.ColumnString
	RedirectURI postgres.ColumnString
	Secret      postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type ClientsTable struct {
	clientsTable

	EXCLUDED clientsTable
}

// AS creates new ClientsTable with assigned alias
func (a ClientsTable) AS(alias string) *ClientsTable {
	return newClientsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ClientsTable with assigned schema name
func (a ClientsTable) FromSchema(schemaName string) *ClientsTable {
	return newClientsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ClientsTable with assigned table prefix
func (a ClientsTable) WithPrefix(prefix string) *ClientsTable {
	return newClientsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ClientsTable with assigned table suffix
func (a ClientsTable) WithSuffix(suffix string) *ClientsTable {
	return newClientsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newClientsTable(schemaName, tableName, alias string) *ClientsTable {
	return &ClientsTable{
		clientsTable: newClientsTableImpl(schemaName, tableName, alias),
		EXCLUDED:     newClientsTableImpl("", "excluded", ""),
	}
}

func newClientsTableImpl(schemaName, tableName, alias string) clientsTable {
	var (
		IDColumn          = postgres.StringColumn("id")
		NameColumn        = postgres.StringColumn("name")
		URLColumn         = postgres.StringColumn("url")
		RedirectURIColumn = postgres.StringColumn("redirect_uri")
		SecretColumn      = postgres.StringColumn("secret")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, URLColumn, RedirectURIColumn, SecretColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, URLColumn, RedirectURIColumn, SecretColumn}
	)

	return clientsTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		URL:         URLColumn,
		RedirectURI: RedirectURIColumn,
		Secret:      SecretColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
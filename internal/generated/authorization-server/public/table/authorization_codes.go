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

var AuthorizationCodes = newAuthorizationCodesTable("public", "authorization_codes", "")

type authorizationCodesTable struct {
	postgres.Table

	// Columns
	Code           postgres.ColumnString
	ClientID       postgres.ColumnString
	RedirectURI    postgres.ColumnString
	ExpirationTime postgres.ColumnTimestampz
	Scope          postgres.ColumnString
	Used           postgres.ColumnBool

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type AuthorizationCodesTable struct {
	authorizationCodesTable

	EXCLUDED authorizationCodesTable
}

// AS creates new AuthorizationCodesTable with assigned alias
func (a AuthorizationCodesTable) AS(alias string) *AuthorizationCodesTable {
	return newAuthorizationCodesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new AuthorizationCodesTable with assigned schema name
func (a AuthorizationCodesTable) FromSchema(schemaName string) *AuthorizationCodesTable {
	return newAuthorizationCodesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new AuthorizationCodesTable with assigned table prefix
func (a AuthorizationCodesTable) WithPrefix(prefix string) *AuthorizationCodesTable {
	return newAuthorizationCodesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new AuthorizationCodesTable with assigned table suffix
func (a AuthorizationCodesTable) WithSuffix(suffix string) *AuthorizationCodesTable {
	return newAuthorizationCodesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newAuthorizationCodesTable(schemaName, tableName, alias string) *AuthorizationCodesTable {
	return &AuthorizationCodesTable{
		authorizationCodesTable: newAuthorizationCodesTableImpl(schemaName, tableName, alias),
		EXCLUDED:                newAuthorizationCodesTableImpl("", "excluded", ""),
	}
}

func newAuthorizationCodesTableImpl(schemaName, tableName, alias string) authorizationCodesTable {
	var (
		CodeColumn           = postgres.StringColumn("code")
		ClientIDColumn       = postgres.StringColumn("client_id")
		RedirectURIColumn    = postgres.StringColumn("redirect_uri")
		ExpirationTimeColumn = postgres.TimestampzColumn("expiration_time")
		ScopeColumn          = postgres.StringColumn("scope")
		UsedColumn           = postgres.BoolColumn("used")
		allColumns           = postgres.ColumnList{CodeColumn, ClientIDColumn, RedirectURIColumn, ExpirationTimeColumn, ScopeColumn, UsedColumn}
		mutableColumns       = postgres.ColumnList{ClientIDColumn, RedirectURIColumn, ExpirationTimeColumn, ScopeColumn, UsedColumn}
	)

	return authorizationCodesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Code:           CodeColumn,
		ClientID:       ClientIDColumn,
		RedirectURI:    RedirectURIColumn,
		ExpirationTime: ExpirationTimeColumn,
		Scope:          ScopeColumn,
		Used:           UsedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

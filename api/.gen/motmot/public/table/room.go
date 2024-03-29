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

var Room = newRoomTable("public", "room", "")

type roomTable struct {
	postgres.Table

	// Columns
	ID          postgres.ColumnInteger
	Name        postgres.ColumnString
	Description postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type RoomTable struct {
	roomTable

	EXCLUDED roomTable
}

// AS creates new RoomTable with assigned alias
func (a RoomTable) AS(alias string) *RoomTable {
	return newRoomTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new RoomTable with assigned schema name
func (a RoomTable) FromSchema(schemaName string) *RoomTable {
	return newRoomTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new RoomTable with assigned table prefix
func (a RoomTable) WithPrefix(prefix string) *RoomTable {
	return newRoomTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new RoomTable with assigned table suffix
func (a RoomTable) WithSuffix(suffix string) *RoomTable {
	return newRoomTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newRoomTable(schemaName, tableName, alias string) *RoomTable {
	return &RoomTable{
		roomTable: newRoomTableImpl(schemaName, tableName, alias),
		EXCLUDED:  newRoomTableImpl("", "excluded", ""),
	}
}

func newRoomTableImpl(schemaName, tableName, alias string) roomTable {
	var (
		IDColumn          = postgres.IntegerColumn("id")
		NameColumn        = postgres.StringColumn("name")
		DescriptionColumn = postgres.StringColumn("description")
		allColumns        = postgres.ColumnList{IDColumn, NameColumn, DescriptionColumn}
		mutableColumns    = postgres.ColumnList{NameColumn, DescriptionColumn}
	)

	return roomTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}

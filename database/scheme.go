package database

import "github.com/xeviknal/background-jobs/models"

type tables []table
type table struct {
	Name  string
	Model interface{}
	Key   string
}

func CreateScheme() error {
	for _, table := range getTables() {
		dbmap.AddTableWithName(table.Model, table.Name).SetKeys(true, table.Key)
	}
	// TODO: Implement migration tool (rubenv/sql-migrate)
	return dbmap.CreateTablesIfNotExists()
}

func DropTables() error {
	return dbmap.DropTablesIfExists()
}

func getTables() tables {
	return tables{
		{
			"jobs",
			models.Job{},
			"Id",
		},
	}
}

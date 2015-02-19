package dat

import "database/sql"

// Executable is an object that can be queried.
type Executable interface {
	Exec() (sql.Result, error)
	Query() (*sql.Rows, error)
	QueryScan(destinations ...interface{}) error
	QuerySlice(dest interface{}) (int64, error)
	QueryStruct(dest interface{}) error
	QueryStructs(dest interface{}) (int64, error)
}

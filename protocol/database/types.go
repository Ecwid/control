package database

/*
Unique identifier of Database object.
*/
type DatabaseId string

/*
Database object.
*/
type Database struct {
	Id      DatabaseId `json:"id"`
	Domain  string     `json:"domain"`
	Name    string     `json:"name"`
	Version string     `json:"version"`
}

/*
Database error.
*/
type Error struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

type ExecuteSQLArgs struct {
	DatabaseId DatabaseId `json:"databaseId"`
	Query      string     `json:"query"`
}

type ExecuteSQLVal struct {
	ColumnNames []string      `json:"columnNames,omitempty"`
	Values      []interface{} `json:"values,omitempty"`
	SqlError    *Error        `json:"sqlError,omitempty"`
}

type GetDatabaseTableNamesArgs struct {
	DatabaseId DatabaseId `json:"databaseId"`
}

type GetDatabaseTableNamesVal struct {
	TableNames []string `json:"tableNames"`
}

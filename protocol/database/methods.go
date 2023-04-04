package database

import (
	"github.com/ecwid/control/protocol"
)

/*
Disables database tracking, prevents database events from being sent to the client.
*/
func Disable(c protocol.Caller) error {
	return c.Call("Database.disable", nil, nil)
}

/*
Enables database tracking, database events will now be delivered to the client.
*/
func Enable(c protocol.Caller) error {
	return c.Call("Database.enable", nil, nil)
}

/*
 */
func ExecuteSQL(c protocol.Caller, args ExecuteSQLArgs) (*ExecuteSQLVal, error) {
	var val = &ExecuteSQLVal{}
	return val, c.Call("Database.executeSQL", args, val)
}

/*
 */
func GetDatabaseTableNames(c protocol.Caller, args GetDatabaseTableNamesArgs) (*GetDatabaseTableNamesVal, error) {
	var val = &GetDatabaseTableNamesVal{}
	return val, c.Call("Database.getDatabaseTableNames", args, val)
}

package database

import (
	"database/sql"
	"fmt"

	config "../config"
	//3rd party package
	_ "github.com/go-sql-driver/mysql"
)

//Repositries :Collection of DB connections
type Repositries struct {
	_db [eMax]*sql.DB
}

//Instance :
var Instance *Repositries = nil

//Connect :Connect to database
func Connect(gameDB, logDB config.DBSetting) error {

	if Instance == nil {
		Instance = new(Repositries)

		connectionString := fmt.Sprintf("%s:%s@%s/%s", gameDB.DBUser, gameDB.DBPassword, gameDB.DB, gameDB.DBSchema)
		db1, err := sql.Open("mysql", connectionString)
		if err != nil {
			return err
		}

		connectionString = fmt.Sprintf("%s:%s@%s/%s", logDB.DBUser, logDB.DBPassword, logDB.DB, logDB.DBSchema)
		db2, err := sql.Open("mysql", connectionString)
		if err != nil {
			return err
		}

		db1.SetMaxIdleConns(MaxIdleConnection)
		db1.SetMaxOpenConns(MaxOpenConnection)
		db2.SetMaxIdleConns(MaxIdleConnection)
		db2.SetMaxOpenConns(MaxOpenConnection)

		Instance._db[eGameDB] = db1
		Instance._db[eLogDB] = db2
	}
	return nil
}

//Close :Close all connections
func Close() {
	if Instance != nil {
		for _, db := range Instance._db {
			db.Close()
		}
	}
}

package database

import (
	"fmt"

	config "../config"
	xorm "github.com/go-xorm/xorm"
)

//ORMEngines :Collection of DB connections
type ORMEngines struct {
	_engine [eMax]*xorm.Engine
}

//Instance2 :
var Instance2 *ORMEngines = nil

//CreateEngine :Connect to database by orm
func CreateEngine(gameDB, logDB config.DBSetting) error {

	if Instance2 == nil {
		Instance2 = new(ORMEngines)

		connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", gameDB.DBUser, gameDB.DBPassword, gameDB.DB, gameDB.DBSchema)
		engine1, err := xorm.NewEngine("mysql", connectionString)
		if err != nil {
			return err
		}

		connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", logDB.DBUser, logDB.DBPassword, logDB.DB, logDB.DBSchema)
		engine2, err := xorm.NewEngine("mysql", connectionString)
		if err != nil {
			return err
		}

		engine1.SetMaxIdleConns(MaxIdleConnection)
		engine1.SetMaxOpenConns(MaxOpenConnection)
		engine2.SetMaxIdleConns(MaxIdleConnection)
		engine2.SetMaxOpenConns(MaxOpenConnection)

		Instance2._engine[eGameDB] = engine1
		Instance2._engine[eLogDB] = engine2
	}

	return nil
}

//CloseEngine :Close all connections
func CloseEngine() {
	if Instance2 != nil {
		for _, engine := range Instance2._engine {
			engine.Close()
		}
	}
}

//SyncDBTables :Create/Update all the database tables.
func SyncDBTables() error {
	if Instance2 != nil {
		err := Instance2._engine[eGameDB].Sync2(new(Team))
		//...
		return err
	}

	return fmt.Errorf("Xorm instance is not created yet.")
}

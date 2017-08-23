package database

import "time"

const (
	eGameDB = 0 //Index of db array
	eLogDB  = 1 //Index of db array
	eMax    = 2 //Max
)

//MaxIdleConnection :
const MaxIdleConnection = 100

//MaxOpenConnection :Max connected connection
const MaxOpenConnection = 100

//Team :DB ORM Table
type Team struct {
	Owner    int64     `xorm:"'Owner' pk index(OwnerCardID)"`
	TeamID   int       `xorm:"'TeamID' pk"`
	Position int       `xorm:"'Position' pk"`
	CardID   int       `xorm:"'CardID' NOT NULL index(OwnerCardID)"`
	Name     string    `xorm:"'Name' NOT NULL varchar(50) index"`
	Created  time.Time `xorm:"'Created' NOT NULL"`
	Updated  time.Time `xorm:"'Updated' NOT NULL"`
}

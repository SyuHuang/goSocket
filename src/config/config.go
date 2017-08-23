package config

import (
	"encoding/json"
	"io/ioutil"
)

//Instance :
var Instance ServerSetting
var initialized = false

//ServerSetting :
type ServerSetting struct {
	ID      int    `json:"id"`
	Address string `json:"address"`
	Port    int    `json:"port"`

	GameDB DBSetting `json:"gamedb"`
	LogDB  DBSetting `json:"logdb"`
}

//DBSetting :
type DBSetting struct {
	DB         string `json:"db"`
	DBUser     string `json:"dbuser"`
	DBPassword string `json:"dbpassword"`
	DBSchema   string `json:"dbschema"` 
}

//ReadSetting :Read server setting to struct
func ReadSetting() (*ServerSetting, error) {
	if initialized == false {
		config, err := ioutil.ReadFile("config.json")
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(config, &Instance)
		if err != nil {
			return nil, err
		}
		initialized = true
	}
	return &Instance, nil
}

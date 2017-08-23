package main

import (
	"fmt"
	"os"
	"time"

	client "./client"
	config "./config"
	db "./database"
	proto "./proto"
)

func main() {

	//Get server setting.
	setting, err := config.ReadSetting()
	CheckError(err)

	// //Connect to database.
	// err = db.Connect(setting.GameDB, setting.LogDB)
	// CheckError(err)
	// defer db.Close()

	//Connect to database with orm.
	db.CreateEngine(setting.GameDB, setting.LogDB)
	CheckError(err)
	defer db.CloseEngine()

	//Create/Update all database tables.
	db.SyncDBTables()
	CheckError(err)

	//Regist all protocols.
	proto.InitProtocol()

	//Start server listening...
	ServerStart(setting.Address, setting.Port)

	//Create test client
	for i := 0; i < 1; i++ {
		player := client.NewPlayer()
		go player.Start()
	}

	//Game logic loop
	for {

		time.Sleep(10000 * time.Millisecond)

		fmt.Println("Now => " + time.Now().String())
	}
}

//CheckError :Print the err and exit program.
func CheckError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "[Server Info] Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

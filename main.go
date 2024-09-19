package main

import (
	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/serve"
)

func main(){
	var configApp,err = config.Load()
	if err != nil {
		panic(err)
	}
	database.Connect(configApp.Database.Uri)
	serve.Run()
}
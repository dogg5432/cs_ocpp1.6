package main

import (
	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/serve"
)

func main(){
	config.Load()
	database.Connect(config.Config.Database.Uri)
	serve.Run()
}
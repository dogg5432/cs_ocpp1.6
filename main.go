package main

import (
	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/serve"
)

func main(){
	if err := config.Load(); err != nil {
		panic(err)
	}
	if err := database.Connect(); err != nil {
		panic(err)
	}
	serve.Run()
}
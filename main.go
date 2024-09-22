package main

import (
	"github.com/dogg5432/central_charger/config"
	"github.com/dogg5432/central_charger/database"
	"github.com/dogg5432/central_charger/serve"
)

func main(){
	err := config.Load()
	if err != nil {
		panic(err)
	}
	err = database.Connect()
	if err != nil {
		panic(err)
	}
	serve.Run()
}
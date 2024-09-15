package main

import (
	"github.com/dogg5432/central_charger/serve"
	"github.com/dogg5432/central_charger/config"
)

func main(){
	config.Load()
	serve.Run()
}
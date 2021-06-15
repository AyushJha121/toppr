package main

import (
	"api-final/config"
	"api-final/handlers"
)

func main() {
	//Initialize My DB
	config.Init()
	defer config.Db.Close()
	//Intialize Router and Handle CRUD Operations
	handlers.Init()
}

package main

import (
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/database"
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/router"
)

func main() {
	database.InitDatabase()

	r := router.SetupRouter()

	r.Run(":8080")
}

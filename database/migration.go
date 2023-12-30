package database

import (
	"github.com/mochachmar/pbi-btpns-fd-mochachmar/app"
)

func MigrateTables() {
	DB.AutoMigrate(&app.User{}, &app.Photo{})
}

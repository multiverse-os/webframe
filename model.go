package maglev

import database "github.com/multiverse-os/maglev/database"

type Model struct {
	Framework

	Database *database.Database

	Name string

	BeforeCreate []func()
	AfterCreate  []func()

	BeforeSave []func()
	AfterSave  []func()
}

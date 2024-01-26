package webkit

import database "github.com/multiverse-os/webkit/database"

type Model struct {
	Framework *Framework

	Database *database.Database

	Name string

	BeforeCreate []func()
	AfterCreate  []func()

	BeforeSave []func()
	AfterSave  []func()
}

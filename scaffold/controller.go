package main

type ControllerData struct {
	Name string
	// TODO: As job object is created we fill out the data
	// and use that to build the template. This functionality
	// is itended to reflect the Rails `rails generate` functionality
	// which allows you to define aspects about the object  and that
	// is used to generate the corresponding template
}

func ControllerTemplate() string {
	return `package controllers

import (
	"fmt"
)

func {{ .Name }}Controller() {

}

`
}

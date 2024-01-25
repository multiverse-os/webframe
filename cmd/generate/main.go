package main

import (
	"fmt"
)

type TemplateType int

// TODO: Need a compound object like the generic scaffold provided that
// alloweddefinition of an object and coresponding CRUD (template for each
// component in MVC) was generated
const (
	Model TemplateType = iota
	Controller
	View
	Job
)

func (self TemplateType) String() string {
	switch self {
	case Model:
		return "model"
	case Controller:
		return "controller"
	case View:
		return "view"
	case Job:
		return "job"
	default:
		return ""
	}
}

func RenderTemplate(templateType TemplateType, data map[string]string) string {
	fmt.Println("functionality not yet implated")
	fmt.Println("should have generated a [", templateType.String(), "] template and installed it in the corresponding project folder")
}

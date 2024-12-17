package main

import (
	"groupie-tracker/routes"
	temp "groupie-tracker/templates"
)

func main() {
	temp.InitTemplates()
	routes.InitServe()
}

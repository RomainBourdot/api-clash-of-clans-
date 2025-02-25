package temp

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func InitTemplates() {
	temp, tempErr := template.ParseGlob("./templates/*.html")
	if tempErr != nil {
		fmt.Printf("Erreur Template - Une erreur lors du chargement des template \n message d'erreur : %v\n", tempErr.Error())
		os.Exit(1)
	}
	Temp = temp
}

package templates

import (
	"html/template"
	"log"
)

var Tpl *template.Template

func InitTemplates() {
	var err error
	Tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}
}

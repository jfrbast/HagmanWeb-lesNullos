package main

import (
	"Hangmanweb/pages"
	"Hangmanweb/templates"
	"Hangmanweb/utils"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var err error
	utils.Mots, err = utils.LireMots("mots.txt")
	if err != nil {
		log.Fatal("Erreur lors de la lecture des mots :", err)
	}
	templates.InitTemplates()
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/play", pages.PlayPage)
	http.HandleFunc("/end", pages.EndPage)
	http.HandleFunc("/scores", pages.ScoresPage)
	http.HandleFunc("/propos", pages.ProposPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

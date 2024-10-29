package pages

import (
	"Hangmanweb/game"
	"Hangmanweb/utils"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var mots []string
var tpl *template.Template

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pseudo := r.FormValue("pseudo")
		difficulte := r.FormValue("difficulte")

		if !utils.ValiderPseudo(pseudo) {
			tpl.ExecuteTemplate(w, "index", "Pseudo invalide. Seuls les lettres, chiffres, _ et - sont autorisés.")
			return
		}

		utils.Session = game.NouvellePartie(mots, difficulte)
		utils.Session.Pseudo = pseudo

		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index", nil)
}

func PlayPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		val := r.FormValue("value")
		utils.CheckValue(val)
	}
	if utils.Session.EstTermine() {
		utils.EnJeu = false
		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "game", utils.Session)
}

func EndPage(w http.ResponseWriter, r *http.Request) {
	message := ""
	if !strings.Contains(utils.Session.MotAffiche, "_") {
		message = "Félicitations, " + utils.Session.Pseudo + ", vous avez gagné !"
	} else {
		message = "Dommage, " + utils.Session.Pseudo + ". Vous avez perdu. Le mot était : " + utils.Session.MotATrouver
	}

	utils.EnregistrerScore()

	tpl.ExecuteTemplate(w, "end", message)
}

func ScoresPage(w http.ResponseWriter, r *http.Request) {
	scores, err := utils.LireScores()
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier des scores :", err)
	}

	tpl.ExecuteTemplate(w, "scores", scores)
}
func ProposPage(w http.ResponseWriter, r *http.Request) {
	err := tpl.ExecuteTemplate(w, "propos", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur lors de l'affichage de la page", http.StatusInternalServerError)
	}
}

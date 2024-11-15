package pages

import (
	"Hangmanweb/templates"
	"Hangmanweb/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

var win bool

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		pseudo := r.FormValue("pseudo")
		difficulte := r.FormValue("difficulte")

		if !utils.ValiderPseudo(pseudo) {
			err := templates.Tpl.ExecuteTemplate(w, "index", "Pseudo invalide. Seuls les lettres, chiffres, _ et - sont autorisés.")
			if err != nil {
				log.Println("Erreur lors de l'exécution du template :", err)
				http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
			}
			return
		}

		utils.Session = utils.NouvellePartie(utils.Mots, difficulte)
		utils.Session.Pseudo = pseudo
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	err := templates.Tpl.ExecuteTemplate(w, "index", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func PlayPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		val := r.FormValue("value")
		utils.Check(val)
	}
	if utils.EnJeu == false {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	if utils.Session.EstTermine() {
		utils.EnJeu = false
		time.Sleep(2 * time.Second)
		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}
	err := templates.Tpl.ExecuteTemplate(w, "game", utils.Session)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}
func EndPage(w http.ResponseWriter, r *http.Request) {
	message := ""
	if !strings.Contains(utils.Session.MotAffiche, "_") {

		win = true
		message = utils.PhrasesALeatoire(win)

	} else {
		win = false
		message = utils.PhrasesALeatoire(win) + " Le mot était : " + utils.Session.MotATrouver

	}

	utils.EnregistrerScore()

	err := templates.Tpl.ExecuteTemplate(w, "end", message)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func ScoresPage(w http.ResponseWriter, r *http.Request) {
	scores, err := utils.LireScores()
	if err != nil {
		log.Println("Erreur lors de la lecture du fichier des scores :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
		return
	}

	err = templates.Tpl.ExecuteTemplate(w, "scores", scores)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

func ProposPage(w http.ResponseWriter, r *http.Request) {
	err := templates.Tpl.ExecuteTemplate(w, "propos", nil)
	if err != nil {
		log.Println("Erreur lors de l'exécution du template :", err)
		http.Error(w, "Erreur interne du serveur", http.StatusInternalServerError)
	}
}

package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

type GameSession struct {
	MotATrouver     string
	LettresEssayees []string
	EssaisRestants  int
	MotAffiche      string
}

var session GameSession
var mots []string

func LireMotsDepuisFichier(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}
	mots := strings.Split(string(contenu), "\n")
	return mots, nil
}

func NouvellePartie(mots []string) GameSession {

	mot := mots[rand.Intn(len(mots))]
	return GameSession{
		MotATrouver:     strings.TrimSpace(mot),
		LettresEssayees: []string{},
		EssaisRestants:  10,
		MotAffiche:      genererMotAffiche(mot, []string{}),
	}
}

func genererMotAffiche(mot string, lettresEssayees []string) string {
	affichage := ""
	for _, lettre := range mot {
		if contains(lettresEssayees, string(lettre)) {
			affichage += string(lettre) + " "
		} else {
			affichage += "_ "
		}
	}
	return affichage
}

func (g *GameSession) DevinerLettre(lettre string) bool {
	if !contains(g.LettresEssayees, lettre) {
		g.LettresEssayees = append(g.LettresEssayees, lettre)
		if !strings.Contains(g.MotATrouver, lettre) {
			g.EssaisRestants--
			return false
		}
		g.MotAffiche = genererMotAffiche(g.MotATrouver, g.LettresEssayees)
	}
	return true
}

func (g *GameSession) EstTerminee() bool {
	return g.EssaisRestants <= 0 || !strings.Contains(g.MotAffiche, "_")
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		session = NouvellePartie(mots)

		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index.html", nil)
}

func playPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lettre := r.FormValue("lettre")

		session.DevinerLettre(lettre)
	}

	if session.EstTerminee() {
		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "game.html", session)
}

func endPage(w http.ResponseWriter, r *http.Request) {
	message := ""
	if !strings.Contains(session.MotAffiche, "_") {
		message = "Félicitations, vous avez gagné !"
	} else {
		message = "Dommage, vous avez perdu. Le mot était : " + session.MotATrouver
	}
	tpl.ExecuteTemplate(w, "end.html", message)
}

var tpl *template.Template

func main() {

	var err error
	mots, err = LireMotsDepuisFichier("mots.txt")
	if err != nil {
		log.Fatal("Erreur lors de la lecture des mots :", err)
	}

	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playPage)
	http.HandleFunc("/end", endPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

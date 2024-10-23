package main

import (
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type GameSession struct {
	Pseudo          string
	MotATrouver     string
	LettresEssayees []string
	EssaisRestants  int
	MotAffiche      string
	Difficulte      string
	Mob             string
}

var session GameSession
var mots []string

func ValiderPseudo(pseudo string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", pseudo)
	return matched
}

func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}
	mots := strings.Split(string(contenu), "\n")
	return mots, nil
}

func NouvellePartie(mots []string, difficulte string) GameSession {
	mot := mots[rand.Intn(len(mots))]
	essais := determinerEssais(difficulte)

	mob := assignerMob(difficulte)

	return GameSession{
		MotATrouver:     strings.TrimSpace(mot),
		LettresEssayees: []string{},
		EssaisRestants:  essais,
		MotAffiche:      genererMotAffiche(mot, []string{}),
		Difficulte:      difficulte,
		Mob:             mob,
	}
}

func determinerEssais(difficulte string) int {
	switch difficulte {
	case "Normal":
		return 8
	case "Difficile":
		return 6
	case "Extreme":
		return 4
	case "Nullos":
		return 200000
	case "Entrainement":
		return 12
	default:
		return 0
	}
}

func assignerMob(difficulte string) string {
	if difficulte == "Normal" {
		mobs := []string{"Zombie", "Squelette", "Piglin", "Cochon"}
		return mobs[rand.Intn(len(mobs))]
	} else if difficulte == "Difficile" {
		return "Creeper"
	} else if difficulte == "Extreme" {
		return "Slime"
	} else if difficulte == "Nullos" {
		return "Enderman"
	} else if difficulte == "Entrainement" {
		return "Armorstand"
	}
	return "Inconnu"
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

func (g *GameSession) TryLetter(lettre string) bool {
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

func (g *GameSession) EstTermine() bool {
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
		pseudo := r.FormValue("pseudo")
		difficulte := r.FormValue("difficulte")

		if !ValiderPseudo(pseudo) {
			tpl.ExecuteTemplate(w, "index", "Pseudo invalide. Seuls les lettres, chiffres, _ et - sont autorisés.")
			return
		}

		session = NouvellePartie(mots, difficulte)
		session.Pseudo = pseudo

		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "index", nil)
}

func playPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		lettre := r.FormValue("lettre")
		session.TryLetter(lettre)
	}

	if session.EstTermine() {
		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "game", session)
}

func endPage(w http.ResponseWriter, r *http.Request) {
	message := ""
	if !strings.Contains(session.MotAffiche, "_") {
		message = "Félicitations, " + session.Pseudo + ", vous avez gagné !"
	} else {
		message = "Dommage, " + session.Pseudo + ". Vous avez perdu. Le mot était : " + session.MotATrouver
	}

	score := session.Pseudo + " - Mot: " + session.MotATrouver + " - Essais restants: " + fmt.Sprint(session.EssaisRestants)
	enregistrerScore(score)

	tpl.ExecuteTemplate(w, "end", message)
}

func enregistrerScore(score string) {
	f, err := os.OpenFile("Tabscore.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier des scores :", err)
	}
	defer f.Close()

	if _, err = f.WriteString(score + "\n"); err != nil {
		log.Fatal("Erreur lors de l'écriture du score :", err)
	}
}

func scoresPage(w http.ResponseWriter, r *http.Request) {
	contenu, err := os.ReadFile("Tabscore.txt")
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier des scores :", err)
	}
	tpl.ExecuteTemplate(w, "scores", string(contenu))
}


var tpl *template.Template

func main() {
	var err error
	mots, err = LireMots("mots.txt")
	if err != nil {
		log.Fatal("Erreur lors de la lecture des mots :", err)
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}

	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playPage)
	http.HandleFunc("/end", endPage)
	http.HandleFunc("/scores", scoresPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*






func CheckValue(str string) (bool, string) {
	strings.ToLower(str)
	match := regexp.MustCompile("^[a-zA-Z]$").MatchString(str)
	if match {
		if len(str) == 1 {
			return session.TryLetter(str)
		}
		if len(str) > 1 {
			session.TryMot(str)
		}
	}
	return false, "Opus valeur invalide..."
}
func playPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		value := r.FormValue("value")
		fmt.Println(CheckValue(value))
	}

	if session.EstTermine() {
		enJeu = false
		http.Redirect(w, r, "/end", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "game", session)
}

func endPage(w http.ResponseWriter, r *http.Request) {
	if enJeu {
		http.Redirect(w, r, "/play", http.StatusSeeOther)
		return
	}

	message := ""
	if !strings.Contains(session.MotAffiche, "_") {
		message = "Félicitations, " + session.Pseudo + ", vous avez gagné !"
	} else {
		message = "Dommage, " + session.Pseudo + ". Vous avez perdu. Le mot était : " + session.MotATrouver
	}

	enregistrerScore()
	tpl.ExecuteTemplate(w, "end", message)
}

func enregistrerScore() {
	f, err := os.OpenFile("Tabscore.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier des scores :", err)
	}
	defer f.Close()

	date := time.Now().Format("02/01/2006")

	score := Score{
		Pseudo:         session.Pseudo,
		Difficulte:     session.Difficulte,
		MotATrouver:    session.MotATrouver,
		CoupsJoues:     len(session.LettresEssayees),
		EssaisRestants: session.EssaisRestants,
		Date:           date,
	}

	scoreComplet := fmt.Sprintf("%s - Difficulté: %s - Mot: %s - Coups joués: %d - Essais restants: %d - Date: %s",
		score.Pseudo, score.Difficulte, score.MotATrouver, score.CoupsJoues, score.EssaisRestants, score.Date)

	if _, err = f.WriteString(scoreComplet + "\n"); err != nil {
		log.Fatal("Erreur lors de l'écriture du score :", err)
	}
}

func lireScores() ([]Score, error) {
	contenu, err := os.ReadFile("Tabscore.txt")
	if err != nil {
		return nil, err
	}

	lignes := strings.Split(string(contenu), "\n")
	var scores []Score
	for _, ligne := range lignes {
		if ligne == "" {
			continue
		}

		var score Score

		_, err := fmt.Sscanf(ligne, "%s - Difficulté: %s - Mot: %s - Coups joués: %d - Essais restants: %d - Date: %s",
			&score.Pseudo, &score.Difficulte, &score.MotATrouver, &score.CoupsJoues, &score.EssaisRestants, &score.Date)
		if err != nil {
			log.Println("Erreur lors de l'extraction du score :", err)
			continue
		}
		scores = append(scores, score)
	}

	return scores, nil
}

func scoresPage(w http.ResponseWriter, r *http.Request) {
	scores, err := lireScores()
	if err != nil {
		log.Fatal("Erreur lors de la lecture du fichier des scores :", err)
	}

	tpl.ExecuteTemplate(w, "scores", scores)
}

func proposPage(w http.ResponseWriter, r *http.Request) {
	scores, err := lireScores()
	if err != nil {
		log.Fatal("Erreur lors de la lecture du à propos :", err)
	}

	tpl.ExecuteTemplate(w, "propos", scores)
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
	http.HandleFunc("/propos", proposPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
*/

package main

import (
	"Hangmanweb/pages"
	"Hangmanweb/utils"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var mots []string

var tpl *template.Template

func main() {
	var err error
	mots, err = utils.LireMots("mots.txt")
	if err != nil {
		log.Fatal("Erreur lors de la lecture des mots :", err)
	}

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	tpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Erreur lors du chargement des templates :", err)
	}

	http.HandleFunc("/", pages.HomePage)
	http.HandleFunc("/play", pages.PlayPage)
	http.HandleFunc("/end", pages.EndPage)
	http.HandleFunc("/scores", pages.ScoresPage)
	http.HandleFunc("/propos", pages.ProposPage)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

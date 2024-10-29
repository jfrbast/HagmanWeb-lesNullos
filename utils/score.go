package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func EnregistrerScore() {
	f, err := os.OpenFile("Tabscore.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal("Erreur lors de l'ouverture du fichier des scores :", err)
	}
	defer f.Close()

	date := time.Now().Format("02/01/2006")

	score := Score{
		Pseudo:         Session.Pseudo,
		Difficulte:     Session.Difficulte,
		MotATrouver:    Session.MotATrouver,
		CoupsJoues:     len(Session.LettresEssayees),
		EssaisRestants: Session.EssaisRestants,
		Date:           date,
	}

	scoreComplet := fmt.Sprintf("%s - Difficulté: %s - Mot: %s - Coups joués: %d - Essais restants: %d - Date: %s",
		score.Pseudo, score.Difficulte, score.MotATrouver, score.CoupsJoues, score.EssaisRestants, score.Date)

	if _, err = f.WriteString(scoreComplet + "\n"); err != nil {
		log.Fatal("Erreur lors de l'écriture du score :", err)
	}
}
func LireScores() ([]Score, error) {
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
func LireMots(nomFichier string) ([]string, error) {
	contenu, err := os.ReadFile(nomFichier)
	if err != nil {
		return nil, err
	}
	Mots := strings.Split(string(contenu), "\n")
	return Mots, nil
}

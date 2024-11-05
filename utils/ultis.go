package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
)

type GameSession struct {
	Pseudo          string
	MotATrouver     string
	LettresEssayees []string
	MotEssayes      []string
	EssaisRestants  int
	MotAffiche      string
	Difficulte      string
	Mob             string
}

type Score struct {
	Pseudo         string
	Difficulte     string
	MotATrouver    string
	CoupsJoues     int
	EssaisRestants int
	Date           string
}

var Session GameSession
var Mots []string
var EnJeu bool

func ValiderPseudo(pseudo string) bool {
	matched, _ := regexp.MatchString("^[a-zA-Z0-9_-]+$", pseudo)
	return matched
}

func GenererMotAffiche(mot string, lettresEssayees []string) string {
	affichage := ""
	for _, lettre := range mot {
		if Contains(lettresEssayees, string(lettre)) {
			affichage += string(lettre) + " "
		} else {
			affichage += "_ "
		}
	}
	return affichage
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
func (g *GameSession) EstTermine() bool {
	return g.EssaisRestants <= 0 || !strings.Contains(g.MotAffiche, "_")
}

func NouvellePartie(Mots []string, difficulte string) GameSession {
	EnJeu = true
	fmt.Println(len(Mots))
	mot := Mots[rand.Intn(len(Mots))]
	essais := DeterminerEssais(difficulte)
	mob := AssignerMob(difficulte)

	return GameSession{
		MotATrouver:     strings.TrimSpace(mot),
		LettresEssayees: []string{},
		EssaisRestants:  essais,
		MotAffiche:      GenererMotAffiche(mot, []string{}),
		Difficulte:      difficulte,
		Mob:             mob,
	}
}

func IsAlpha(str string) bool {
	match := regexp.MustCompile("^[a-zA-Z]+$").MatchString(str)
	return match

}

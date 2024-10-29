package utils

import (
	"regexp"
	"strings"
)

/*
	func (g *GameSession) TryLetter(lettre string) bool {
		strings.ToLower(lettre)
		if !Contains(g.LettresEssayees, lettre) {
			g.LettresEssayees = append(g.LettresEssayees, lettre)
			if !strings.Contains(g.MotATrouver, lettre) {
				g.EssaisRestants--
				return false
			}
			g.MotAffiche = GenererMotAffiche(g.MotATrouver, g.LettresEssayees)
		}
		return true
	}
*/
func (g *GameSession) TryLetter(lettre string) (bool, string) {
	strings.ToLower(lettre)
	if !Contains(g.LettresEssayees, lettre) {
		g.LettresEssayees = append(g.LettresEssayees, lettre)
		if !strings.Contains(g.MotATrouver, lettre) {
			g.EssaisRestants--
			return false, lettre

		}
		g.MotAffiche = GenererMotAffiche(g.MotATrouver, g.LettresEssayees)
	}
	return true, lettre
}

func (g *GameSession) TryMot(mot string) (bool, string) {
	strings.ToLower(mot)
	if !Contains(g.MotEssayes, mot) {
		g.MotEssayes = append(g.MotEssayes, mot)
		if strings.TrimSpace(mot) != g.MotATrouver {
			g.EssaisRestants = g.EssaisRestants - 2
			return false, mot

		}
		g.MotAffiche = mot
	}
	return true, mot
}
func CheckValue(str string) (bool, string) {
	strings.ToLower(str)
	match := regexp.MustCompile("^[a-zA-Z]$").MatchString(str)
	if match {
		if len(str) == 1 {
			return Session.TryLetter(str)
		}
		if len(str) > 1 {
			Session.TryMot(str)
		}
	}
	return false, "Opus valeur invalide..."
}

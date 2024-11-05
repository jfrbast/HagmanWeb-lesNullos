package utils

import (
	"strings"
)

func (g *GameSession) TryLetter(lettre string) (bool, string) {
	lettre = strings.ToLower(lettre)
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
	mot = strings.ToLower(mot)
	if !Contains(g.MotEssayes, mot) {
		if strings.TrimSpace(mot) != g.MotATrouver {
			g.EssaisRestants -= 2
			return false, mot
		}
		g.MotAffiche = mot
	}
	return true, mot
}

func CheckValue(str string) (bool, string) {
	str = strings.ToLower(str)

	if IsAlpha(str) {
		if len(str) == 1 {
			return Session.TryLetter(str)
		}
		if len(str) > 1 {
			return Session.TryMot(str)
		}
	}
	return false, "Valeur invalide..."
}

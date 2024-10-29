package game

import (
	"Hangmanweb/utils"
	"math/rand"
	"strings"
)

func NouvellePartie(mots []string, difficulte string) utils.GameSession {
	utils.EnJeu = true
	mot := mots[rand.Intn(len(mots))]
	essais := utils.DeterminerEssais(difficulte)

	mob := utils.AssignerMob(difficulte)

	return utils.GameSession{
		MotATrouver:     strings.TrimSpace(mot),
		LettresEssayees: []string{},
		EssaisRestants:  essais,
		MotAffiche:      utils.GenererMotAffiche(mot, []string{}),
		Difficulte:      difficulte,
		Mob:             mob,
	}
}

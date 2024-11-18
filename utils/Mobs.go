package utils

import "math/rand"

func DeterminerEssais(difficulte string) int {
	switch difficulte {
	case "Normal":
		return 8
	case "Difficile":
		return 6
	case "Extreme":
		return 4
	case "Nullos":
		return 27
	case "Entrainement":
		return 12
	default:
		return 0
	}
}

func AssignerMob(difficulte string) string {
	switch difficulte {
	case "Normal":
		mobs := []string{"Zombie", "Squelette", "Piglin", "Cochon"}
		return mobs[rand.Intn(len(mobs))]
	case "Difficile":
		return "Creeper"
	case "Extreme":
		return "Slime"
	case "Nullos":
		return "Enderman"
	case "Entrainement":
		return "Armorstand"
	default:
		return "Inconnu"
	}
}
func PhrasesALeatoire(W bool) string {
	L1 := "Même un cochon aurait fait mieux !"
	L2 := "J'en ai vu des choses mais toi tu es le pire !"
	L3 := "La vérité t'es guez !"
	L4 := "Game Over !"
	L5 := "Toi t'es pas le couteau le plus aiguisé du tiroir !"
	L6 := "Perdu , vous pensez ? Moi je pense pas ."
	L7 := "La lave a eu raison de toi… évite de creuser tout droit la prochaine fois !"
	L8 := "Tu sleep slept slept sur ton keyboard.."
	L9 := "Flop !"
	L10 := "Tu as perdu !"
	W1 := "Je te pensais plus nul..."
	W2 := "Je sais pas quoi dire..."
	W3 := "Pas mal !"
	W4 := "Tu mérite pas ."
	W5 := "J'ai plus d'inpi"
	W6 := "Tu as gagné !"
	W7 := "Gagné !"
	W8 := " Ah la on y bien !"
	W9 := " GG WP !"
	W10 := "EZ!"

	Phrases := []string{L1, L2, L3, L4, L5, L6, L7, L8, L9, L10, W1, W2, W3, W4, W5, W6, W7, W8, W9, W10}
	if W == true {
		return Phrases[rand.Intn(10)+10]
	} else {
		return Phrases[rand.Intn(10)]
	}
}

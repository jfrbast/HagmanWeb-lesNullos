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
		return 200000
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

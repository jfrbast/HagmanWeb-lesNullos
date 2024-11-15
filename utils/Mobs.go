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
func PhrasesALeatoire(W bool) string {
	L1 := "Tu as été explosé par un Creeper… mieux vaut être plus prudent la prochaine fois !"
	L2 := "Les zombies t’ont eu ! Récupère tes forces et retente ta chance !"
	L3 := "Les squelettes étaient trop rapides pour toi… mais tu peux encore devenir plus fort !"
	L4 := "Les mobs de la nuit ont été plus malins… mais pas pour longtemps !"
	L5 := "Ta pioche s'est cassée… il va falloir te réapprovisionner !"
	L6 := "Le Nether t’a vaincu, mais le combat n’est pas terminé !"
	L7 := "La lave a eu raison de toi… évite de creuser tout droit la prochaine fois !"
	L8 := "Le Dragon de l’End t’a surclassé… mais tu reviendras plus puissant !"
	L9 := "Les pillards ont pris l’avantage. Reviens mieux préparé !"
	L10 := "Ton village a été détruit, mais tu pourras tout reconstruire plus grand !"
	W1 := "Bravo ! Tu as vaincu le Dragon de l'End et libéré l’End !"
	W2 := "Tu as récupéré toutes les ressources nécessaires – le monde est à toi !"
	W3 := "Tu es revenu sain et sauf du Nether, victorieux !"
	W4 := "Les pillards n'ont eu aucune chance contre toi – le village est sauvé !"
	W5 := "Les zombies, squelettes et autres monstres ne t’ont pas arrêté ! Quelle aventure !"
	W6 := "Ton équipement en netherite est parfait, rien ne peut t’arrêter maintenant !"
	W7 := "Tu as dompté le Nether et trouvé des trésors rares – un vrai conquérant !"
	W8 := "Ton abri est maintenant indestructible ! Bien joué, bâtisseur !"
	W9 := "Le Wither a été vaincu, et tu as triomphé ! Quelle épopée !"
	W10 := "Félicitations ! Ton royaume s'étend et prospère, tu as tout pour réussir !"

	Phrases := []string{L1, L2, L3, L4, L5, L6, L7, L8, L9, L10, W1, W2, W3, W4, W5, W6, W7, W8, W9, W10}
	if W == true {
		return Phrases[rand.Intn(10)+10]
	} else {
		return Phrases[rand.Intn(10)]
	}
}

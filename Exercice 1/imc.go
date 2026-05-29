package main

import "fmt"

const (
	IMCMaigreur float64 = 18.5
	IMCNormal   float64 = 25.0
	IMCSurpoids float64 = 30.0
	Nom         string  = "Timoté"
)

func main() {
	var poids float64 = 70
	var taille float64 = 1.80

	imc := poids / (taille * taille)

	fmt.Printf("Bonjour %s !\n", Nom)
	fmt.Printf("Poids  : %.2f kg\n", poids)
	fmt.Printf("Taille : %.2f m\n", taille)
	fmt.Printf("IMC    : %.2f\n", imc)

	var categorie string
	if imc < IMCMaigreur {
		categorie = "Maigreur"
	} else if imc < IMCNormal {
		categorie = "Normal"
	} else if imc < IMCSurpoids {
		categorie = "Surpoids"
	} else {
		categorie = "Obésité"
	}

	fmt.Printf("Catégorie : %s\n", categorie)
}

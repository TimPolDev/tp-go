package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func creerOperation(op string) func(float64, float64) float64 {
	switch op {
	case "+":
		return func(a, b float64) float64 { return a + b }
	case "-":
		return func(a, b float64) float64 { return a - b }
	case "*":
		return func(a, b float64) float64 { return a * b }
	case "/":
		return func(a, b float64) float64 { return a / b }
	default:
		return nil
	}
}

func operer(a, b float64, op string) (float64, error) {
	fn := creerOperation(op)
	if fn == nil {
		return 0, fmt.Errorf("opération inconnue : %q (utilisez +, -, *, / ou 'quit')", op)
	}
	if op == "/" && b == 0 {
		return 0, fmt.Errorf("division par zéro interdite")
	}
	return fn(a, b), nil
}

func main() {
	fmt.Println("Calculatrice CLI — format : <a> <b> <op>  (ex: 10 5 +)  |  'quit' pour quitter")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		ligne := strings.TrimSpace(scanner.Text())
		if ligne == "" {
			continue
		}

		champs := strings.Fields(ligne)

		if champs[0] == "quit" {
			fmt.Println("Au revoir !")
			break
		}

		if len(champs) != 3 {
			fmt.Println("Erreur : format attendu '<a> <b> <op>' (3 valeurs séparées par des espaces)")
			continue
		}

		a, err := strconv.ParseFloat(champs[0], 64)
		if err != nil {
			fmt.Printf("Erreur : '%s' n'est pas un nombre valide\n", champs[0])
			continue
		}

		b, err := strconv.ParseFloat(champs[1], 64)
		if err != nil {
			fmt.Printf("Erreur : '%s' n'est pas un nombre valide\n", champs[1])
			continue
		}

		op := champs[2]

		resultat, err := operer(a, b, op)
		if err != nil {
			fmt.Printf("Erreur : %s\n", err)
			continue
		}

		fmt.Printf("%g %s %g = %g\n", a, op, b, resultat)
	}
}

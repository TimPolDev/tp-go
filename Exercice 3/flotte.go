package main

import "fmt"

type Appareil struct {
	Nom  string
	Role string 
}


func installer(role string) {
	switch role {
	case "admin":
		fmt.Println("  + outils admin (monitoring, accès SSH)")
		fallthrough
	case "dev":
		fmt.Println("  + outils dev (git, IDE)")
		fallthrough
	case "user":
		fmt.Println("  + outils user (navigateur, bureautique)")
	}
}

func main() {
	flotte := []Appareil{
		{"PC-001", "user"},
		{"PC-002", "dev"},
	}

	flotte = append(flotte, Appareil{"PC-003", "admin"})

	fmt.Printf("Flotte : %d appareils\n\n", len(flotte))

	for i, a := range flotte {
		fmt.Printf("[%d] %s (rôle: %s)\n", i, a.Nom, a.Role)
		installer(a.Role)
		fmt.Println()
	}

	n := 0
	for {
		n++
		if n == 3 {
			break
		}
	}
	fmt.Printf("Boucle for{} arrêtée à n=%d\n", n)
}

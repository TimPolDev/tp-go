package main

import (
	"errors"
	"fmt"
	"strings"
)

type Produit struct {
	ID        int
	Nom       string
	Marque    string
	Prix      float64
	Stock     int
	Categorie string
	Actif     bool
}

type Catalogue struct {
	Produits []Produit
}

func (c *Catalogue) AjouterProduit(p Produit) error {
	for _, ex := range c.Produits {
		if ex.ID == p.ID {
			return fmt.Errorf("ID %d déjà utilisé", p.ID)
		}
	}
	c.Produits = append(c.Produits, p)
	return nil
}

func (c *Catalogue) TrouverParID(id int) (Produit, error) {
	for _, p := range c.Produits {
		if p.ID == id {
			return p, nil
		}
	}
	return Produit{}, fmt.Errorf("produit ID %d introuvable", id)
}

func (c *Catalogue) TrouverParCategorie(cat string) []Produit {
	var res []Produit
	for _, p := range c.Produits {
		if strings.EqualFold(p.Categorie, cat) {
			res = append(res, p)
		}
	}
	return res
}

func (c *Catalogue) AppliquerReduction(categorie string, pct float64) int {
	n := 0
	for i := range c.Produits {
		if strings.EqualFold(c.Produits[i].Categorie, categorie) {
			c.Produits[i].Prix -= c.Produits[i].Prix * pct / 100
			n++
		}
	}
	return n
}

func (c *Catalogue) Vendre(id, qte int) error {
	for i := range c.Produits {
		if c.Produits[i].ID == id {
			if c.Produits[i].Stock < qte {
				return errors.New("stock insuffisant")
			}
			c.Produits[i].Stock -= qte
			return nil
		}
	}
	return fmt.Errorf("produit ID %d introuvable", id)
}

func (c *Catalogue) Rapport() string {
	var valeur float64
	for _, p := range c.Produits {
		valeur += p.Prix * float64(p.Stock)
	}
	return fmt.Sprintf("%d produit(s) — valeur totale du stock : %.2f €",
		len(c.Produits), valeur)
}

func main() {
	cat := &Catalogue{Produits: []Produit{
		{1, "iPhone15Pro", "Apple", 1229, 12, "Smartphone", true},
		{2, "MacBookAirM3", "Apple", 1499, 7, "Laptop", true},
		{3, "GalaxyS24", "Samsung", 999, 15, "Smartphone", true},
		{4, "ThinkPadX1", "Lenovo", 2199, 4, "Laptop", true},
		{5, "AirPodsPro2", "Apple", 279, 30, "Audio", true},
	}}

	for {
		fmt.Println("\n[1] Ajouter [2] Chercher [3] Soldes [4] Vendre [5] Rapport [0] Quitter")
		var choix int
		fmt.Print("> ")
		fmt.Scan(&choix)

		switch choix {
		case 0:
			return

		case 1:
			var p Produit
			fmt.Print("ID Nom Marque Prix Stock Categorie : ")
			fmt.Scan(&p.ID, &p.Nom, &p.Marque, &p.Prix, &p.Stock, &p.Categorie)
			p.Actif = true
			if err := cat.AjouterProduit(p); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Produit ajouté")
			}

		case 2:
			var id int
			fmt.Print("ID : ")
			fmt.Scan(&id)
			p, err := cat.TrouverParID(id)
			if err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Printf("%+v\n", p)
			}

		case 3:
			var c string
			var pct float64
			fmt.Print("Categorie Pourcentage : ")
			fmt.Scan(&c, &pct)
			fmt.Printf("%d produit(s) mis à jour\n", cat.AppliquerReduction(c, pct))

		case 4:
			var id, qte int
			fmt.Print("ID Quantite : ")
			fmt.Scan(&id, &qte)
			if err := cat.Vendre(id, qte); err != nil {
				fmt.Println("Erreur :", err)
			} else {
				fmt.Println("Vente OK")
			}

		case 5:
			fmt.Println(cat.Rapport())

		default:
			fmt.Println("Choix invalide")
		}
	}
}

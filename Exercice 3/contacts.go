package main

import "fmt"

type Personne struct {
	Prenom, Nom, Email string
	Age                int
}

func (p Personne) NomComplet() string   { return p.Prenom + " " + p.Nom }
func (p Personne) Presentation() string { return fmt.Sprintf("%s, %d ans (%s)", p.NomComplet(), p.Age, p.Email) }

type Adresse struct{ Rue, Ville, CodePostal string }

func (a Adresse) Format() string { return fmt.Sprintf("%s, %s %s", a.Rue, a.CodePostal, a.Ville) }

// Employé : embedding de Personne ET Adresse → accès direct aux champs/méthodes.
type Employe struct {
	Personne
	Adresse
	Poste   string
	Salaire float64
}

func (e Employe) FicheEmploye() string {
	return fmt.Sprintf("[Employé] %s — %s, %.2f € — %s",
		e.Presentation(), e.Poste, e.Salaire, e.Format())
}

// Receveur pointeur : indispensable pour modifier l'employé en place.
func (e *Employe) AugmenterSalaire(pct float64) { e.Salaire += e.Salaire * pct / 100 }

type Etudiant struct {
	Personne
	Promo   string
	Moyenne float64
}

func (e Etudiant) MentionObtenue() string {
	switch {
	case e.Moyenne >= 16:
		return "TB"
	case e.Moyenne >= 14:
		return "B"
	case e.Moyenne >= 12:
		return "AB"
	case e.Moyenne >= 10:
		return "P"
	default:
		return "Insuffisant"
	}
}

func main() {
	emp1 := Employe{Personne{"Alice", "Durand", "alice@corp.fr", 32}, Adresse{"12 rue de la Paix", "Paris", "75002"}, "Ingénieure", 3800}
	emp2 := Employe{Personne{"Bruno", "Martin", "bruno@corp.fr", 45}, Adresse{"5 av. Jaurès", "Lyon", "69007"}, "Responsable RH", 4500}
	etu1 := Etudiant{Personne{"Chloé", "Bernard", "chloe@univ.fr", 21}, "M1 Info", 15.4}
	etu2 := Etudiant{Personne{"David", "Petit", "david@univ.fr", 19}, "L2 Maths", 9.8}

	fmt.Println(emp1.FicheEmploye())
	fmt.Println(emp2.FicheEmploye())

	emp1.AugmenterSalaire(10)
	fmt.Printf(">> Après +10%% : %s gagne %.2f €\n", emp1.NomComplet(), emp1.Salaire)

	fmt.Printf("[Étudiant] %s — %s — moy. %.2f → %s\n", etu1.Presentation(), etu1.Promo, etu1.Moyenne, etu1.MentionObtenue())
	fmt.Printf("[Étudiant] %s — %s — moy. %.2f → %s\n", etu2.Presentation(), etu2.Promo, etu2.Moyenne, etu2.MentionObtenue())
}

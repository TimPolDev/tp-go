package main

import (
	"fmt"
	"math"
	"strings"
)

type Payeur interface {
	Payer(montant float64) (string, error)
}

type CarteCredit struct {
	Numero, Titulaire string
	Solde             float64
}

func (cc *CarteCredit) Payer(montant float64) (string, error) {
	if cc.Solde < montant {
		return "", fmt.Errorf("solde insuffisant : %.2f € disponibles pour %s", cc.Solde, cc.Titulaire)
	}
	if !strings.HasPrefix(cc.Numero, "4") && !strings.HasPrefix(cc.Numero, "5") {
		return "", fmt.Errorf("numéro de carte invalide : %s", cc.Numero)
	}
	cc.Solde -= montant
	suffixe := cc.Numero
	if len(cc.Numero) >= 4 {
		suffixe = cc.Numero[len(cc.Numero)-4:]
	}
	return fmt.Sprintf("Transaction CB #%s confirmée", suffixe), nil
}

type PayPal struct {
	Email string
	Solde float64
}

func (pp *PayPal) Payer(montant float64) (string, error) {
	if pp.Solde < montant {
		return "", fmt.Errorf("solde PayPal insuffisant : %.2f € disponibles pour %s", pp.Solde, pp.Email)
	}
	pp.Solde -= montant
	return fmt.Sprintf("Paiement PayPal de %.2f€ vers %s", montant, pp.Email), nil
}

type Crypto struct {
	Adresse string
	Solde   float64
	Monnaie string
}

func (c *Crypto) Payer(montant float64) (string, error) {
	montantCrypto := math.Round(montant/50000*1000) / 1000
	if c.Solde < montantCrypto {
		return "", fmt.Errorf("solde %s insuffisant : %.3f disponibles, %.3f requis", c.Monnaie, c.Solde, montantCrypto)
	}
	c.Solde -= montantCrypto
	return fmt.Sprintf("Transaction %s : %.3f %s vers %s (%.2f €)", c.Monnaie, montantCrypto, c.Monnaie, c.Adresse, montant), nil
}

func ProcesserPanier(payeur Payeur, articles []float64) {
	var total float64
	for _, prix := range articles {
		total += prix
	}

	switch p := payeur.(type) {
	case *CarteCredit:
		fmt.Printf("→ Mode : Carte bancaire — %s (****%s)\n", p.Titulaire, p.Numero[len(p.Numero)-4:])
	case *PayPal:
		fmt.Printf("→ Mode : PayPal — %s\n", p.Email)
	case *Crypto:
		fmt.Printf("→ Mode : Crypto — %s (%s)\n", p.Monnaie, p.Adresse)
	default:
		fmt.Println("→ Mode : inconnu")
	}

	fmt.Printf("→ Total panier : %.2f €\n", total)

	msg, err := payeur.Payer(total)
	if err != nil {
		fmt.Printf("✗ Échec : %s\n", err)
		return
	}
	fmt.Printf("✓ %s\n", msg)
}

var _ Payeur = (*CarteCredit)(nil)
var _ Payeur = (*PayPal)(nil)
var _ Payeur = (*Crypto)(nil)

func main() {
	articles := []float64{29.99, 15.50, 8.00}

	fmt.Println("═══ Paiement par carte bancaire ═══")
	ProcesserPanier(&CarteCredit{"4532123456789012", "Alice Martin", 100}, articles)

	fmt.Println("\n═══ Paiement PayPal ═══")
	ProcesserPanier(&PayPal{"alice@email.com", 80}, articles)

	fmt.Println("\n═══ Paiement Crypto ═══")
	ProcesserPanier(&Crypto{"1A2B3C4D5E", 0.01, "BTC"}, articles)

	fmt.Println("\n═══ Solde insuffisant (CB) ═══")
	ProcesserPanier(&CarteCredit{"5123456789012345", "Bruno Petit", 10}, articles)
}

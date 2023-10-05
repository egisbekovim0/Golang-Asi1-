package main

import "fmt"

func main() {
	whitePawn, _ := getChessPiece("white_pawn")
	blackKnight, _ := getChessPiece("black_knight")

	printDetails(whitePawn)
	printDetails(blackKnight)
}

func printDetails(p IChessPiece) {
	fmt.Printf("Piece: %s", p.getName())
	fmt.Println()
	fmt.Printf("Strength: %d", p.getStrength())
	fmt.Println()
	fmt.Printf("Defense: %d", p.getDefense())
	fmt.Println()
}

func getChessPiece(pieceType string) (IChessPiece, error) {
	if pieceType == "white_pawn" {
		return newWhitePawn(), nil
	}
	if pieceType == "black_knight" {
		return newBlackKnight(), nil
	}
	return nil, fmt.Errorf("Wrong chess piece")
}

type whitePawn struct {
	ChessPiece
}

func newWhitePawn() IChessPiece {
	return &whitePawn{
		ChessPiece: ChessPiece{
			name:     "White Pawn",
			strength: 1,
			defense:  2,
		},
	}
}

type blackKnight struct {
	ChessPiece
}

func newBlackKnight() IChessPiece {
	return &blackKnight{
		ChessPiece: ChessPiece{
			name:     "Black Knight",
			strength: 3,
			defense:  5,
		},
	}
}

type ChessPiece struct {
	name     string
	strength int
	defense  int
}

func (p *ChessPiece) setName(name string) {
	p.name = name
}

func (p *ChessPiece) getName() string {
	return p.name
}

func (p *ChessPiece) setStrength(strength int) {
	p.strength = strength
}

func (p *ChessPiece) getStrength() int {
	return p.strength
}

func (p *ChessPiece) setDefense(defense int) {
	p.defense = defense
}

func (p *ChessPiece) getDefense() int {
	return p.defense
}

type IChessPiece interface {
	setName(name string)
	setStrength(strength int)
	getName() string
	getStrength() int
	setDefense(defense int)
	getDefense() int
}

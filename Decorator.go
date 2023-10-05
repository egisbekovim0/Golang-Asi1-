package main

import "fmt"

func main() {
	whitePawn := &ChessPieces{name: "White Pawn", strength: 1, defense: 1}

	enhancedWhitePawn := &StrengthDefenseBoostDecorator{piece: whitePawn, strengthBoost: 2, defenseBoost: 3}

	printDetailsFor(enhancedWhitePawn)
}

func printDetailsFor(p ChessElement) {
	fmt.Printf("Piece: %s\nStrength: %d\nDefense: %d\n", p.getName(), p.getStrength(), p.getDefense())
}

type ChessElement interface {
	getName() string
	getStrength() int
	getDefense() int
}

type ChessPieces struct {
	name     string
	strength int
	defense  int
}

func (p *ChessPieces) getName() string  { return p.name }
func (p *ChessPieces) getStrength() int { return p.strength }
func (p *ChessPieces) getDefense() int  { return p.defense }

type ChessDecorator interface {
	ChessElement
	setStrengthBoost(boost int)
	setDefenseBoost(boost int)
}

type StrengthDefenseBoostDecorator struct {
	piece         ChessElement
	strengthBoost int
	defenseBoost  int
}

func (d *StrengthDefenseBoostDecorator) getName() string { return d.piece.getName() }

func (d *StrengthDefenseBoostDecorator) getStrength() int {
	return d.piece.getStrength() + d.strengthBoost
}

func (d *StrengthDefenseBoostDecorator) getDefense() int {
	return d.piece.getDefense() + d.defenseBoost
}

package game

import (
	"github.com/klumhru/4hex/hex"
)

// Player represents a player in the game.
type Player interface {
	AddUnit(u hex.Unit)
	GetName() string
	GetUnits() []hex.Unit
}

// concretePlayer implements the Player interface.
type concretePlayer struct {
	name  string
	units []hex.Unit
}

func (p *concretePlayer) AddUnit(u hex.Unit) {
	p.units = append(p.units, u)
}

func (p *concretePlayer) GetName() string {
	return p.name
}

func (p *concretePlayer) GetUnits() []hex.Unit {
	return p.units
}

// NewPlayer creates a new Player.
func NewPlayer(name string) Player {
	return &concretePlayer{name: name, units: []hex.Unit{}}
}

package game

import (
	"fmt"

	"github.com/klumhru/4hex/hex"
)

// Player represents a player in the game.
type Player interface {
	AddUnit(u hex.Unit)
	GetName() string
	GetUnitCount() int
	GetUnitAt(index int) (hex.Unit, error)
	SetUnitAt(index int, unit hex.Unit) error
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

// GetUnitCount returns the number of units the player has.
func (p *concretePlayer) GetUnitCount() int {
	return len(p.units)
}

// GetUnitAt returns the unit at the specified index.
// Returns an error if the index is out of bounds.
func (p *concretePlayer) GetUnitAt(index int) (hex.Unit, error) {
	if index < 0 || index >= len(p.units) {
		return nil, fmt.Errorf("index out of bounds: %d", index)
	}
	return p.units[index], nil
}

// SetUnitAt sets the unit at the specified index.
// Returns an error if the index is out of bounds.
// Allows setting a nil unit.
func (p *concretePlayer) SetUnitAt(index int, unit hex.Unit) error {
	if index < 0 || index >= len(p.units) {
		return fmt.Errorf("index out of bounds: %d", index)
	}
	p.units[index] = unit
	return nil
}

// NewPlayer creates a new Player.
func NewPlayer(name string) Player {
	return &concretePlayer{name: name, units: []hex.Unit{}}
}

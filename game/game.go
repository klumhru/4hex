package game

import (
	"github.com/klumhru/4hex/hex"
)

// Game represents the main game state.
type Game interface {
	SetMap(m hex.Map)
	AddPlayer(p Player)
}

// concreteGame implements the Game interface.
type concreteGame struct {
	gameMap hex.Map
	players []Player
}

func (g *concreteGame) SetMap(m hex.Map) {
	g.gameMap = m
}

func (g *concreteGame) AddPlayer(p Player) {
	g.players = append(g.players, p)
}

// NewGame creates a new Game.
func NewGame() Game {
	return &concreteGame{players: []Player{}}
}

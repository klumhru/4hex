# 4 HEX

4X Strategy Game Hex System

# 4X Strategy Game Hex System
This is a simple hex system for a 4X strategy game. It is designed to be easy to use and understand, while still being powerful enough to handle complex game mechanics.

It is built using Golang and is designed as a pure Go library, with no external dependencies.

# Features

- Hexagonal grid system
- Unit positioning and Movemement
- Terrain types
- Combat mechanics
- Resource management
- Event system
- Simple AI
- Map generation
- Save and load game state

# Installation

To install the library, simply run:

```bash
go get github.com/klumhru/4hex
```

# Usage

```go
package main
import (
  "fmt"
  "github.com/klumhru/4hex/hex"
  "github.com/klumhru/4hex/game"
  "github.com/klumhru/4hex/player"
  "github.com/klumhru/4hex/unit"
)

func main() {
  // Create a new game
  game := hex.NewGame()

  // Create a new map
  m := hex.NewMap(10, 10)

  // Add the map to the game
  game.SetMap(m)

  // Create a new player
  player := hex.NewPlayer("Player 1")

  // Add the player to the game
  game.AddPlayer(player)

  // Create a new unit
  warrior := unit.NewUnit("Warrior", player)

  // Add the unit to the player's army
  player.AddUnit(warrior)

  // Move the unit to a new position
  warrior.Move(hex.NewPosition(1, 1))

  // Print the unit's position
  fmt.Println(warrior.Position())
}
```

# Tests

To run the tests, simply run:

```bash
go test ./...
```

# Contributing

If you would like to contribute to the project, check out [CONTRIBUTING.md](CONTRIBUTING.md). All contributions are welcome!

# License

This project is licensed under the Apache 2.0 License. See the [LICENSE](LICENSE) file for details.

# Contact

If you have any questions or comments, please feel free to contact me at [klumhru@gmail.com](mailto:klumhru@gmail.com).

# Acknowledgements

This project was inspired by the [Civilization](https://civilization.com/) series of games, as well as other 4X strategy games. The 

hexagonal grid system is based on the work of [Martin O'Leary](https://www.redblobgames.com/grids/hexagons/).

# TODO
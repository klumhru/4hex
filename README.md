# 4 HEX

4X Strategy Game Hex System

# 4X Strategy Game Hex System

This is a simple hex system for a 4X strategy game. It is designed to be easy to use and understand, while still being powerful enough to handle complex game mechanics.

## Features

- Hexagonal grid system
- Unit positioning and movement
- Terrain types
- Combat mechanics
- Resource management
- Event system
- Simple AI
- Map generation
- Save and load game state

## Design

It is built using Golang and is designed as a pure Go library, with no external dependencies. It is intented to be used as an engine for a 4X strategy game. It is not a complete game, but rather a set of tools and libraries that can be used to build a game. The design is modular and extensible, allowing for easy customization and expansion.

## Data Structures

The library uses a hexagonal grid system to represent the game world. Each hexagon is represented as a struct, with properties such as position, terrain type, and resources. The library also includes data structures for units, players, and game events.

The hexagonal grid system is based on the work of [Martin O'Leary](https://www.redblobgames.com/grids/hexagons/), and is designed to be easy to use and understand. The library includes functions for converting between different coordinate systems, as well as functions for calculating distances and angles between hexagons.

## Map Data

The library uses Axial coordinates to represent hexagonal grids, which allows for easy manipulation of hexagon positions. Algorithms for calculating distances, angles, and neighbors are included, making it easy to implement movement and combat mechanics. These are implemented using both axial and cube coordinates, and methods are included for easy conversion between the two systems.

## Map Shapes

The library includes functions for generating different shapes of hexagonal maps, such as rectangular, hexagonal, circular, spiral, up- and down-triangles. New shapes can be added by implementing the `Shape` interface, which defines a method for generating the hexagons in the shape. This allows for easy customization of the map generation process.

## Map Layers

Each hex map is composed of multiple layers of hexagons, allowing for complex terrain and resource management. The library includes functions for generating maps, and custom implementations for applying layers can be added as needed, by passing a function that operates on the hexagons in the layer. This allows for easy customization of the map generation process.

Layers have a root position on the map. This position is arbitrary, allowing for flexibility in how layers are applied. Each layer can be thought of as a separate map, with its own set of hexagons and properties. For example, you can create one up triangle and another down triangle layer, and then add them to the same map. This would create a hour-glass shape on the map. Another alternative is to create two circular layers that intersect at a certain point, creating a venn diagram shape.

Each layer can stores different types of data, such as terrain types, resources, and units. The library includes functions for manipulating these layers, allowing for easy customization of the game world. The library includes functions for gathering data from the map, such as terrain types, resources, and units. This data can be used to create game events, such as resource gathering, and apply game mechanics, such as movement and combat.

Filters can be applied to the map data to gather specific information, such as all units of a certain type, or all hexagons with a certain terrain type. This is implemented using simple query system based on filter predicates. This can be used to calculate movement costs, combat outcomes, and other game mechanics.

## Extending

The library is designed to be easily extensible. You can add new features and functionality by creating new packages and modules. The code is well-documented and easy to understand, making it easy to get started. All types are designed as interfaces, allowing for easy implementation of new functionality.

## Testing and Visualization
The library includes a set of tests to ensure that the code is working correctly. The tests cover all major functionality, including hexagon generation, map generation, and unit movement. The tests are run using the Go testing framework, and can be run using the `go test` command.

The library also includes a simple visualization tool that can be used to visualize the hexagonal grid and the game world. This tool uses [lipgloss](https://github.com/charmbracelet/lipgloss) to create a terminal-based visualization of the game world. This can be used to debug the game world and ensure that the code is working correctly. The visualization tool can be run using the `go run viz/main.go [options]` command, and will display the hexagonal grid in the terminal. Check out the `viz` package and `go run viz/main.go -h` for more information on how to use it.

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
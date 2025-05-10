

<div align="center">

# Terminal Pokedex (Go)

![GitHub top language](https://img.shields.io/github/languages/top/alerone/pokedex-go?color=%2377CDFF)
![GitHub last commit](https://img.shields.io/github/last-commit/alerone/pokedex-go?color=%23bc0bbf)
![GitHub Created At](https://img.shields.io/github/created-at/alerone/pokedex-go?color=%230dba69)
![GitHub repo size](https://img.shields.io/github/repo-size/alerone/pokedex-go?color=%23390385)

<br>

<img src="" alt="golang y http" width="250" height="250"/>

</div>

A simple and interactive terminal-based Pokedex written in Go. This application communicates with the [PokeAPI](https://pokeapi.co) to fetch Pokémon data, locations, and other relevant information. It provides a command-line interface to explore the Pokémon world and build your own Pokedex!

## 🚀 Getting Started

Make sure you have **Go** installed on your system.

To start the Pokedex, run the following command from the root of the project:

```bash
go run .
```


## 📋 Available Commands

| Command   | Description |
|-----------|-------------|
| `map`     | Display the next page of locations |
| `mapb`    | Display the previous page of locations |
| `explore` | Explore a location and discover which Pokémon can appear there |
| `catch`   | Try to catch a Pokémon and save it to your Pokedex |
| `inspect` | View detailed information of a captured Pokémon |
| `pokedex` | List all Pokémon you've caught |
| `help`    | Display this help message |
| `exit`    | Exit the Pokedex |


## 🔍 Inspect Command Output

The `inspect` command displays detailed information only for Pokémon that you've successfully caught.

Example output:

```
Pokedex ID: 19
Name: rattata
Height: 0.30 m
Weight: 3.50 kg
Stats:
  -hp: 30
  -attack: 56
  -defense: 35
  -special-attack: 25
  -special-defense: 35
  -speed: 72
Types:
  -normal

Evolution Chain
---------------
  - rattata -> level up At level: 20 = raticate
```
> [!WARNING]
> When you exit the pokedex all the pokemon captured are flushed away!


## 📦 Dependencies

- [Go](https://golang.org/dl/)  installed (version 1.18+)





package pokedex

import (
    "testing"
)

func TestGetPokemonInfo(t *testing.T) {
    // Test case 1: Valid Pokemon name
    pokemonName := "pikachu"
    pokemon, err := GetPokemonInfo(pokemonName)
    if err != nil {
        t.Errorf("Failed to retrieve Pokemon info for %s: %v", pokemonName, err)
    }
    if pokemon == nil {
        t.Errorf("Received nil Pokemon info for %s", pokemonName)
    }

    // Add more test cases as needed
}

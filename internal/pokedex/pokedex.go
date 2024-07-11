package pokedex

import (
    "encoding/json"
    "fmt"
    "net/http"
)

// Pokemon struct represents a Pokemon object
type Pokemon struct {
    Name      string `json:"name"`
    Abilities []struct {
        Ability struct {
            Name string `json:"name"`
        } `json:"ability"`
    } `json:"abilities"`
    BaseExperience int `json:"base_experience"`
    // Add more fields as needed
}

// Item struct represents an Item object
type Item struct {
    Name        string `json:"name"`
    Description string `json:"effect_entries"`
    // Add more fields as needed
}

// Move struct represents a Move object
type Move struct {
    Name   string `json:"name"`
    Power  int    `json:"power"`
    // Add more fields as needed
}

// ListResponse struct represents the response format for listing items or moves
type ListResponse struct {
    Count   int `json:"count"`
    Results []struct {
        Name string `json:"name"`
    } `json:"results"`
    Next string `json:"next"`
}

// GetPokemonList retrieves a list of Pokémon
func GetPokemonList(limit int) ([]string, error) {
    var allPokemonNames []string
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon?limit=%d", limit)

    for url != "" {
        resp, err := http.Get(url)
        if err != nil {
            return nil, fmt.Errorf("failed to get Pokémon list: %v", err)
        }
        defer resp.Body.Close()

        var pokemonList ListResponse
        if err := json.NewDecoder(resp.Body).Decode(&pokemonList); err != nil {
            return nil, fmt.Errorf("failed to decode Pokémon list: %v", err)
        }

        for _, pokemon := range pokemonList.Results {
            allPokemonNames = append(allPokemonNames, pokemon.Name)
        }

        url = pokemonList.Next
    }

    return allPokemonNames, nil
}

// GetPokemonDetails retrieves detailed information about a Pokémon by name
func GetPokemonDetails(name string) (*Pokemon, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to get Pokémon info: %v", err)
    }
    defer resp.Body.Close()

    var pokemon Pokemon
    if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
        return nil, fmt.Errorf("failed to decode Pokémon info: %v", err)
    }

    return &pokemon, nil
}

// GetItemsList retrieves a list of Items
func GetItemsList(limit int) ([]string, error) {
    var allItemNames []string
    url := fmt.Sprintf("https://pokeapi.co/api/v2/item?limit=%d", limit)

    for url != "" {
        resp, err := http.Get(url)
        if err != nil {
            return nil, fmt.Errorf("failed to get Items list: %v", err)
        }
        defer resp.Body.Close()

        var itemList ListResponse
        if err := json.NewDecoder(resp.Body).Decode(&itemList); err != nil {
            return nil, fmt.Errorf("failed to decode Items list: %v", err)
        }

        for _, item := range itemList.Results {
            allItemNames = append(allItemNames, item.Name)
        }

        url = itemList.Next
    }

    return allItemNames, nil
}

// GetItemDetails retrieves detailed information about an Item by name
func GetItemDetails(name string) (*Item, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/item/%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to get Item info: %v", err)
    }
    defer resp.Body.Close()

    var item Item
    if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
        return nil, fmt.Errorf("failed to decode Item info: %v", err)
    }

    return &item, nil
}

// GetMovesList retrieves a list of Moves
func GetMovesList(limit int) ([]string, error) {
    var allMoveNames []string
    url := fmt.Sprintf("https://pokeapi.co/api/v2/move?limit=%d", limit)

    for url != "" {
        resp, err := http.Get(url)
        if err != nil {
            return nil, fmt.Errorf("failed to get Moves list: %v", err)
        }
        defer resp.Body.Close()

        var moveList ListResponse
        if err := json.NewDecoder(resp.Body).Decode(&moveList); err != nil {
            return nil, fmt.Errorf("failed to decode Moves list: %v", err)
        }

        for _, move := range moveList.Results {
            allMoveNames = append(allMoveNames, move.Name)
        }

        url = moveList.Next
    }

    return allMoveNames, nil
}

// GetMoveDetails retrieves detailed information about a Move by name
func GetMoveDetails(name string) (*Move, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/move/%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to get Move info: %v", err)
    }
    defer resp.Body.Close()

    var move Move
    if err := json.NewDecoder(resp.Body).Decode(&move); err != nil {
        return nil, fmt.Errorf("failed to decode Move info: %v", err)
    }

    return &move, nil
}

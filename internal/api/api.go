package api

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Pokemon struct {
    Name string `json:"name"`
    // Add more fields as needed
}

func GetPokemon(name string) (*Pokemon, error) {
    url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var pokemon Pokemon
    if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
        return nil, err
    }

    return &pokemon, nil
}

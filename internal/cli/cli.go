package cli

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"

    "github.com/naradashen/pokedexcli/internal/pokedex"
)

// ShowMenu displays the main menu and handles user input
func ShowMenu() {
    fmt.Println("Welcome to PokedexCLI!")
    fmt.Println("Select an option:")
    fmt.Println("1. Characters (Pokemon)")
    fmt.Println("2. Items")
    fmt.Println("3. Moves")
    fmt.Print("Enter your choice: ")

    reader := bufio.NewReader(os.Stdin)
    choice, _ := reader.ReadString('\n')
    choice = strings.TrimSpace(choice)

    switch choice {
    case "1":
        fmt.Println("\nFetching Pokemon characters...")
        characters, err := pokedex.GetPokemonList(1000) // Fetching 1000 Pokémon, adjust as needed
        if err != nil {
            fmt.Printf("Error fetching Pokemon characters: %v\n", err)
            return
        }
        fmt.Println("Pokemon characters:")
        for i, character := range characters {
            fmt.Printf("%d. %s\n", i+1, character)
        }
        fmt.Print("Enter the number of the character to view details: ")

        charChoice, _ := reader.ReadString('\n')
        charChoice = strings.TrimSpace(charChoice)
        index, _ := strconv.Atoi(charChoice)
        if index > 0 && index <= len(characters) {
            pokemonName := characters[index-1]
            fmt.Printf("\nFetching details for %s...\n", pokemonName)
            if err := PrintPokemonInfo(pokemonName); err != nil {
                fmt.Printf("Error fetching Pokemon details: %v\n", err)
                return
            }
        } else {
            fmt.Println("Invalid selection.")
        }

    case "2":
        fmt.Println("\nFetching Items...")
        items, err := pokedex.GetItemsList(20) // Fetching 20 Items, adjust as needed
        if err != nil {
            fmt.Printf("Error fetching Items: %v\n", err)
            return
        }
        fmt.Println("Items:")
        for i, item := range items {
            fmt.Printf("%d. %s\n", i+1, item)
        }
        fmt.Print("Enter the number of the item to view details: ")

        itemChoice, _ := reader.ReadString('\n')
        itemChoice = strings.TrimSpace(itemChoice)
        index, _ := strconv.Atoi(itemChoice)
        if index > 0 && index <= len(items) {
            itemName := items[index-1]
            fmt.Printf("\nFetching details for %s...\n", itemName)
            if err := PrintItemInfo(itemName); err != nil {
                fmt.Printf("Error fetching Item details: %v\n", err)
                return
            }
        } else {
            fmt.Println("Invalid selection.")
        }

    case "3":
        fmt.Println("\nFetching Moves...")
        moves, err := pokedex.GetMovesList(20) // Fetching 20 Moves, adjust as needed
        if err != nil {
            fmt.Printf("Error fetching Moves: %v\n", err)
            return
        }
        fmt.Println("Moves:")
        for i, move := range moves {
            fmt.Printf("%d. %s\n", i+1, move)
        }
        fmt.Print("Enter the number of the move to view details: ")

        moveChoice, _ := reader.ReadString('\n')
        moveChoice = strings.TrimSpace(moveChoice)
        index, _ := strconv.Atoi(moveChoice)
        if index > 0 && index <= len(moves) {
            moveName := moves[index-1]
            fmt.Printf("\nFetching details for %s...\n", moveName)
            if err := PrintMoveInfo(moveName); err != nil {
                fmt.Printf("Error fetching Move details: %v\n", err)
                return
            }
        } else {
            fmt.Println("Invalid selection.")
        }

    default:
        fmt.Println("Invalid choice.")
    }
}

// PrintPokemonInfo prints detailed information about a Pokemon
func PrintPokemonInfo(name string) error {
    pokemon, err := pokedex.GetPokemonDetails(name)
    if err != nil {
        return fmt.Errorf("failed to get Pokemon info: %v", err)
    }

    // Print Pokemon details
    fmt.Printf("Name: %s\n", pokemon.Name)
    fmt.Println("Abilities:")
    for _, ability := range pokemon.Abilities {
        fmt.Printf("- %s\n", ability.Ability.Name)
    }
    fmt.Printf("Base Experience: %d\n", pokemon.BaseExperience)
    // Add more fields as needed

    return nil
}

// PrintItemInfo prints detailed information about an Item
func PrintItemInfo(name string) error {
    item, err := pokedex.GetItemDetails(name)
    if err != nil {
        return fmt.Errorf("failed to get Item info: %v", err)
    }

    // Print Item details
    fmt.Printf("Name: %s\n", item.Name)
    fmt.Printf("Description: %s\n", item.Description)
    // Add more fields as needed

    return nil
}

// PrintMoveInfo prints detailed information about a Move
func PrintMoveInfo(name string) error {
    move, err := pokedex.GetMoveDetails(name)
    if err != nil {
        return fmt.Errorf("failed to get Move info: %v", err)
    }

    // Print Move details
    fmt.Printf("Name: %s\n", move.Name)
    fmt.Printf("Power: %d\n", move.Power)
    // Add more fields as needed

    return nil
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
	"repl-project/internal/pokeapi"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("no pokemon name provided")
	}

	pokemonName := args[0]

	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)

	fmt.Println(pokemon.BaseExperience, threshold, randNum)
	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
		// fmt.Printf("%s escaped", pokemon.Name)
	}
	if cfg.caughtPokemon == nil {
		cfg.caughtPokemon = make(map[string]pokeapi.Pokemon) // Replace 'PokemonType' with the actual type of 'pokemon'
	}
	cfg.caughtPokemon[pokemon.Name] = pokemon
	fmt.Printf("%s was caught!\n", pokemon.Name)

	return nil
}

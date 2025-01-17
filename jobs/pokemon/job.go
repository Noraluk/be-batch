package pokemon

import (
	"be-batch/jobs"
	"be-batch/jobs/pokemon/entities"
	"be-batch/jobs/pokemon/models"
	"be-batch/pkg/base"
	"be-batch/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type pokemonJob struct {
	id         string
	repository base.BaseRepository[any]
	logger     logger.Logger
}

func NewPokemonJob(
	repository base.BaseRepository[any],
) jobs.Job {
	return &pokemonJob{
		id:         "pokemon",
		repository: repository,
		logger:     logger.WithPrefix("job/pokemon"),
	}
}

func (j pokemonJob) GetID() string {
	return j.id
}

func (j pokemonJob) Run() error {
	errCh := make(chan error, 1)

	j.getPokemons(errCh)

	close(errCh)

	if e := <-errCh; e != nil {
		return e
	}

	return nil
}

func (j pokemonJob) getPokemons(errCh chan error) {
	resp, err := http.Get("https://pokeapi.co/api/v2/pokemon?offset=0&limit=1302")
	if err != nil {
		errCh <- err
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errCh <- err
		return
	}

	var pokemonsResp models.PokemonsResponse
	err = json.Unmarshal(body, &pokemonsResp)
	if err != nil {
		errCh <- err
		return
	}

	pokemons := []entities.Pokemon{}
	pokemonAbilities := []entities.PokemonAbility{}
	pokemonStats := []entities.PokemonStat{}
	pokemonTypes := []entities.PokemonType{}
	pokemonWeaknesses := []entities.PokemonWeakness{}

	pokemonDetailCh := make(chan models.PokemonDetail)
	stopChannel := make(chan struct{})
	channelWg := &sync.WaitGroup{}
	channelWg.Add(1)
	go func() {
		defer channelWg.Done()
		for {
			select {
			case pd := <-pokemonDetailCh:
				pokemons = append(pokemons, pd.Pokemon)
				pokemonAbilities = append(pokemonAbilities, pd.PokemonAbilities...)
				pokemonStats = append(pokemonStats, pd.PokemonStats...)
				pokemonTypes = append(pokemonTypes, pd.PokemonTypes...)
				pokemonWeaknesses = append(pokemonWeaknesses, pd.PokemonWeaknesses...)
			case <-stopChannel:
				log.Println("stop channel")
				return
			}
		}
	}()

	wg2 := &sync.WaitGroup{}
	for i, p := range pokemonsResp.Results {
		wg2.Add(1)
		go func(index int, name string) {
			defer wg2.Done()
			pd := j.getPokemonDetail(index+1, name)
			pokemonDetailCh <- pd
		}(i, p.Name)
	}

	wg2.Wait()
	close(stopChannel)
	channelWg.Wait()

	err = j.repository.Where("1 = 1").Delete(&entities.Pokemon{}).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Where("1 = 1").Delete(&entities.PokemonType{}).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Where("1 = 1").Delete(&entities.PokemonAbility{}).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Where("1 = 1").Delete(&entities.PokemonStat{}).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Where("1 = 1").Delete(&entities.PokemonWeakness{}).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Create(&pokemons).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Create(&pokemonTypes).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Create(&pokemonAbilities).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Create(&pokemonStats).Error()
	if err != nil {
		errCh <- err
		return
	}

	err = j.repository.Create(&pokemonWeaknesses).Error()
	if err != nil {
		errCh <- err
		return
	}
}

func (j pokemonJob) getPokemonDetail(index int, name string) models.PokemonDetail {
	j.logger.Wrap("searching pokemon : %s", name).Info()
	pokemon, err := j.getPokemon(name)
	if err != nil {
		j.logger.Wrap("get pokemon failed, error: %v", err).Error()
		return models.PokemonDetail{}
	}

	wg := &sync.WaitGroup{}
	pokemonTypes := make([]entities.PokemonType, 0)
	pokemonWeaknesses := make([]entities.PokemonWeakness, 0)
	for _, v := range pokemon.Types {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			pokemonType, _ := j.getPokemonType(url)

			pokemonTypes = append(pokemonTypes, entities.PokemonType{
				Name:      pokemonType.Name,
				PokemonID: pokemon.ID,
			})

			for _, v2 := range pokemonType.DamageRelations.DoubleDamageFrom {
				pokemonWeaknesses = append(pokemonWeaknesses, entities.PokemonWeakness{
					Name:          v2.Name,
					PokemonID:     pokemon.ID,
					PokemonTypeID: pokemonType.ID,
				})
			}
		}(v.Type.URL)
	}

	var minimumLevel int = 1
	var evolvedPokemonID *int = nil
	var basePokemonID int

	wg.Add(1)
	go func() {
		defer wg.Done()

		pokemonSpecies, _ := j.getPokemonSpecies(pokemon.Species.URL)
		if len(pokemonSpecies.EvolutionChain.URL) > 0 {
			pokemonEvolutionChainResp, _ := j.getPokemonEvolutionChain(pokemonSpecies.EvolutionChain.URL)
			basePokemon, _ := j.getPokemon(pokemonEvolutionChainResp.Chain.Species.Name)
			basePokemonID = basePokemon.ID
			if pokemon.Name == pokemonEvolutionChainResp.Chain.Species.Name && len(pokemonEvolutionChainResp.Chain.EvolvesTo) > 0 {
				secondPokemon, _ := j.getPokemon(pokemonEvolutionChainResp.Chain.EvolvesTo[0].Species.Name)
				evolvedPokemonID = &secondPokemon.ID
			} else if len(pokemonEvolutionChainResp.Chain.EvolvesTo) > 0 && pokemon.Name == pokemonEvolutionChainResp.Chain.EvolvesTo[0].Species.Name {
				if len(pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolutionDetails) > 0 {
					minimumLevel = pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolutionDetails[0].MinLevel
				}

				if len(pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo) > 0 {
					thirdPokemon, _ := j.getPokemon(pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo[0].Species.Name)
					evolvedPokemonID = &thirdPokemon.ID
				}

			} else if len(pokemonEvolutionChainResp.Chain.EvolvesTo) > 0 && len(pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo) > 0 && pokemon.Name == pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo[0].Species.Name && len(pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo[0].EvolutionDetails) > 0 {
				minimumLevel = pokemonEvolutionChainResp.Chain.EvolvesTo[0].EvolvesTo[0].EvolutionDetails[0].MinLevel
			}
		}
	}()

	wg.Wait()

	pokemonStats := make([]entities.PokemonStat, 0)
	for _, v := range pokemon.Stats {
		pokemonStats = append(pokemonStats, entities.PokemonStat{
			Name:      v.Stat.Name,
			BaseStat:  v.BaseStat,
			PokemonID: pokemon.ID,
		})
	}

	pokemonAbilities := make([]entities.PokemonAbility, 0)
	for _, v := range pokemon.Abilities {
		pokemonAbilities = append(pokemonAbilities, entities.PokemonAbility{
			Name:      v.Ability.Name,
			PokemonID: pokemon.ID,
		})
	}

	return models.PokemonDetail{
		Pokemon: entities.Pokemon{
			ID:                                   index,
			PokemonID:                            pokemon.ID,
			Name:                                 pokemon.Name,
			SpriteFrontDefaultShowdownURL:        pokemon.Sprites.Other.Showdown.FrontDefault,
			SpriteFrontDefaultOfficialArtworkURL: pokemon.Sprites.Other.OfficialArtwork.FrontDefault,
			Height:                               pokemon.Height,
			Weight:                               pokemon.Weight,
			BaseExperience:                       pokemon.BaseExperience,
			MinimumLevel:                         minimumLevel,
			EvolvedPokemonID:                     evolvedPokemonID,
			BasePokemonID:                        basePokemonID,
		},
		PokemonTypes:      pokemonTypes,
		PokemonStats:      pokemonStats,
		PokemonAbilities:  pokemonAbilities,
		PokemonWeaknesses: pokemonWeaknesses,
	}
}

func (j pokemonJob) getPokemon(name string) (models.PokemonResponse, error) {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", name))
	if err != nil {
		j.logger.Wrap("call pokemon %s failed, error: %v", name, err).Error()
		return models.PokemonResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read pokemon response body, error: %v", err).Error()
		return models.PokemonResponse{}, err
	}

	var pokemon models.PokemonResponse
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		j.logger.Wrap("unmarshal pokemon %s response %s, error: %v", name, string(body), err).Error()
		return models.PokemonResponse{}, err
	}

	return pokemon, nil
}

func (j pokemonJob) getPokemonType(url string) (models.PokemonTypeResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		j.logger.Wrap("call pokemon failed, error: %v", err).Error()
		return models.PokemonTypeResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read pokemon response body, error: %v", err).Error()
		return models.PokemonTypeResponse{}, err
	}

	var pokemonType models.PokemonTypeResponse
	err = json.Unmarshal(body, &pokemonType)
	if err != nil {
		j.logger.Wrap("unmarshal pokemon type response, error: %v", err).Error()
		return models.PokemonTypeResponse{}, err
	}

	return pokemonType, nil
}

func (j pokemonJob) getPokemonSpecies(url string) (models.PokemonSpeciesResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		j.logger.Wrap("call pokemon failed, error: %v", err).Error()
		return models.PokemonSpeciesResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read pokemon response body, error: %v", err).Error()
		return models.PokemonSpeciesResponse{}, err
	}

	var pokemonSpecies models.PokemonSpeciesResponse
	err = json.Unmarshal(body, &pokemonSpecies)
	if err != nil {
		j.logger.Wrap("unmarshal pokemon species response, error: %v", err).Error()
		return models.PokemonSpeciesResponse{}, err
	}

	return pokemonSpecies, nil
}

func (j pokemonJob) getPokemonEvolutionChain(url string) (models.PokemonEvolutionChainResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		j.logger.Wrap("call pokemon failed, error: %v", err).Error()
		return models.PokemonEvolutionChainResponse{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read pokemon response body, error: %v", err).Error()
		return models.PokemonEvolutionChainResponse{}, err
	}

	var pokemonEvolutionChainResponse models.PokemonEvolutionChainResponse
	err = json.Unmarshal(body, &pokemonEvolutionChainResponse)
	if err != nil {
		j.logger.Wrap("unmarshal pokemon evolution chain response, error: %v", err).Error()
		return models.PokemonEvolutionChainResponse{}, err
	}

	return pokemonEvolutionChainResponse, nil
}

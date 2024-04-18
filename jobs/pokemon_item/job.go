package pokemonitem

import (
	"be-batch/jobs"
	"be-batch/jobs/pokemon_item/entities"
	"be-batch/jobs/pokemon_item/models"
	"be-batch/pkg/base"
	"be-batch/pkg/logger"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type pokemonItemJob struct {
	id         string
	repository base.BaseRepository[any]
	logger     logger.Logger
}

func NewPokemonItemJob(
	repository base.BaseRepository[any],
) jobs.Job {
	return &pokemonItemJob{
		id:         "pokemon_item",
		repository: repository,
		logger:     logger.WithPrefix("job/pokemon_item"),
	}
}

func (j pokemonItemJob) GetID() string {
	return j.id
}

func (j pokemonItemJob) Run() error {
	err := j.repository.Where("1 = 1").Delete(&entities.PokemonItem{}).Error()
	if err != nil {
		j.logger.Wrap("delete all items failed, error: %v", err).Error()
		return err
	}

	for i := 0; i <= 2110; i += 1000 {
		err := j.getItems(i)
		if err != nil {
			return err
		}
	}

	return nil
}

func (j pokemonItemJob) getItems(i int) error {
	resp, err := http.Get(fmt.Sprintf("https://pokeapi.co/api/v2/item?offset=%d&limit=1000", i))
	if err != nil {
		j.logger.Wrap("call items failed, error: %v", err).Error()
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read items response body, error: %v", err).Error()
		return err
	}

	var itemsResp models.PokemonItemsResponse
	err = json.Unmarshal(body, &itemsResp)
	if err != nil {
		j.logger.Wrap("unmarshal items response %s, error: %v", string(body), err).Error()
		return err
	}

	errCh := make(chan error, 1)
	items := make([]entities.PokemonItem, len(itemsResp.Results))
	wg := &sync.WaitGroup{}
	for k, v := range itemsResp.Results {
		wg.Add(1)
		go func(index int, url string, k2 int) {
			defer wg.Done()
			item, err := j.getItem(url, index)
			if err != nil {
				errCh <- err
				return
			}
			items[k2] = item
		}(i+k+1, v.URL, k)
	}

	wg.Wait()
	close(errCh)

	if e := <-errCh; e != nil {
		return e
	}

	j.logger.Wrap("items : %d", len(items)).Info()
	err = j.repository.Create(&items).Error()
	if err != nil {
		j.logger.Wrap("create items failed, error: %v", err).Error()
		return err
	}

	return nil
}

func (j pokemonItemJob) getItem(url string, index int) (entities.PokemonItem, error) {
	resp, err := http.Get(url)
	if err != nil {
		j.logger.Wrap("get item failed, error: %v", err).Error()
		return entities.PokemonItem{}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		j.logger.Wrap("read item response body, error: %v", err).Error()
		return entities.PokemonItem{}, err
	}

	var itemResp models.PokemonItemResponse
	err = json.Unmarshal(body, &itemResp)
	if err != nil {
		j.logger.Wrap("unmarshal item %s response %s, error: %v", url, string(body), err).Error()
		return entities.PokemonItem{}, err
	}

	return entities.PokemonItem{
		ID:        index,
		ItemID:    itemResp.ID,
		Name:      itemResp.Name,
		Cost:      itemResp.Cost,
		SpriteURL: itemResp.Sprites.Default,
	}, nil
}

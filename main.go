package main

import (
	"be-batch/jobs"
	pokemonlist "be-batch/jobs/pokemon"
	pokemonListEtt "be-batch/jobs/pokemon/entities"
	"be-batch/pkg/base"
	"be-batch/pkg/config"
	"be-batch/pkg/database"
	"flag"
)

func main() {
	jobName := flag.String("job_name", "", "")
	if jobName == nil {
		panic("job name is required")
	}

	flag.Parse()

	err := config.Init()
	if err != nil {
		panic(err)
	}

	err = database.Init()
	if err != nil {
		panic(err)
	}

	db := database.GetDatabase()
	db.AutoMigrate(&pokemonListEtt.Pokemon{})
	db.AutoMigrate(&pokemonListEtt.PokemonType{})
	db.AutoMigrate(&pokemonListEtt.PokemonAbility{})
	db.AutoMigrate(&pokemonListEtt.PokemonWeakness{})
	db.AutoMigrate(&pokemonListEtt.PokemonStat{})

	pokemonJob := pokemonlist.NewPokemonJob(
		base.NewBaseRepository[any](),
		base.NewBaseRepository[[]pokemonListEtt.Pokemon](),
		base.NewBaseRepository[[]pokemonListEtt.PokemonType](),
		base.NewBaseRepository[[]pokemonListEtt.PokemonAbility](),
		base.NewBaseRepository[[]pokemonListEtt.PokemonStat](),
		base.NewBaseRepository[[]pokemonListEtt.PokemonWeakness](),
	)

	jobs := []jobs.Job{
		pokemonJob,
	}

	jobMap := make(map[string]func() error, 0)
	for _, job := range jobs {
		jobMap[job.GetID()] = job.Run
	}

	err = jobMap[*jobName]()
	if err != nil {
		panic(err)
	}
}

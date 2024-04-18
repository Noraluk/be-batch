package main

import (
	"be-batch/jobs"
	"be-batch/jobs/pokemon"
	pokemonitem "be-batch/jobs/pokemon_item"
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

	repository := base.NewBaseRepository[any]()
	pokemonJob := pokemon.NewPokemonJob(repository)
	pokemonItemJob := pokemonitem.NewPokemonItemJob(repository)

	jobs := []jobs.Job{
		pokemonJob,
		pokemonItemJob,
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

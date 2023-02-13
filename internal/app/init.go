package app

import (
  "github.com/FadhilAF/s-tech-pplbo/config"

  "github.com/FadhilAF/s-tech-pplbo/internal/repository"
)

type App struct {
	Config   config.Config
	delivery deliveries
	usecase  usecases
	store    repository.Store
}

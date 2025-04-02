package app

import (
	"context"
	"inventory_service/config"
)

type Application struct {
}

func New(ctx context.Context, config *config.Config) (*Application, error) {
	return &Application{}, nil
}

func (app *Application) Run() error {
	return nil
}

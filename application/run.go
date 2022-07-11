package application

import "sut-notification-go/config"

func (app *Application) Run(cfg *config.Config) error {
	return grpcRun(cfg)(app)
}

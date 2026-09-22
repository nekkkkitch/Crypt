package app

import (
	"context"
)

// pусть app при инициализации принимает на вход 4 объекта подходящих под данный интерфейс, раскидывая их в мапу, пусть будет map[string]Cypher
type CryptAnalysator interface {
}

// App struct
type App struct {
	a   CryptAnalysator
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(a CryptAnalysator) *App {
	app := &App{}
	app.a = a
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

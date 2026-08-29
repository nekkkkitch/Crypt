package app

import (
	"context"
	"log/slog"
)

// pусть app при инициализации принимает на вход 4 объекта подходящих под данный интерфейс, раскидывая их в мапу, пусть будет map[string]Cypher
type Cypher interface {
	Cypher(input string) string
	Decypher(input string) string
	ChangeParams(params []interface{})
}

// App struct
type App struct {
	cyphers       map[string]Cypher
	currentCypher Cypher
	ctx           context.Context
}

// NewApp creates a new App application struct
func NewApp(a, s, p, c, g, v, pl Cypher) *App {
	app := &App{}
	app.cyphers = map[string]Cypher{
		"atbash":    a,
		"scytale":   s,
		"polybius":  p,
		"caesar":    c,
		"gronsfeld": g,
		"vigener":   v,
		"playfair":  pl,
	}
	app.currentCypher = a
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Cypher(input string) string {
	return a.currentCypher.Cypher(input)
}

func (a *App) Decypher(input string) string {
	return a.currentCypher.Decypher(input)
}

func (a *App) ChangeParams(params []interface{}) {
	a.currentCypher.ChangeParams(params)
}

func (a *App) ChangeCypher(choice string) {
	slog.Info("Got cypher to change", "cyhper", choice)
	if cypher, ok := a.cyphers[choice]; ok {
		a.currentCypher = cypher
	}
}

package app

import (
	crypanalysis "13/backend/services/cryptanalysis"
	"context"
	"encoding/json"
	"log/slog"
)

type CryptAnalysator interface {
	Analyse() crypanalysis.Data
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

func (a *App) Analyse() []byte {
	data := a.a.Analyse()
	toRet, err := json.Marshal(data)
	if err != nil {
		slog.Error("json err", "err", err)
	}
	return toRet
}

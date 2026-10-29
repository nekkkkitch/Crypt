package app

import (
	"13/backend/services/cryptanalysis"
	"context"
	"log/slog"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type CryptAnalysator interface {
	Analyse() cryptanalysis.Data
	SetText(text string, isRus bool)
	SetFile(filePath string, isRus bool)
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

func (a *App) Analyse() cryptanalysis.Data {
	data := a.a.Analyse()
	return data
}

func (a *App) SetText(text string, language string) {
	isRus := false
	if language == "russian" {
		isRus = true
	}
	a.a.SetText(text, isRus)
}

func (a *App) SetFile(language string) {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{DefaultDirectory: ".", Title: "Select file"})
	if err != nil {
		slog.Error("SetFile", "err", err)
		return
	}
	slog.Info("setfile", "file", filePath)
	isRus := false
	if language == "russian" {
		isRus = true
	}
	a.a.SetFile(filePath, isRus)
}

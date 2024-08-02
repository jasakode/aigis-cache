package aigiscache

import (
	"os"
)

type AigisCache interface{}

type Config struct {
	file_path string
}

type App struct {
	config Config
}

func New(config ...Config) AigisCache {
	app := App{}
	if len(config) > 0 {
		app.config = config[0]
	}
	app.config = app.init()

	return &app
}

func (app *App) init() Config {
	if app.config.file_path == "" {
		app.config.file_path = "aigiscache.bin"
	}
	_, err := os.Stat(app.config.file_path)
	if os.IsNotExist(err) {
		if err := os.WriteFile(app.config.file_path, []byte{63, 65, 73, 71, 73, 83, 32, 67, 65, 67, 72, 69}, os.ModeDevice); err != nil {
			panic(err.Error())
		}
	}
	return app.config
}

package tests

import (
	"github.com/stretchr/testify/require"
	"log"
	"path/filepath"
	"receipt-loader/internal/app"
	"receipt-loader/internal/db"
	"receipt-loader/internal/utils"
	"testing"
)

func AppSetup(t *testing.T) *app.App {
	t.Log("Starting AppSetup")
	app := app.NewApp()

	configPath := filepath.Join(utils.GetProjectRoot(), ".env.test")
	err := app.Config.Load(configPath)
	if err != nil {
		panic(err)
	}
	require.Nil(t, err)

	err = app.Setup()
	if err != nil {
		panic(err)
	}
	require.Nil(t, err)

	db.MigrateUp(app.DB)
	t.Log("App setup complete")
	return &app
}

func AppTeardown(app *app.App) {
	log.Println("Starting AppTeardown")
	db.MigrateDown(app.DB)
	app.Teardown()
	log.Println("App teardown complete")
}

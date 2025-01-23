package tests

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"receipt-loader/internal/app"
	"receipt-loader/internal/db"
	"receipt-loader/internal/utils"
	"testing"
)

func AppSetup(t *testing.T) *app.App {
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

	return &app
}

func AppTeardown(app *app.App) {
	db.MigrateDown(app.DB)
	app.Teardown()
}

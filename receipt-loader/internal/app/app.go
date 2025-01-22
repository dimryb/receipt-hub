package app

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"os"
	"receipt-loader/internal/handlers"
)

type App struct {
	Config Config
	Router *mux.Router
	DB     *gorm.DB
}

type Config struct {
	DatabaseUrl string
}

func (config *Config) Load(filename string) error {
	err := godotenv.Load(filename)
	if err != nil {
		return errors.New("Error loading " + filename)
	}

	config.DatabaseUrl = os.Getenv("DATABASE_URL")
	if config.DatabaseUrl == "" {
		return errors.New("missing DATABASE_URL")
	}

	return nil
}

func NewApp() App {
	return App{}
}

func (app *App) Setup() error {
	db, err := gorm.Open(postgres.Open(app.Config.DatabaseUrl), &gorm.Config{})
	if err != nil {
		return err
	}

	r := mux.NewRouter()
	r.HandleFunc(`/receipt/{id:\d+}`, handlers.GetReceiptByID(db)).Methods("GET")
	r.HandleFunc(`/receipt`, handlers.AddReceipt(db)).Methods("POST")

	app.Router = r
	app.DB = db

	return nil
}

func (app *App) Teardown() error {
	sqlDB, err := app.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func (app *App) Run() error {
	return http.ListenAndServe(":8080", app.Router)
}

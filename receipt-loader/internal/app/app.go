package app

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"receipt-loader/internal/db"
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
	//db, err := gorm.Open(postgres.Open(app.Config.DatabaseUrl), &gorm.Config{})
	dataBase, err := db.Connect(app.Config.DatabaseUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	r := mux.NewRouter()
	r.Use(Middleware)
	r.HandleFunc(`/receipt/{id:\d+}`, handlers.GetReceiptByID(dataBase)).Methods("GET")
	r.HandleFunc(`/receipt`, handlers.AddReceipt(dataBase)).Methods("POST")
	r.HandleFunc(`/ping`, handlers.Ping)

	app.Router = r
	app.DB = dataBase

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

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w.Header().Get("Content-Type") == "" {
			w.Header().Set("Content-Type", "application/json")
		}
		next.ServeHTTP(w, r)
	})
}

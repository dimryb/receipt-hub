package db

import (
	"github.com/pressly/goose"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Connect устанавливает соединение с базой данных и возвращает экземпляр *gorm.DB
func Connect(dsn string) *gorm.DB {
	// Подключение к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = runSqlMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run SQL migration: %v", err)
	}

	log.Println("Database connection established successfully")
	return db
}

func runSqlMigrations(db *gorm.DB) error {
	// Получаем объект sql.DB для работы с миграциями через goose
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// Устанавливаем диалект для goose
	if err := goose.SetDialect("postgres"); err != nil {
		panic(err)
	}

	// Выполняем миграции из папки migrations
	migrationsDir := "internal/migrations" // Путь к папке с миграциями
	if err := goose.Up(sqlDB, migrationsDir); err != nil {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
	return nil
}

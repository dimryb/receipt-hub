package db

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"receipt-loader/internal/utils"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	configPath := filepath.Join(utils.GetProjectRoot(), ".env.test")
	if err := godotenv.Load(configPath); err != nil {
		t.Fatalf("Failed to load environment variables from .env.test: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		t.Fatalf("DATABASE_URL is not set in .env.test")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	require.NoError(t, err, "Failed to connect to database: %v", err)

	err = db.Exec("SELECT 1").Error
	require.NoError(t, err, "Database is not ready: %v", err)

	sqlDB, err := db.DB()
	require.NoError(t, err, "Failed to get database instance")
	defer sqlDB.Close()
}

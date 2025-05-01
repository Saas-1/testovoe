package internal

import (
	"awesomeProject19/logger"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	pgmigrate "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

var (
	db   *gorm.DB
	once sync.Once
)

func init() {
	requiredEnvVars := []string{"DB_HOST", "DB_USER", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			logger.Log.Fatalf("Ошибка: переменная окружения %s не установлена", envVar)
		}
	}
	logger.Log.Infof("Переменные окружения успешно загружены")
}

func ConnectDB() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"))

		logger.Log.Debugf("Попытка подключения к БД с DSN: %s", dsn)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			logger.Log.Errorf("Ошибка подключения к БД: %v", err)
			return
		}
		logger.Log.Infof("Успешное подключение к БД")

		sqlDB, err := db.DB()
		if err != nil {
			logger.Log.Errorf("Ошибка получения sql.DB: %v", err)
			return
		}

		if err := runMigrations(sqlDB); err != nil {
			logger.Log.Errorf("Ошибка миграций: %v", err)
		}
	})

	return db, err
}

func runMigrations(sqlDB *sql.DB) error {
	logger.Log.Debug("Запуск миграций...")
	driver, err := pgmigrate.WithInstance(sqlDB, &pgmigrate.Config{})
	if err != nil {
		return fmt.Errorf("не удалось создать драйвер миграций: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"person_db",
		driver,
	)
	if err != nil {
		return fmt.Errorf("не удалось инициализировать мигратор: %w", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("не удалось применить миграции: %w", err)
	}

	logger.Log.Infof("SQL-миграции успешно применены")
	return nil
}

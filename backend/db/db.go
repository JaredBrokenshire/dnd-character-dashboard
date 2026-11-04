package db

import (
	"dnd-character-dashboard/pkg/file_storage"
	"dnd-character-dashboard/pkg/go-migrations/store"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"strconv"

	"github.com/ShkrutDenis/go-migrations/engine/migrator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"

	"dnd-character-dashboard/db/migrations/process"
	"dnd-character-dashboard/db/seeders"
)

// Init opens a connection to MySQL, runs all pending migrations, then seeds
// baseline data.
func Init(fileStore file_storage.Store) (*gorm.DB, error) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		os.Getenv("DB_PASSWORD"),
		host,
		port,
		os.Getenv("DB_NAME"),
	)

	log.Printf("Connecting to %v on port %v with username %v", host, port, user)

	db, err := gorm.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

	// Reuse the already-open connection for migrations so no second DB
	// connection or .env file is needed.
	sqlxDB := sqlx.NewDb(db.DB(), driver)
	runMigrations(sqlxDB, driver, process.List())

	seeder := seeders.NewSeeder(db, fileStore)
	seeder.Seed()

	return db, nil
}

// runMigrations creates the migrations table if needed and runs any pending
// migrations using the existing sqlx connection.
func runMigrations(db *sqlx.DB, driver string, migrations []store.Migratable) {
	m := migrator.NewMigrator(db, driver)
	firstRun := m.CreateMigrationsTable()

	for _, migration := range migrations {
		if !firstRun && m.MigrationExist(migration.GetName()) {
			continue
		}
		log.Println("Migrating", migration.GetName())
		migration.Up(db)
		m.AddMigrationRaw(migration.GetName())
		log.Println("Migrated", migration.GetName())
	}
}

// Paginate takes an echo context and returns the page, page size and a paginate
// method from the query parameters
func Paginate(c echo.Context) (int, int, func(db *gorm.DB) *gorm.DB) {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 25
	}

	return page, pageSize, func(db *gorm.DB) *gorm.DB {
		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

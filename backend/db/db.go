package db

import (
	"dnd/db/migrations/process"
	"dnd/db/seeders"
	"dnd/pkg/file_storage"
	"dnd/pkg/go-migrations/store"
	"dnd/pkg/logging"
	"fmt"
	"github.com/ShkrutDenis/go-migrations/engine/migrator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
	"os"
)

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

	db, err := gorm.Open(driver, dsn)
	if err != nil {
		panic(err.Error())
	}

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
		logging.Logger.LogInfo().Msg(fmt.Sprintf("Migrating %v", migration.GetName()))
		migration.Up(db)
		m.AddMigrationRaw(migration.GetName())
		logging.Logger.LogInfo().Msg(fmt.Sprintf("Migrated %v", migration.GetName()))
	}
}

type Scopes []func(db *gorm.DB) *gorm.DB

func Paginate(page, pageSize int) (int, func(db *gorm.DB) *gorm.DB) {
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 25
	}

	return pageSize, func(db *gorm.DB) *gorm.DB {
		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

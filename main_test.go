package gormgrate_test

import (
	"log"
	"os"
	"testing"

	"github.com/CrowderSoup/gormgrate"
	"github.com/CrowderSoup/gormgrate/examples"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}

	// Run Auto Migrations
	db.AutoMigrate(examples.User{})

	return db
}

func cleanupDB(db *gorm.DB) {
	db.Close()

	err := os.Remove("test.db")
	if err != nil {
		log.Fatal(err)
	}
}

func getFiles() []gormgrate.MigrationFile {
	return []gormgrate.MigrationFile{
		examples.NewExampleMigration(),
	}
}

func TestNewMigrator(t *testing.T) {
	db := initDB()
	defer cleanupDB(db)

	migrator, err := gormgrate.NewMigrator(db, true, getFiles())
	assert.Nil(t, err)
	assert.NotNil(t, migrator)
}

func TestRunMigrations(t *testing.T) {
	db := initDB()
	defer cleanupDB(db)

	migrator, err := gormgrate.NewMigrator(db, true, getFiles())
	assert.Nil(t, err)
	assert.NotNil(t, migrator)

	err = migrator.RunMigrations()
	assert.Nil(t, err)

	var count int
	db.Where(&gormgrate.Migration{Successful: true}).Find(&gormgrate.Migration{}).Count(&count)
	assert.Equal(t, len(getFiles()), count)

	migratorDown, err := gormgrate.NewMigrator(db, false, getFiles())
	assert.Nil(t, err)
	assert.NotNil(t, migratorDown)

	err = migratorDown.RunMigrations()
	assert.Nil(t, err)

	db.Where(&gormgrate.Migration{Successful: true}).Find(&gormgrate.Migration{}).Count(&count)
	assert.Equal(t, len(getFiles())-1, count)
}

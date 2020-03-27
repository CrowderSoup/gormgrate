package gormgrate

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// MigrationFile a migration file
type MigrationFile interface {
	Up(*gorm.DB) error
	Down(*gorm.DB) error
	Name() string
}

// Migrator migrates
type Migrator struct {
	DB         *gorm.DB
	Migrations []Migration
	MigrateUp  bool
	Files      []MigrationFile
}

// Migration a migration
type Migration struct {
	gorm.Model

	Name       string `gorm:"type:varchar(100);unique_index"`
	Successful bool
}

// NewMigrator returns a new Migrator
func NewMigrator(db *gorm.DB, up bool, files []MigrationFile) (*Migrator, error) {
	// Ensures that migrations table exists on first run
	db.AutoMigrate(&Migration{})

	var migrations []Migration
	if err := db.Find(&migrations).Error; err != nil {
		return nil, err
	}

	return &Migrator{
		DB:         db,
		Migrations: migrations,
		MigrateUp:  up,
		Files:      files,
	}, nil
}

// RunMigrations runs all our migrations (if they haven't been run successfully)
func (mig *Migrator) RunMigrations() error {
	if mig.MigrateUp {
		for _, migrationFile := range mig.Files {
			if !mig.shouldRun(migrationFile) {
				continue
			}

			migration, err := mig.insertIfNotExisting(migrationFile)
			if err != nil {
				return err
			}

			err = mig.runMigrationUp(migrationFile)
			if err != nil {
				return err
			}

			migration.Successful = true
			mig.updateMigration(migration)
		}
	} else {
		lastMigration := mig.Migrations[len(mig.Files)-1]
		migrationFile := mig.findFileByName(lastMigration.Name)

		err := mig.runMigrationDown(migrationFile, lastMigration)
		if err != nil {
			return err
		}
	}

	return nil
}

func (mig *Migrator) findFileByName(name string) MigrationFile {
	for _, file := range mig.Files {
		if file.Name() == name {
			return file
		}
	}

	return nil
}

func (mig *Migrator) runMigrationUp(m MigrationFile) error {
	fmt.Printf("Running Migration: %s\n", m.Name())

	// Run Migration
	err := m.Up(mig.DB)
	if err != nil {
		return err
	}

	return nil
}

func (mig *Migrator) runMigrationDown(mf MigrationFile, m Migration) error {
	fmt.Printf("Rolling back Migration: %s\n", mf.Name())

	// Roll back Migration
	err := mf.Down(mig.DB)
	if err != nil {
		return err
	}

	// Delete Migration
	if err := mig.DB.Unscoped().Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

func (mig *Migrator) shouldRun(m MigrationFile) bool {
	for _, migration := range mig.Migrations {
		if m.Name() == migration.Name && migration.Successful {
			return false
		}
	}

	return true
}

func (mig *Migrator) insertIfNotExisting(mf MigrationFile) (*Migration, error) {
	var existingMigration Migration

	for _, m := range mig.Migrations {
		if m.Name == mf.Name() {
			existingMigration = m
			break
		}
	}

	if existingMigration.Name != "" {
		return nil, nil
	}

	// Insert Migration
	migration := &Migration{
		Name:       mf.Name(),
		Successful: false,
	}
	if err := mig.DB.Create(migration).Error; err != nil {
		return nil, err
	}

	return migration, nil
}

func (mig *Migrator) updateMigration(m *Migration) error {
	m.Successful = true
	if err := mig.DB.Save(m).Error; err != nil {
		return err
	}

	return nil
}

package examples

import (
	"github.com/CrowderSoup/gormgrate"
	"github.com/jinzhu/gorm"
)

type exampleMigration struct {
	name string
}

// NewExampleMigration returns a migration file
func NewExampleMigration() gormgrate.MigrationFile {
	return &exampleMigration{
		name: "exampleMigration",
	}
}

// Name returns the name
func (m *exampleMigration) Name() string {
	return m.name
}

// Up runs the migration
func (m *exampleMigration) Up(db *gorm.DB) error {
	// Profile -> User
	if err := db.Model(&Profile{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT").Error; err != nil {
		return err
	}
}

// Down rolls back the migration
func (m *exampleMigration) Down(db *gorm.DB) error {
	// Profile -> User
	if err := db.Model(&Profile{}).RemoveForeignKey("user_id", "users(id)").Error; err != nil {
		return err
	}
}

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
	db.Model(&User{}).AddIndex("idx_user_email", "email")

	return nil
}

// Down rolls back the migration
func (m *exampleMigration) Down(db *gorm.DB) error {
	db.Model(&User{}).RemoveIndex("idx_user_email")

	return nil
}

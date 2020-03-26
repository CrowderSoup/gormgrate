# gormgrate

Managed manual migrations for GORM

## Usage 

`gormgrate` provides a `MigrationFile` interface. Each of your migrations must
satisfy this interface.

```go
type MigrationFile interface {
	Up(*gorm.DB) error
	Down(*gorm.DB) error
	Name() string
}
```

You can find an example migration in this repo at
`examples/migration_example.go`.

Once you've created a migration, you use `gormgate` in your project like so
(it's assumed you already have `db` initialized):

```go 
migrator, err := migrations.NewMigrator(db, true)
if err != nil {
    log.Fatal(err)
}

err = migrator.RunMigrations()
if err != nil {
	log.Fatal(err)
}
```

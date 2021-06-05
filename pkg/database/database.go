package database

import (
	"io/ioutil"
	"os"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type Config struct {
	Dsn string `json:"dsn"`
}

type Database struct {
	config *Config
	db     *sqlx.DB
	logger *zerolog.Logger
}

func NewDatabase(config *Config, logger *zerolog.Logger) *Database {
	return &Database{
		config: config,
		logger: logger,
	}
}

func (d *Database) Init() error {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = d.config.Dsn
	}
	db, err := sqlx.Connect("pgx", url)
	d.db = db
	return err
}

func (d *Database) Close() error {
	if d.db != nil {
		return d.db.Close()
	}
	return nil
}

func (d *Database) Migrate(migrationsPath string) error {
	d.logger.Info().Msg("Running migrations ...")
	files, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		d.logger.Error().Msg(err.Error())
		return err
	}

	for _, file := range files {
		data, err := ioutil.ReadFile(migrationsPath + "/" + file.Name())

		if err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		if _, err := d.db.Exec(string(data)); err != nil {
			d.logger.Error().Msg(err.Error())
			return err
		}

		d.logger.Info().Msg("- " + file.Name() + ": done.")
	}
	d.logger.Info().Msg("Migrated successfully.")
	return nil
}

func (d *Database) DB() *sqlx.DB {
	return d.db
}

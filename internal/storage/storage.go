package storage

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Storage struct {
	db *sql.DB
}

func New(dbURL string, migrationsPath string) (Storage, error) {
	db, err := sql.Open(pgDriver, dbURL)
	if err != nil {
		return Storage{}, err
	}

	err = db.Ping()
	if err != nil {
		return Storage{}, err
	}

	err = migrateDB(db, migrationsPath, pgDriver)
	if err != nil {
		return Storage{}, err
	}

	return Storage{
		db: db,
	}, nil
}

func (s Storage) GetModels(ctx context.Context) ([]Model, error) {
	q := `
		SELECT
			id,
			name,
			repository,
			version
		FROM mlist.models
	`
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	models := make([]Model, 0)
	err = sqlx.StructScan(rows, &models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s Storage) CreateModel(ctx context.Context, model Model) (Model, error) {
	q := `
		INSERT INTO mlist.models (
			name,
			repository,
			version
		) VALUES (
		 	$1,
			$2,
			$3
		) RETURNING
		 	id
	`
	var newID int
	err := s.db.QueryRowContext(ctx, q, model.Name).Scan(&newID)
	if err != nil {
		return Model{}, err
	}

	model.ID = newID
	return model, nil
}

func (s Storage) UpdateModel(ctx context.Context, modelID int, model Model) (Model, error) {
	q := `
        UPDATE mlist.models
        	SET name = $1,
				repository = $2,
				version = $3
        WHERE id = $4
		RETURNING
			id,
			name,
			repository,
			version
    `
	err := s.db.QueryRowContext(ctx, q, model.Name, model.Repository, model.Version, modelID).
		Scan(&model.ID, &model.Name, &model.Repository, &model.Version)
	if err != nil {
		return Model{}, err
	}

	return model, nil
}

func (s Storage) DeleteModel(ctx context.Context, modelID int) error {
	q := `
        DELETE FROM mlist.models
        WHERE id = $1
    `
	_, err := s.db.ExecContext(ctx, q, modelID)
	switch err {
	case sql.ErrNoRows:
		return nil
	default:
		return err
	}
}

func (s Storage) LookupModel(ctx context.Context, repositry string, version string) error {
	q := `
		SELECT
			1,
		FROM mlist.models
		WHERE repository = $1
			AND version = $2
	`
	_, err := s.db.QueryContext(ctx, q)
	return err
}

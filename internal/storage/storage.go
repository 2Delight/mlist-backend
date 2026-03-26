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
			name
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
			name
		) VALUES (
		 	$1
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

func (s Storage) UpdateModel(ctx context.Context, model Model) error {
	q := `
        UPDATE mlist.models
        	SET name = $1
        WHERE id = $2
    `
	_, err := s.db.ExecContext(ctx, q, model.Name, model.ID)
	return err
}

func (s Storage) DeleteModel(ctx context.Context, modelID int) error {
	q := `
        DELETE FROM mlist.models
        WHERE id = $1
    `
	_, err := s.db.ExecContext(ctx, q, modelID)
	return err
}

func (s Storage) LookupModel(ctx context.Context, repositry string, version string) (bool, error) {
	q := `
		SELECT
			id,
			name
		FROM mlist.models
	`
	_, err := s.db.QueryContext(ctx, q)
	switch err {
	case nil:
		return true, nil
	case sql.ErrNoRows:
		return false, nil
	default:
		return false, err
	}
}

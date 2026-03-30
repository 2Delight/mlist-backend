package models

import (
	"context"

	"github.com/2Delight/mlist-backend/internal/storage"
	"github.com/olegdayo/omniconv"
)

//go:generate mockery --name --case=snake --with-expecter --exported
type Storage interface {
	GetModels(ctx context.Context) ([]storage.Model, error)
	CreateModel(ctx context.Context, model storage.Model) (storage.Model, error)
	UpdateModel(ctx context.Context, modelID int, model storage.Model) (storage.Model, error)
	DeleteModel(ctx context.Context, modelID int) error
	LookupModel(ctx context.Context, repository string, version string) error
}

type Provider struct {
	storage Storage
}

func NewProvider(storage Storage) Provider {
	return Provider{
		storage: storage,
	}
}

func (p Provider) GetModels(ctx context.Context) ([]Model, error) {
	models, err := p.storage.GetModels(ctx)
	if err != nil {
		return nil, err
	}
	return omniconv.ConvertSlice(models, func(m storage.Model) Model { return Model(m) }), nil
}

func (p Provider) CreateModel(ctx context.Context, model Model) (Model, error) {
	resp, err := p.storage.CreateModel(ctx, storage.Model(model))
	if err != nil {
		return Model{}, err
	}
	return Model(resp), nil
}

func (p Provider) UpdateModel(ctx context.Context, modelID int, model Model) (Model, error) {
	resp, err := p.storage.UpdateModel(ctx, modelID, storage.Model(model))
	if err != nil {
		return Model{}, err
	}
	return Model(resp), nil
}

func (p Provider) DeleteModel(ctx context.Context, modelID int) error {
	return p.storage.DeleteModel(ctx, modelID)
}

func (p Provider) LookupModel(ctx context.Context, repository string, version string) error {
	return p.storage.LookupModel(ctx, repository, version)
}

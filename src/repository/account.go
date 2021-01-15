package repository

import (
	"github.com/base_skeleton_go/config"
	"github.com/base_skeleton_go/config/database"
)

type accountRepositoryCtx struct {
	cfg *config.Config
	DB  database.GormDatabase
}

type AccountRepository interface {
}

// NewAccountRepository ...
func NewAccountRepository(cfg *config.Config) AccountRepository {
	return &accountRepositoryCtx{
		cfg: cfg,
		DB:  cfg.DB(),
	}
}

package usecase

import (
	"github.com/base_skeleton_go/config"
	"github.com/base_skeleton_go/src/repository"
)

// Account ...
type Account interface {
}

type accountCtx struct {
	cfg  config.Config
	repo repository.AccountRepository
}

// NewAccountUc ...
func NewAccountUc(repo repository.AccountRepository, cfg config.Config) Account {
	return &accountCtx{cfg: cfg, repo: repo}
}

package registry

import (
	"sync"

	"github.com/base_skeleton_go/src/usecase"

	"github.com/base_skeleton_go/config"
)

// UsecaseRegistry ...
type UsecaseRegistry interface {
	Account() usecase.Account
}

type usecaseRegistry struct {
	repo RepositoryRegistry
	cfg  config.Config
}

// NewUsecaseRegistry ...
func NewUsecaseRegistry(cfg config.Config) (r UsecaseRegistry) {
	var ucRegistry usecaseRegistry
	var once sync.Once

	once.Do(func() {
		repoReg := NewRepoRegistry(cfg)
		ucRegistry = usecaseRegistry{repo: repoReg, cfg: cfg}
	})

	return &ucRegistry
}

func (u *usecaseRegistry) Account() usecase.Account {
	return nil
}

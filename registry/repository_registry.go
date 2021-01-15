package registry

import (
	"sync"

	"github.com/base_skeleton_go/config"
)

// RepositoryRegistry ...
type RepositoryRegistry interface {
}

type repositoryRegistry struct {
	cfg config.Config
}

// NewRepoRegistry ...
func NewRepoRegistry(cfg config.Config) RepositoryRegistry {
	var r repositoryRegistry
	var once sync.Once

	once.Do(func() {
		r = repositoryRegistry{cfg: cfg}
	})

	return r
}

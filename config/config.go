package config

import (
	"crypto/rsa"
	"os"
	"strconv"
	"time"

	"github.com/base_skeleton_go/config/database"
	"github.com/base_skeleton_go/config/keys"
	"github.com/base_skeleton_go/shared/logger"
)

// Config struct
type Config struct {
	serviceName string
	environment string
	debug       bool
	port        int

	token TokenConfig

	db database.GormDatabase
}

// TokenConfig ...
type TokenConfig struct {
	Name                string
	Type                string
	TemporaryLifetime   time.Duration
	AccessTokenLifeTime time.Duration
	PrivateKey          *rsa.PrivateKey
	PublicKey           *rsa.PublicKey
}

// NewConfig func
func NewConfig() *Config {
	cfg := new(Config)
	cfg.ConnectDB()

	return cfg
}

// ServiceName ...
func (c *Config) ServiceName() string {
	return os.Getenv(`SERVICE_NAME`)
}

// Environment ...
func (c *Config) Environment() string {
	return os.Getenv(`ENVIRONMENT`)
}

// Port func
func (c *Config) Port() int {
	v := os.Getenv("PORT")
	c.port, _ = strconv.Atoi(v)

	return c.port
}

// Debug func
func (c *Config) Debug() bool {
	v := os.Getenv("DEBUG")
	c.debug, _ = strconv.ParseBool(v)

	return c.debug
}

// ConnectDB func
func (c *Config) ConnectDB() {
	c.db = database.InitGorm()
}

// DB func
func (c *Config) DB() database.GormDatabase {
	return c.db
}

// InitToken ...
func (c *Config) InitToken() {
	temporaryLifetimeStr := os.Getenv("TOKEN_TEMPORARY_LIFETIME")
	temporaryLifetime, err := strconv.Atoi(temporaryLifetimeStr)
	if err != nil {
		logger.Panic(`ENV TOKEN_TEMPORARY_LIFETIME ERROR `, err)
	}

	accessTokenLifetimeStr := os.Getenv("ACCESS_TOKEN_LIFETIME")
	accessTokenLifetime, err := strconv.Atoi(accessTokenLifetimeStr)
	if err != nil {
		logger.Panic(`ENV ACCESS_TOKEN_LIFETIME ERROR `, err)
	}

	publicKey, err := keys.InitPublicKey()
	if err != nil {
		logger.Panic(err)
	}

	privateKey, err := keys.InitPrivateKey()
	if err != nil {
		logger.Panic(err)
	}

	c.token = TokenConfig{
		Name:                os.Getenv("TOKEN_NAME"),
		Type:                os.Getenv("TOKEN_TYPE"),
		TemporaryLifetime:   time.Duration(temporaryLifetime) * time.Second,
		AccessTokenLifeTime: time.Duration(accessTokenLifetime) * time.Hour,
		PrivateKey:          privateKey,
		PublicKey:           publicKey,
	}
}

// Token ...
func (c *Config) Token() TokenConfig {
	return c.token
}

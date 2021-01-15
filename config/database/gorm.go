package database

import (
	"fmt"
	logger2 "github.com/base_skeleton_go/shared/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)


type gormInstance struct {
	connDB *gorm.DB
}

// Master initialize DB for master data
func (g *gormInstance) ConnDB() *gorm.DB {
	return g.connDB
}

// GormDatabase abstraction
type GormDatabase interface {
	ConnDB() *gorm.DB
}

// InitGorm ...
func InitGorm() GormDatabase {
	inst := new(gormInstance)

	dbLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,         // Disable color
		},
	)

	gormConfig := &gorm.Config{
		// enhance performance config
		PrepareStmt: true,
		SkipDefaultTransaction: true,
		Logger: dbLogger,
	}

	// username, password, host, port, database
	dsnDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		os.Getenv("DB_USERNAME"),  os.Getenv("DB_PASSWORD"),  os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"), os.Getenv("DB_MASTER_NAME"))

	connDB, errConnDB := gorm.Open(mysql.Open(dsnDB), gormConfig)

	if errConnDB != nil {
		logger2.Panic(errConnDB)
	}

	inst.connDB = connDB
	
	return inst
}

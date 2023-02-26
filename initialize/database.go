package initialize

import (
	"blog/server/initialize/global"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct{}
var d = new(Database)

func DataBase() *gorm.DB {
	switch global.GLOBAL_CONFIG.Db.Dialects {
	case "psql":
		return d.DBPsql()
	default:
		return d.DBPsql()
	}
}


// InitializePsql
func (d *Database) DBPsql() *gorm.DB {
	// postgres config
	psqlConfig := postgres.Config{
		DSN: d.psqlDsn(),
	}
	
	db, err := gorm.Open(postgres.New(psqlConfig), d.gormConfig())
	if err != nil {
		panic(err)
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(global.GLOBAL_CONFIG.Db.MaxIdle)
	sqlDb.SetMaxOpenConns(global.GLOBAL_CONFIG.Db.MaxOpen)

	return db
}

func (d *Database) psqlDsn() string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		global.GLOBAL_CONFIG.Db.Host,
		global.GLOBAL_CONFIG.Db.Port,
		global.GLOBAL_CONFIG.Db.Username,
		global.GLOBAL_CONFIG.Db.Db,
		global.GLOBAL_CONFIG.Db.Password,
	)
}

func (d *Database) gormConfig() *gorm.Config {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,		// 
		  	LogLevel:      logger.Error, 	// Log level
		  	Colorful:      true,         	// color of print
		},
	  )
	  
	config := &gorm.Config{
		Logger: newLogger,
	}
	return config
}
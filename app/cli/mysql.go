package cli

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
	"github.com/douyacun/go-websocket-protobuf-ts/config"
)

var (
	mysqlOnce = sync.Once{}
	DB        *gorm.DB
)

func InitDB() {
	if config.Client.MysqlDsn() == "" {
		log.Errorf("mysql dsn empty !!!")
		return
	}
	mysqlOnce.Do(func() {
		var (
			err   error
			level = logger.Error
		)
		if config.App.IsTest() {
			level = logger.Info
		}
		DB, err = gorm.Open(mysql.Open(config.Client.MysqlDsn()), &gorm.Config{
			Logger: NewGormLogger(&logger.Config{
				LogLevel: level,
			}),
		})
		if err != nil {
			log.Errorf("gorm.open %v", err)
		}
	})
}

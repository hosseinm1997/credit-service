package infrastructures

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
)

var (
	dbsOnce sync.Once
	gormDB  *gorm.DB
)

func dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
		viper.Get("DB_HOST").(string),
		viper.Get("DB_USER").(string),
		viper.Get("DB_PASS").(string),
		viper.Get("DB_NAME").(string),
		viper.Get("DB_PORT").(string),
		viper.Get("DB_TIMEZONE").(string),
	)
}

func DatabaseSetup() {
	if gormDB == nil {
		dbsOnce.Do(func() {
			var err error
			gormDB, err = gorm.Open(postgres.Open(dsn()), &gorm.Config{
				//Logger: logger.Default.LogMode(logger.Info),
			})

			if err != nil {
				panic("can not connect to database")
			}

			sqlDB, err := gormDB.DB()
			if viper.IsSet("DB_MAX_OPEN_CONNECTIONS") {
				sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_OPEN_CONNECTIONS"))
			}

			if viper.IsSet("DB_MAX_IDLE_CONNECTIONS") {
				sqlDB.SetMaxOpenConns(viper.GetInt("DB_MAX_IDLE_CONNECTIONS"))
			}

			if err != nil {
				panic("can not set database configs")
			}
		})
	}
}

func DB() *gorm.DB {
	return gormDB
}

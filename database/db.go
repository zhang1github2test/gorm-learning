package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var GLOBALDB *gorm.DB

func init() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:my-secret-pw@tcp(192.168.188.101:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	GLOBALDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // true: Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		}),
	})
	if err != nil {
		panic(err)
	}
	sqlDB, err := GLOBALDB.DB()
	fmt.Println(sqlDB.Stats())

}

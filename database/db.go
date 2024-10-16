package database

import (
	"fmt"
	"github.com/zhang1github2test/gorm-learning/model"
	"github.com/zhang1github2test/gorm-learning/plugin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"log"
	"os"
	"time"
)

var GLOBALDB *gorm.DB

func init() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:my-secret-pw@tcp(192.168.188.101:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	dsn2 := "root:my-secret-pw@tcp(192.168.188.101:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	GLOBALDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // true: Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		}),
	})
	// 注册加解密的回调
	//callback.Register(GLOBALDB)
	GLOBALDB.Use(&plugin.Encrypt{})
	GLOBALDB.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsn)},
		Replicas: []gorm.Dialector{mysql.Open(dsn2)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}).Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(dsn2)},
		//Replicas: []gorm.Dialector{mysql.Open(dsn2)},
		// sources/replicas load balancing policy
		Policy: dbresolver.RandomPolicy{},
		// print sources/replicas mode in logger
		TraceResolverMode: true,
	}, &model.Student{}))
	if err != nil {
		panic(err)
	}
	sqlDB, err := GLOBALDB.DB()
	fmt.Println(sqlDB.Stats())
}

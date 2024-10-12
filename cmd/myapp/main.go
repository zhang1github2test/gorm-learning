package main

import (
	"fmt"
	"github.com/zhang1github2test/gorm-learning/callback"
	. "github.com/zhang1github2test/gorm-learning/database"
	"github.com/zhang1github2test/gorm-learning/model"
	"time"
)

func main() {
	test_aes()
}
func test_aes() {
	for i := 0; i < 10; i++ {
		fmt.Println(callback.AesEncrypt([]byte("zhangshenglu"), callback.DATA_KEY))
	}
	callback.AesEncrypt([]byte("zhangshenglu"), callback.DATA_KEY)
	str, err := callback.AesDecrypt("Y7At/UvknzizhdHq2d1MwPSLBarA9HDnYUxoA4ZyBTg=", []byte("0123456789123456"))
	if err != nil {
		panic(err)
	}
	fmt.Println(str)
}

func test_getDB() {
	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, err := GLOBALDB.DB()
	if err != nil {
		panic(err)
	}

	// Ping
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
	GLOBALDB.AutoMigrate(&model.User{})
	email := "89954554554@163.com"
	birthday := time.Now().AddDate(-18, 0, 0)

	user := &model.User{
		ID:       100003,
		Name:     "zhangshenglu",
		Email:    &email,
		Age:      18,
		Birthday: &birthday,
	}
	GLOBALDB.Save(user)
}

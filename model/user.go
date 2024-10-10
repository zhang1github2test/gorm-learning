package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID        uint       `gorm:"type:int;comment:主键" json:"id"`      // Standard field for the primary key
	Name      string     `gorm:"size:128;comment:人员姓名" json:"name"`  // 一个常规字符串字段
	Email     *string    `gorm:"size:128;comment:邮箱地址" json:"email"` // 一个指向字符串的指针, allowing for null values
	Age       uint8      `json:"age"`                                // 一个未签名的8位整数
	Phone     string     `gorm:"size:11;comment:手机号码" json:"phone"`  // 手机号码
	Birthday  *time.Time `json:"birthday"`                           // A pointer to time.Time, can be null
	CreatedAt *time.Time `json:"created_at"`                         // 创建时间（由GORM自动管理）
	UpdatedAt *time.Time `json:"updated_at"`                         // 最后一次更新时间（由GORM自动管理）
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	if u.CreatedAt != nil {
		u.CreatedAt = nil
	}
	return nil
}

func (u *User) AfterUpdate(tx *gorm.DB) error {
	log.Println("BeforeSave ...")
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	log.Println("BeforeSave ...")
	return nil
}

func (u *User) AfterSave(tx *gorm.DB) error {
	log.Println("AfterSave ...")
	return nil
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	log.Println("BeforeCreate ...")
	// 将结构体转为 JSON 格式
	userJSON, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return err
	}

	// 输出 JSON 格式
	log.Println(string(userJSON))
	// 返回错误信息后，GORM 将停止后续的操作并回滚事务
	// return errors.New("BeforeCreate")
	return nil
}
func (u *User) AfterCreate(tx *gorm.DB) error {
	log.Println("AfterCreate ...")
	return nil
}

func (u *User) AfterFind(tx *gorm.DB) error {
	ctx := tx.Statement.Context
	s, ok := ctx.Value("pageSize").(int)
	log.Printf("AfterFind pageSize:%v,ok:%v", s, ok)
	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) error {
	log.Println("AfterDelete ...")
	return nil
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	log.Println("BeforeDelete ...")
	return nil
}

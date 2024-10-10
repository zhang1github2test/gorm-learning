package model

import "time"

type Student struct {
	ID           uint       `gorm:"type:int;comment:主键" json:"id"`     // Standard field for the primary key
	Name         string     `gorm:"size:128;comment:学生姓名" json:"name"` // 一个常规字符串字段
	GuardianName string     `gorm:"size:128;comment:监护人姓名" json:"guardian_name"`
	Age          uint8      `json:"age"`                               // 一个未签名的8位整数
	Phone        string     `gorm:"size:11;comment:手机号码" json:"phone"` // 手机号码
	Birthday     *time.Time `json:"birthday"`                          // A pointer to time.Time, can be null
	CreatedAt    *time.Time `json:"created_at"`                        // 创建时间（由GORM自动管理）
	UpdatedAt    *time.Time `json:"updated_at"`                        // 最后一次更新时间（由GORM自动管理）
}

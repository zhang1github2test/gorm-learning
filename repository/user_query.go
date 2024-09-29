package repository

import (
	"github.com/zhang1github2test/gorm-learning/model"
	"gorm.io/gorm"
)

type UserDao struct {
	Db *gorm.DB
}

// Create 使用Create方法保存user
func (usedao *UserDao) Create(user *model.User) (int64, error) {
	db := usedao.Db.Create(user)
	return db.RowsAffected, db.Error
}

// Save 使用Save方法保存user
func (usedao *UserDao) Save(user *model.User) (int64, error) {
	db := usedao.Db.Save(user)
	return db.RowsAffected, db.Error
}

// CreateInBatches 批量保存数据
func (userDao *UserDao) CreateInBatches(users *[]model.User, batchSize int) (int64, error) {
	db := userDao.Db.CreateInBatches(users, batchSize)
	return db.RowsAffected, db.Error
}

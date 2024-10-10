package repository

import (
	"github.com/zhang1github2test/gorm-learning/model"
	"gorm.io/gorm"
)

type StudentDao struct {
	Db *gorm.DB
}

func (stDao *StudentDao) Create(student *model.Student) error {
	tx := stDao.Db.Create(student)
	return tx.Error
}

func (stDao *StudentDao) First(student *model.Student) error {
	tx := stDao.Db.First(student)
	return tx.Error
}

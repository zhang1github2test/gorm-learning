package repository

import (
	"context"
	"errors"
	"github.com/zhang1github2test/gorm-learning/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
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
	// 类似SQL: INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`created_at`,`updated_at`)
	// VALUES (?,?,?,?,?,?),(?,?,?,?,?,?),(?,?,?,?,?,?)
	db := userDao.Db.CreateInBatches(users, batchSize)
	return db.RowsAffected, db.Error
}

// First 使用主键升序的方式获取到第一条记录
func (userDao *UserDao) First(user *model.User) (int64, error) {
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	db := userDao.Db.First(user)
	return db.RowsAffected, db.Error
}

// Take 不排序直接获取第一条记录
func (userDao *UserDao) Take(user *model.User) (int64, error) {
	// SELECT * FROM `users`  LIMIT 1
	db := userDao.Db.Take(user)
	return db.RowsAffected, db.Error
}

// Last 使用主键降序序的方式获取到第一条记录
func (userDao *UserDao) Last(user *model.User) (int64, error) {
	// SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	db := userDao.Db.Last(user)
	return db.RowsAffected, db.Error
}

// Find 查找所有对象
func (userDao *UserDao) Find(users *[]model.User) (int64, error) {
	db := userDao.Db.Find(users)
	return db.RowsAffected, db.Error
}

// Scan 查找所有对象
func (userDao *UserDao) Scan(users *[]model.User) (int64, error) {
	//db := userDao.Db.Table("users").Scan(users)
	db := userDao.Db.Raw("SELECT * FROM `users`").Scan(users)
	return db.RowsAffected, db.Error
}

// SelectSpecField 查找所有对象
func (userDao *UserDao) SelectSpecField(users *[]model.User, fields ...string) (int64, error) {
	db := userDao.Db.Select("name", "age", fields).Find(users)
	return db.RowsAffected, db.Error
}

// StringQuery  使用where方法增加查询的条件 : where方法的第一个入参为条件语句，后面为查询参数
func (userDao *UserDao) StringQuery(users *[]model.User) (int64, error) {
	db := userDao.Db.Where("name = ?", "zhangshenglu1").Find(&users)
	return db.RowsAffected, db.Error
}

// StructQuery  使用where方法增加查询的条件 : where方法的第一个入参为模型对应的结构体
func (userDao *UserDao) StructQuery(users *[]model.User, user *model.User) (int64, error) {
	db := userDao.Db.Where(user).Find(&users)
	return db.RowsAffected, db.Error
}

// MapQuery  使用where方法增加查询的条件 : where方法的第一个入参为Map对象
func (userDao *UserDao) MapQuery(users *[]model.User) (int64, error) {
	// SELECT * FROM `users` WHERE `Name` = "zhangshenglu1"
	db := userDao.Db.Where(map[string]interface{}{"Name": "zhangshenglu1"}).Find(&users)
	return db.RowsAffected, db.Error
}

// NotQuery  使用Not方法增加查询的条件
func (userDao *UserDao) NotQuery(users *[]model.User) error {
	// SELECT * FROM `users` WHERE `Name` <> 'zhangshenglu1'
	db := userDao.Db.Not(map[string]interface{}{"Name": "zhangshenglu1"}).Find(&users)

	// SELECT * FROM `users` WHERE `Name` <> 'zhangshenglu1'
	db = userDao.Db.Not(&model.User{Name: "zhangshenglu1"}).Find(&users)
	// SELECT * FROM `users` WHERE NOT name = 'zhangshenglu1'
	db = userDao.Db.Not("name = ?", "zhangshenglu1").Find(&users)

	// SELECT * FROM `users` WHERE `Name` NOT IN ('zhangshenglu1','zhangshenglu2')
	db = userDao.Db.Not(map[string]interface{}{"Name": []string{"zhangshenglu1", "zhangshenglu2"}}).Find(&users)
	return db.Error
}

// OrQuery  使用Or方法增加查询的条件
func (userDao *UserDao) OrQuery(users *[]model.User) error {
	// SELECT * FROM `users` WHERE age = 0 OR `Name` = 'zhangshenglu1'
	db := userDao.Db.Where("age = ?", 0).Or(map[string]interface{}{"Name": "zhangshenglu1"}).Find(&users)

	// SELECT * FROM `users` WHERE age = 0 OR `users`.`name` = 'zhangshenglu1'
	db = userDao.Db.Where("age = ?", 0).Or(&model.User{Name: "zhangshenglu1"}).Find(&users)
	// SELECT * FROM `users` WHERE age = 0 OR name = 'zhangshenglu1'
	db = userDao.Db.Where("age = ?", 0).Or("name = ?", "zhangshenglu1").Find(&users)

	// SELECT * FROM `users` WHERE age = 0 OR `Name` IN ('zhangshenglu1','zhangshenglu2')
	db = userDao.Db.Where("age = ?", 0).Or(map[string]interface{}{"Name": []string{"zhangshenglu1", "zhangshenglu2"}}).Find(&users)
	return db.Error
}

// Order 使用Order方法对查询的数据进行排序
func (userdao *UserDao) Order(users *[]model.User) error {
	// SELECT * FROM `users` WHERE `name` = 'zhangshenglu' ORDER BY name desc,age
	tx := userdao.Db.Order("name desc,age").Where("name", "zhangshenglu1").Find(users)
	return tx.Error
}

// Group
func (userdao *UserDao) Group() error {
	type result struct {
		Name  string
		Count int
	}
	re := []result{}
	// SELECT name,count(name) as count FROM `users` WHERE `name` = 'zhangshenglu1' GROUP BY `name` ORDER BY name desc
	tx := userdao.Db.Table("users").Select("name,count(name) as count").Order("name desc").Where("name", "zhangshenglu1").Group("name").Find(&re)
	return tx.Error
}

// Having
func (userdao *UserDao) Having() error {
	type result struct {
		Name  string
		Count int
	}
	re := []result{}
	// SELECT name,count(name) as count FROM `users` WHERE `name` = 'zhangshenglu1' GROUP BY `name` HAVING count > 5
	tx := userdao.Db.Table("users").Select("name,count(name) as count").Where("name", "zhangshenglu1").Group("name").Having("count > ?", 5).Find(&re)
	return tx.Error
}

// PageQuery 演示分页查询
func (userdao *UserDao) PageQuery(offset, limit int) error {
	type result struct {
		users []model.User
		Total int64
	}
	re := result{}
	// SELECT * FROM `users` LIMIT 10 OFFSET 10
	tx := userdao.Db.Table("users").Count(&re.Total)
	tx = userdao.Db.Table("users").Offset(offset).Limit(limit).Find(&re.users)
	return tx.Error
}

// 更新相关操作开始

// UpdateSingle 更新单个字段
func (userDao *UserDao) UpdateSingle() error {
	// UPDATE `users` SET `name`='zhangshenglu' WHERE id = 100003
	tx := userDao.Db.Table("users").Where("id = ?", 100003).Update("name", "zhangshenglu")
	return tx.Error
}

func (userDao *UserDao) Updates() error {
	// 使用结构体进行更新
	// UPDATE `users` SET `name`='zhangshenglu3',`age`=18,`updated_at`='2024-09-30 17:00:50.08' WHERE id = 100003
	tx := userDao.Db.Table("users").Where("id = ?", 100003).Updates(&model.User{
		Name: "zhangshenglu3",
		Age:  18,
	})

	// 使用select 来选择需要更新的字段
	// UPDATE `users` SET `name`='zhangshenglu3',`updated_at`='2024-09-30 17:02:48.623' WHERE id = 100003
	tx = userDao.Db.Table("users").Select("name").Where("id = ?", 100003).Updates(&model.User{
		Name: "zhangshenglu3",
		Age:  18,
	})

	// 使用map结构更新   UPDATE `users` SET `name`='zhangshenglu4',`updated_at`='2024-09-30 17:23:06.757' WHERE id = 100003
	tx = userDao.Db.Model(&model.User{}).Select("name").Where("id = ?", 100003).Updates(map[string]interface{}{
		"name": "zhangshenglu4",
		"age":  22,
	})
	return tx.Error
}

// 更新相关操作结束
func (userDao *UserDao) Delete() error {
	user := model.User{
		ID: 100003,
	}
	// 基于主键删除
	// DELETE FROM `users` WHERE `users`.`id` = 100003
	tx := userDao.Db.Delete(user)

	// 基于查询条件删除所有匹配到的数据
	// DELETE FROM `users` WHERE name like '%zhangshenglu%'
	//tx = userDao.Db.Where("name like ?", "%zhangshenglu%").Delete(&model.User{})
	re := clause.Returning{Columns: []clause.Column{{Name: "name"}}}
	tx = userDao.Db.Clauses(&re).Where("name like ?", "%zhangshenglu%").Delete(&model.User{})
	return tx.Error
}

func (userdao *UserDao) AutoMigrate() error {
	return userdao.Db.AutoMigrate(&model.User{})
}

func (userdao *UserDao) Transaction(f func(tx *gorm.DB, parm interface{}) error, parm interface{}) error {
	return userdao.Db.Transaction(func(tx *gorm.DB) error {
		return f(tx, parm)
	})
}

func (userdao *UserDao) SaveWithTransactionByManual(user *model.User, rollback bool) error {
	db := userdao.Db
	//开启事务
	db.Begin()
	db.Save(user)
	if rollback {
		// 模型出现异常,进行数据回滚
		db.Rollback()
		return errors.New("发生异常，将数据进行回滚!")
	}

	// 成功执行，提交事务
	db.Commit()
	return nil
}

// SingleSessionContext 单会话模式使用Context示例
func (userdao *UserDao) SingleSessionContext(user *model.User) error {
	// 设置查询的时候超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 设置其他的上下文信息
	ctx = context.WithValue(ctx, "pageSize", 10)
	defer cancel()
	tx := userdao.Db.WithContext(ctx).First(user)
	return tx.Error
}

// SessionContext 持续会话模式使用Context示例
func (userdao *UserDao) SessionContext(user *model.User) error {
	// 设置查询的时候超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	// 设置其他的上下文信息
	ctx = context.WithValue(ctx, "pageSize", 10)
	defer cancel()

	tx := userdao.Db.WithContext(ctx)
	tx.First(user)
	tx.First(&model.User{})
	return tx.Error
}

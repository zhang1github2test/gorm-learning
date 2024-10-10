### GORM 学习

---

### 1. **GORM介绍**

官网地址：https://gorm.io/zh_CN/docs/index.html

GORM 提供了简单的 API，使得开发者可以使用 Go 语言的结构体来表示数据库表，减少了 SQL 语句的直接书写。

GORM 是 Golang 项目中一个强大的工具，能够有效提升数据库操作的效率和可维护性。

- **特性**：
    
    * 全功能 ORM
    * Create，Save，Update，Delete，Find 中钩子方法
    
    * 事务，嵌套事务，Save Point，Rollback To Saved Point
    * Auto Migration

---

### 2. **环境配置**

* **项目结构**：

​	初始化项目：

```sh
 go mod init github.com/zhang1github2test/gorm-learning
```

建立如下目录：

```txt
E:.
│  go.mod
│  go.sum
│  README.md
├─cmd
│  └─myapp
│          main.go
│
├─database
│      db.go
│
├─model
└─repository
```

- **数据库准备**：使用 Docker安装mysql数据库，见《docker安装mysql.md》

- **依赖安装**：GORM 和数据库驱动（如 MySQL、PostgreSQL）。

    ```sh
     go get -u gorm.io/gorm
     go get -u gorm.io/driver/mysql
    ```

    

---

### 3. **GORM 基本操作**

#### 数据库连接

- 使用 GORM 连接数据库，演示 DSN 字符串的配置。

导入完成后，需要mysql地址信息，整个链接格式如下：

```
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```

各个参数说明如下：

username：mysql数据的账号名称

password：mysql数据的账号密码

protocol：协议

address：host+port

dbname: 数据库名

param1： 连接的参数1

value1： 连接的参数1对应的值

连接地址的举例：root:my-secret-pw@tcp(192.168.188.101:3306)/test?charset=utf8mb4&parseTime=True&loc=Local

**完整示例** ：

```go
package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBALDB *gorm.DB

func init() {
	var err error
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:my-secret-pw@tcp(192.168.188.101:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	GLOBALDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := GLOBALDB.DB()
	fmt.Println(sqlDB.Stats())

}
```

测试连接：

```go
package main

import (
	"database/sql"
	. "github.com/zhang1github2test/gorm-learning/database"
	"github.com/zhang1github2test/gorm-learning/model"
	"time"
)

func main() {
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
	actvie := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	birthday := time.Now().AddDate(-18, 0, 0)

	memberNumber := sql.NullString{
		String: "10",
		Valid:  true,
	}
	user := &model.User{
		ID:           100003,
		Name:         "zhangshenglu",
		Email:        &email,
		Age:          18,
		Birthday:     &birthday,
		ActivatedAt:  actvie,
		MemberNumber: memberNumber,
	}
	GLOBALDB.Save(user)
}

```





#### CRUD 

定义userDao结构体

```
type UserDao struct {
	Db *gorm.DB
}
```

##### Create操作

* **单条插入** 

  Create()方法插入单条数据

  `Create` 方法用于**插入新的记录**到数据库。

  ```go
  // Create 使用Create方法保存user
  func (usedao *UserDao) Create(user *model.User) (int64, error) {
  	db := usedao.Db.Create(user)
  	return db.RowsAffected, db.Error
  }
  ```

  测试Create()单条插入

  ```go
  package repository
  
  import (
  	"github.com/zhang1github2test/gorm-learning/database"
  	"github.com/zhang1github2test/gorm-learning/model"
  	"gorm.io/gorm"
  	"testing"
  	"time"
  )
  
  func TestUserDao_Create(t *testing.T) {
  	type fields struct {
  		Db *gorm.DB
  	}
  	type args struct {
  		user *model.User
  	}
  	field := fields{
  		Db: database.GLOBALDB,
  	}
  	email := "89954554554@163.com"
  	birthday := time.Now().AddDate(-18, 0, 0)
  
  	user := &model.User{
  		ID:       100003,
  		Name:     "zhangshenglu",
  		Email:    &email,
  		Age:      18,
  		Birthday: &birthday,
  	}
  
  	arg := args{
  		user: user,
  	}
  
  	tests := []struct {
  		name    string
  		fields  fields
  		args    args
  		want    int64
  		wantErr bool
  	}{
  		{
  			name:    "createFirst",
  			fields:  field,
  			args:    arg,
  			want:    1,
  			wantErr: false,
  		},
  		{
  			name:    "createAgain",
  			fields:  field,
  			args:    arg,
  			want:    0,
  			wantErr: true,
  		},
  	}
  	for _, tt := range tests {
  		t.Run(tt.name, func(t *testing.T) {
  			usedao := &UserDao{
  				Db: tt.fields.Db,
  			}
  			got, err := usedao.Create(tt.args.user)
  			if (err != nil) != tt.wantErr {
  				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
  				return
  			}
  			if got != tt.want {
  				t.Errorf("Create() got = %v, want %v", got, tt.want)
  			}
  		})
  	}
  }
  
  ```

  Create方法只能数据进行插入操作，如果对应数据已经存在那么就会执行失败。

  如在上面的测试用例中，我们将100003这条数据重复插入两次，第一次我们期望能够成功插入，第二次则会插入失败。

  Save()方法插入单条数据

  ```go
  // Save 使用Save方法保存user
  func (usedao *UserDao) Save(user *model.User) (int64, error) {
  	db := usedao.Db.Save(user)
  	return db.RowsAffected, db.Error
  }
  ```

  测试Save()方法

  `Save` 方法用于**保存**一条记录，可以是新记录，也可以是更新现有的记录。

  ```go
  func TestUserDao_Save(t *testing.T) {
  	type fields struct {
  		Db *gorm.DB
  	}
  	type args struct {
  		user *model.User
  	}
  
  	field := fields{
  		Db: database.GLOBALDB,
  	}
  	email := "89954554554@163.com"
  	birthday := time.Now().AddDate(-18, 0, 0)
  
  	user := &model.User{
  		ID:       100002,
  		Name:     "zhangshenglu",
  		Email:    &email,
  		Age:      18,
  		Birthday: &birthday,
  	}
  
  	arg := args{
  		user: user,
  	}
  
  	tests := []struct {
  		name    string
  		fields  fields
  		args    args
  		want    int64
  		wantErr bool
  	}{
  		{
  			name:    "SaveFirst",
  			fields:  field,
  			args:    arg,
  			want:    1,
  			wantErr: false,
  		},
  		{
  			name:    "SaveAgain",
  			fields:  field,
  			args:    arg,
  			want:    1,
  			wantErr: false,
  		},
  	}
  	for _, tt := range tests {
  		t.Run(tt.name, func(t *testing.T) {
  			usedao := &UserDao{
  				Db: tt.fields.Db,
  			}
  			got, err := usedao.Save(tt.args.user)
  			if (err != nil) != tt.wantErr {
  				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
  				return
  			}
  			if got != tt.want {
  				t.Errorf("Save() got = %v, want %v", got, tt.want)
  			}
  		})
  	}
  }
  ```

  上面的测试用例中将同一条数据分执行两次Save操作，我们预期两次都会成功。

* **批量插入**

  `CreateInBatches` 方法用于批量插入数据到数据库中。当你需要一次性插入大量数据时，使用 `CreateInBatches` 比普通的 `Create` 方法效率更高，因为它减少了数据库的交互次数。

  ```go
  // CreateInBatches 批量保存数据
  func (userDao *UserDao) CreateInBatches(users *[]model.User, batchSize int) (int64, error) {
  	db := userDao.Db.CreateInBatches(users, batchSize)
  	return db.RowsAffected, db.Error
  }
  ```

  测试代码

  ```go
  func TestUserDao_CreateInBatches(t *testing.T) {
  	type fields struct {
  		Db *gorm.DB
  	}
  	type args struct {
  		users     *[]model.User
  		batchSize int
  	}
  	us := []model.User{{Name: "zhangshenglu1"}, {Name: "zhangshenglu2"}, {Name: "zhangshenglu3"}}
  	arg := args{
  		users: &us,
  		batchSize: 2,
  	}
  	field := fields{
  		Db: database.GLOBALDB,
  	}
  	tests := []struct {
  		name    string
  		fields  fields
  		args    args
  		want    int64
  		wantErr bool
  	}{
  		{
  			name:    "batchSave",
  			fields:  field,
  			args:    arg,
  			want:    3,
  			wantErr: false,
  		},
  	}
  	for _, tt := range tests {
  		t.Run(tt.name, func(t *testing.T) {
  			userDao := &UserDao{
  				Db: tt.fields.Db,
  			}
  			got, err := userDao.CreateInBatches(tt.args.users, tt.args.batchSize)
  			if (err != nil) != tt.wantErr {
  				t.Errorf("CreateInBatches() error = %v, wantErr %v", err, tt.wantErr)
  				return
  			}
  			if got != tt.want {
  				t.Errorf("CreateInBatches() got = %v, want %v", got, tt.want)
  			}
  		})
  	}
  }
  ```

##### Select操作

查询单条、多条记录，条件查询、分页、排序。

###### 检索单个对象

​	GORM 提供了 `First`、`Take`、`Last` 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 `LIMIT 1` 条件，且没有找到记录时，它会返回 `ErrRecordNotFound` 错误。

**first方法**

使用主键升序的方式获取到第一条记录

```go
// First 使用主键升序的方式获取到第一条记录
func (userDao *UserDao) First(user *model.User) (int64, error) {
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	db := userDao.Db.First(user)
	return db.RowsAffected, db.Error
}
```

测试用例：

```go

type fields struct {
	Db *gorm.DB
}
type args struct {
	user *model.User
}

var field = fields{
	Db: database.GLOBALDB,
}



func TestUserDao_First(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "First",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.First(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("First() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("First() got = %v, want %v", got, tt.want)
			}
		})
	}
}
```

**Take方法**

获取一条记录，没有指定排序字段。

```go
// Take 不排序直接获取第一条记录
func (userDao *UserDao) Take(user *model.User) (int64, error) {
	// SELECT * FROM `users`  LIMIT 1
	db := userDao.Db.Take(user)
	return db.RowsAffected, db.Error
}
```

测试用例：

```go
func TestUserDao_Take(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Take_test",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Take(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Take() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Take() got = %v, want %v", got, tt.want)
			}
		})
	}
}
```

**last方法**

按照主键降序获取第一条数据。

```go
// Last 使用主键降序序的方式获取到第一条记录
func (userDao *UserDao) Last(user *model.User) (int64, error) {
	// SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	db := userDao.Db.Last(user)
	return db.RowsAffected, db.Error
}
```

测试用例：

```go
func TestUserDao_Last(t *testing.T) {
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Last",
			fields: field,
			args: args{
				user: &model.User{},
			},
			want:    1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Last(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("Last() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Last() got = %v, want %v", got, tt.want)
			}
		})
	}
}
```

###### 检索多个对象

Find()方法用于从数据库中查询多条记录，并将查询结果填充到一个切片或数组中。它可以通过查询条件检索数据，并支持复杂的查询结构。

```go
// Find 查找所有对象
func (userDao *UserDao) Find(users *[]model.User) (int64, error) {
	db := userDao.Db.Find(users)
	return db.RowsAffected, db.Error
}
```

测试用例

```go
func TestUserDao_Find(t *testing.T) {
	type args struct {
		users *[]model.User
	}
	us := &[]model.User{}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:   "Find",
			fields: field,
			args: args{
				users: us,
			},
			want:    24,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userDao := &UserDao{
				Db: tt.fields.Db,
			}
			got, err := userDao.Find(tt.args.users)
			if (err != nil) != tt.wantErr {
				t.Errorf("Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Find() got = %v, want %v", got, tt.want)
			}
		})
	}

}
```

Scan方法来查询所有记录

```go
// Scan 查找所有对象
func (userDao *UserDao) Scan(users *[]model.User) (int64, error) {
	db := userDao.Db.Table("users").Scan(users)
	return db.RowsAffected, db.Error
}
```

`Scan` 并不要求将查询结果填充到模型结构体中，任何自定义结构体都可以通过 `Scan` 接收结果。它更多用于执行复杂的 SQL 查询，比如联表查询、选择特定字段等。但是得使用.Table方法指定表，或者使用

###### 查询指定字段Select

​	`Select()` 方法用于指定查询中需要返回的字段或列。它允许选择特定的数据库列而不是返回所有列，从而优化查询性能。`Select()` 方法可以接受字符串、结构体字段或表达式，并且支持 SQL 函数等复杂表达式。它常用于控制查询结果的内容，使查询更高效或满足特定的业务需求

```go
// SelectSpecField 查找所有对象
func (userDao *UserDao) SelectSpecField(users *[]model.User, fields ...string) (int64, error) {
	db := userDao.Db.Select("name", "age",fields).Find(users)
	return db.RowsAffected, db.Error
}
```

###### 条件查询 where

​	`Where()` 方法用于构建 SQL 查询的 `WHERE` 子句，筛选符合特定条件的记录。它支持多种查询形式，包括直接传递条件字符串、使用参数化查询、通过结构体筛选字段，或使用 `map` 形式传递字段条件。`Where()` 方法可以用于精确匹配、范围查询、逻辑运算等，并能够与其他查询方法如 `Or()`、`Not()` 等组合使用，以构建复杂的查询条件。

**string条件查询**

```go
// StringQuery  使用where方法增加查询的条件 : where方法的第一个入参为条件语句，后面为查询参数
func (userDao *UserDao) StringQuery(users *[]model.User) (int64, error) {
	db := userDao.Db.Where("name = ?", "zhangshenglu").Find(&users)
	return db.RowsAffected, db.Error
}
```

**Struct 条件查询**

```go
// StructQuery  使用where方法增加查询的条件 : where方法的第一个入参为条件语句，后面为查询参数
func (userDao *UserDao) StructQuery(users *[]model.User, user *model.User) (int64, error) {
	db := userDao.Db.Where(user).Find(&users)
	return db.RowsAffected, db.Error
}
```

注意：

```txt
在使用结构体查询的时候，如果字段的值为零值，那么GORM不会用其来进行构建查询条件的。如果需要基于零值进行查询，需要使用Map条件进行查询。

db.Where(&User{Name: "zhangshenglu1", Age: 0}).Find(&users)
// SELECT * FROM users WHERE name = "zhangshenglu1";
```



**Map条件查询**

```go
// MapQuery  使用where方法增加查询的条件 : where方法的第一个入参为Map对象
func (userDao *UserDao) MapQuery(users *[]model.User) (int64, error) {
	db := userDao.Db.Where(map[string]interface{}{"Name": "zhangshenglu1"}).Find(&users)
	return db.RowsAffected, db.Error
}
```

###### Not 条件查询

`Not()` 方法用于构建 SQL 查询中的 `NOT` 条件，用来查询不满足特定条件的记录。它通常与 `Where()` 类似，只是用来排除符合条件的记录

```go
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
```

###### Or 条件查询

`Or()` 方法用于构建 SQL 查询中的 `OR` 条件，允许你在查询中添加多个条件，其中一个条件为真时，记录会被查询出来。

```go
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
```

###### 查询排序Order

`Order` 方法用于对查询结果进行排序，类似于 SQL 中的 `ORDER BY` 子句。你可以根据一个或多个字段对数据进行升序或降序排列。

```go
// Order 使用Order方法对查询的数据进行排序
func (userdao *UserDao) Order(users *[]model.User) error {
	// SELECT * FROM `users` WHERE `name` = 'zhangshenglu' ORDER BY name desc,age
	tx := userdao.Db.Order("name desc,age").Where("name", "zhangshenglu").Find(users)
	return tx.Error
}
```

###### 分组查询 Group

`Group` 方法用于对查询结果进行分组，类似于 SQL 中的 `GROUP BY` 子句。它允许你根据指定的字段对记录进行分组，然后可以结合聚合函数（如 `COUNT`, `SUM`, `AVG`, `MAX`, `MIN`）来对分组后的数据进行操作。

```go
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
```

###### 分组查询 Having

`Having()` 方法用于在 `GROUP BY` 查询后过滤分组数据。它与 `WHERE` 类似，但 `WHERE` 作用于分组之前的数据，而 `Having()` 则用于对分组后的聚合结果进行条件过滤

```go
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
```

###### 分页查询

在执行分页查询的时候通常会使用到Offset、Limit、Count方法

**`Limit()`**：用于限制查询结果的数量，即只返回指定数量的记录。

**`Offset()`**：用于跳过指定数量的记录，常用于分页查询。

**`Count()`**：用于统计查询结果的总数，不会影响查询到的实际记录，只返回记录的数量

```go
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

```



##### Update操作

###### 更新单个字段

`Update` 方法用于更新数据库中已存在记录的指定字段值。

```go
// UpdateSingle 更新单个字段
func (userDao *UserDao) UpdateSingle() error {
	// UPDATE `users` SET `name`='zhangshenglu' WHERE id = 100003
	tx := userDao.Db.Table("users").Where("id = ?", 100003).Update("name", "zhangshenglu")
	return tx.Error
}
```

###### 更新多个字段



```go
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
```



##### Delete操作

`Delete` 方法用于从数据库中删除记录。可以根据主键或指定的条件删除一条或多条记录

```go
func (userDao *UserDao) Delete() error {
	user := model.User{
		ID: 100003,
	}
	// 基于主键删除
	// DELETE FROM `users` WHERE `users`.`id` = 100003
	tx := userDao.Db.Delete(user)
	
	// 基于查询条件删除所有匹配到的数据
	// DELETE FROM `users` WHERE name like '%zhangshenglu%'
	tx = userDao.Db.Where("name like ?", "%zhangshenglu%").Delete(&model.User{})
	return tx.Error
}
```





#### 模型定义

- 如何定义模型及其与数据库表的映射（字段类型、标签）

  ```go
  type User struct {
  	ID        uint       `gorm:"type:int;comment:主键"`   // Standard field for the primary key
  	Name      string     `gorm:"size:128;comment:人员姓名"` // 一个常规字符串字段
  	Email     *string    `gorm:"size:128;comment:邮箱地址"` // 一个指向字符串的指针, allowing for null values
  	Age       uint8      // 一个未签名的8位整数
  	Birthday  *time.Time // A pointer to time.Time, can be null
  	CreatedAt *time.Time // 创建时间（由GORM自动管理）
  	UpdatedAt *time.Time // 最后一次更新时间（由GORM自动管理）
  }
  ```

  **约定**

  1. **主键**：GORM 使用一个名为`ID` 的字段作为每个模型的默认主键。
  2. **表名**：默认情况下，GORM 将结构体名称转换为 `snake_case` 并为表名加上复数形式。 例如，一个 `User` 结构体在数据库中的表名变为 `users` 。
  3. **列名**：GORM 自动将结构体字段名称转换为 `snake_case` 作为数据库中的列名。
  4. **时间戳字段**：GORM使用字段 `CreatedAt` 和 `UpdatedAt` 来自动跟踪记录的创建和更新时间。

  | 标签名                 | 说明                                                         |
  | :--------------------- | :----------------------------------------------------------- |
  | column                 | 指定 db 列名                                                 |
  | type                   | 列数据类型，推荐使用兼容性好的通用类型，例如：所有数据库都支持 bool、int、uint、float、string、time、bytes 并且可以和其他标签一起使用，例如：`not null`、`size`, `autoIncrement`… 像 `varbinary(8)` 这样指定数据库数据类型也是支持的。在使用指定数据库数据类型时，它需要是完整的数据库数据类型，如：`MEDIUMINT UNSIGNED not NULL AUTO_INCREMENT` |
  | serializer             | 指定将数据序列化或反序列化到数据库中的序列化器, 例如: `serializer:json/gob/unixtime` |
  | size                   | 定义列数据类型的大小或长度，例如 `size: 256`                 |
  | primaryKey             | 将列定义为主键                                               |
  | unique                 | 将列定义为唯一键                                             |
  | default                | 定义列的默认值                                               |
  | precision              | 指定列的精度                                                 |
  | scale                  | 指定列大小                                                   |
  | not null               | 指定列为 NOT NULL                                            |
  | autoIncrement          | 指定列为自动增长                                             |
  | autoIncrementIncrement | 自动步长，控制连续记录之间的间隔                             |
  | embedded               | 嵌套字段                                                     |
  | embeddedPrefix         | 嵌入字段的列名前缀                                           |
  | autoCreateTime         | 创建时追踪当前时间，对于 `int` 字段，它会追踪时间戳秒数，您可以使用 `nano`/`milli` 来追踪纳秒、毫秒时间戳，例如：`autoCreateTime:nano` |
  | autoUpdateTime         | 创建/更新时追踪当前时间，对于 `int` 字段，它会追踪时间戳秒数，您可以使用 `nano`/`milli` 来追踪纳秒、毫秒时间戳，例如：`autoUpdateTime:milli` |
  | index                  | 根据参数创建索引，多个字段使用相同的名称则创建复合索引，查看 [索引](https://gorm.io/zh_CN/docs/indexes.html) 获取详情 |
  | uniqueIndex            | 与 `index` 相同，但创建的是唯一索引                          |
  | check                  | 创建检查约束，例如 `check:age > 13`，查看 [约束](https://gorm.io/zh_CN/docs/constraints.html) 获取详情 |
  | <-                     | 设置字段写入的权限， `<-:create` 只创建、`<-:update` 只更新、`<-:false` 无写入权限、`<-` 创建和更新权限 |
  | ->                     | 设置字段读的权限，`->:false` 无读权限                        |
  | -                      | 忽略该字段，`-` 表示无读写，`-:migration` 表示无迁移权限，`-:all` 表示无读写迁移权限 |
  | comment                | 迁移时为字段添加注释                                         |

---

### 5. **GORM 高级功能**

##### **自动迁移**：使用 `AutoMigrate` 生成和更新数据库表

`AutoMigrate` 方法用于自动创建或更新数据库中的表结构，使其与模型结构保持同步。它会根据模型定义的字段和关系生成表格、添加或修改列，但不会删除已存在的列或更改其数据类型。`AutoMigrate` 可以帮助确保数据库结构与代码中的模型保持一致。

将上面使用到的User模型，增加一个手机号码字段。然后调用AutoMigrate方法进行表结构更新。

```go
type User struct {
	ID        uint       `gorm:"type:int;comment:主键"`   // Standard field for the primary key
	Name      string     `gorm:"size:128;comment:人员姓名"` // 一个常规字符串字段
	Email     *string    `gorm:"size:128;comment:邮箱地址"` // 一个指向字符串的指针, allowing for null values
	Age       uint8      // 一个未签名的8位整数
	Phone     string     `gorm:"size:11;comment:手机号码"` // 手机号码
	Birthday  *time.Time // A pointer to time.Time, can be null
	CreatedAt *time.Time // 创建时间（由GORM自动管理）
	UpdatedAt *time.Time // 最后一次更新时间（由GORM自动管理）
}
```

示例代码：

```go
func (userdao *UserDao) AutoMigrate() error {
	return userdao.Db.AutoMigrate(&model.User{})
}
```

调用上面的函数能够完成自动给表结构增加一个Phone字段。运行后会有类似下面的sql输出

```sql
ALTER TABLE `users` ADD `phone` varchar(11) COMMENT '手机号码'
```

##### **事务处理**：GORM 事务的使用，回滚与提交操作

自动事务：

```go
func (userdao *UserDao) Transaction(f func(tx *gorm.DB, parm interface{}) error, parm interface{}) error {
	return userdao.Db.Transaction(func(tx *gorm.DB) error {
		return f(tx, parm)
	})
}
```

上面的代码中，如果函数f的返回值 error不为nil，那么对数据库进行的操作将会被回滚。

手动事务：

```go

func (userdao *UserDao) SaveWithTransactionByManual(user *model.User, rollback bool) error {
	db := userdao.Db
	//开启事务
	db.Begin()
	db.Save(user)
	if rollback {
		// 模拟出现异常,进行数据回滚
		db.Rollback()
		return errors.New("发生异常，将数据进行回滚!")
	}

	// 成功执行，提交事务
	db.Commit()
	return nil
}
```

上面的代码演示在roolback为真的时候，会对相应的操作进行回滚操作。否则进行提交。

##### context

ORM 的上下文支持由 `WithContext` 方法启用，是一项强大的功能，可以增强 Go 应用程序中数据库操作的灵活性和控制力。 它允许在不同的操作模式、超时设置以及甚至集成到钩子/回调和中间件中进行上下文管理。

* 单会话模式
* 连续会话模式
* 

##### **钩子（Hooks）**：Before/After 钩子的使用场景和实现方式

* 创建的时候的钩子

  ```txt
  // 开始事务
  BeforeSave
  BeforeCreate
  // 关联前的 save
  // 插入记录至 db
  // 关联后的 save
  AfterCreate
  AfterSave
  // 提交或回滚事务
  
  ```

  ```go
  func (u *User) BeforeSave(tx *gorm.DB) error{
  	log.Println("BeforeSave ...")
  	return nil
  } 
  
  func (u *User) AfterSave(tx *gorm.DB) error{
  	log.Println("AfterSave ...")
  	return nil
  }
  
  
  func (u *User) BeforeCreate(tx *gorm.DB) error{
  	log.Println("AfterCreate ...")
  	return nil
  }
  func (u *User) AfterCreate(tx *gorm.DB) error{
  	log.Println("AfterCreate ...")
  	return nil
  }
  ```

  

* 更新时候的钩子

* 查询时候的钩子

* 删除时候的钩子

##### **多数据库** ：Database Resolver

##### **读写分离**: Database Resolver

##### **自定义插件**

##### **自定义 Logger**

---

### 6. **性能优化**

- **日志与调试**：使用 `db.Debug()` 查看 SQL 日志，调试查询。
- **批量操作**：批量插入、更新与删除。
- **连接池配置**：优化数据库连接池的使用。

---

### 7. **实战项目**

- **用户管理系统开发**：通过 GORM 实现用户增删改查功能。
- **API 集成**：使用 Gin 框架集成 GORM，开发 RESTful API。



### 8. **GORM Gen 自动代码生成**

#### 8.1 **GORM Gen 简介**

- **GORM 与 GORM Gen 的区别**：手写与自动生成代码的对比。
- GORM Gen 的优势
    - 自动生成代码，减少重复劳动。
    - 提供类型安全的查询方法。
- **适用场景**：适合中大型项目，有大量数据库表和复杂查询操作的应用。

#### 8.2 **环境准备**

- 安装依赖
    - 安装 GORM 和 GORM Gen。
    - 依赖工具：Go 和数据库驱动（MySQL/PostgreSQL）。
- 配置 Go 项目
    - 初始化 Go 模块，项目结构介绍。
    - GORM Gen 的基本配置。
- 数据库准备
    - 创建或导入一个简单的数据库供后续操作使用。

#### 8.3 **使用 GORM Gen 生成代码**

- 生成模型
    - 如何生成数据库模型代码。
    - 配置模型字段的类型、标签和关系。
    - `gorm.Model` 的使用与自定义模型。
- 生成 CRUD 代码
    - 自动生成增、删、改、查的代码。
    - 解析生成的代码结构。
- 示例
    - 基于简单的 `users` 表生成并操作模型和 CRUD 代码。

#### 8.4 **高级查询与自定义操作**

- 条件查询
    - 使用生成代码中的类型安全 API 进行复杂查询。
    - 条件查询、排序、分页查询的实现。
- 自定义 SQL 查询
    - 使用 GORM Gen 实现自定义 SQL 查询。
    - 如何编写复杂查询方法，并保持代码的类型安全性。

#### 8.5 **项目实战：人员管理系统**

- 项目需求
    - 开发一个人员管理系统，包括用户管理、部门管理、企业管理，员工管理等功能。
- GORM Gen 使用
    - 利用 GORM Gen 生成用户、部门、企业、员工的数据库模型和 CRUD 操作。
- 系统集成
    - 结合 Gin 框架，开发 RESTful API。


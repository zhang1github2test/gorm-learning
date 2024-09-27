# gorm实战学习

## 一、gorm入门

### 1、安装gorm

这里我们使用GoLand编辑器，然后golang的版本为1.22.0，项目的依赖通过Go mod来进行管理。

```shell
 go get -u gorm.io/gorm@v1.25.8
 go get -u gorm.io/driver/mysql@v1.5.5
```

### 2、测试gorm链接mysql数据库

为了链接mysql数据库，通常需要导入mysql驱动以及gorm依赖，示例如下：

```go
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
```

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

连接地址的举例： root:123456@tcp(192.168.188.155:3306)/szkfpt?charset=utf8mb4&parseTime=True&loc=Local

完整代码如下：

```go
package chapter01

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
  *  演示通过gorm来链接mysql库
     前提条件：mysql数据库已经安装好了
*/

func GetMysqlDb(account string, password string, host string, port int, dbname string) (*gorm.DB, error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// 链接格式： [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", account, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
```

测试代码如下：

```go
package test

import (
	"fmt"
	"go-orm-learn/chapter01"
	"testing"
)

func TestMysqlConnection(t *testing.T) {
	db, err := chapter01.GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {
		t.Errorf("err is not nil")
	}
	sqldb, _ := db.DB()

	fmt.Println("已经建立的链接数", sqldb.Stats().OpenConnections)
}
```

运行测试代码，结果如下：

```txt
=== RUN   TestMysqlConnection
root:123456@tcp(192.168.188.155:3306)/szkfpt?charset=utf8mb4&parseTime=True&loc=Local
已经建立的链接数 1
--- PASS: TestMysqlConnection (0.01s)
PASS
```

从上面的测试结果来看，目前已经能够正确链接到mysql数据库中了！

### 3、插入数据

##### 准备数据模型

```go
// chapter01/gorm_user.go
package chapter01

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           // Standard field for the primary key
	Name         string         // A regular string field
	Email        *string        // A pointer to a string, allowing for null values
	Age          uint8          // An unsigned 8-bit integer
	Birthday     *time.Time     // A pointer to time.Time, can be null
	MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
	ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
	CreatedAt    time.Time      // Automatically managed by GORM for creation time
	UpdatedAt    time.Time      // Automatically managed by GORM for update time
}
```



##### 执行插入操作

```go
package chapter01

import (
	"fmt"
	"time"
)

func CreateSingle() {
	db, err := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	if err != nil {

	}
	birthTime := time.Now()
	user := User{Name: "Jinzhu", Age: 18, Birthday: &birthTime}

	// 如果表不存在，则进行自动创建
	db.AutoMigrate(&User{})
	result := db.Create(&user) // 通过数据的指针来创建
	fmt.Println("影响的数据行数为:", result.RowsAffected)
}
```

测试代码：

```go
func TestCreate(t *testing.T) {
	chapter01.CreateSingle()
}
```



##### 插入的数据确认

运行的结果如下：

![image-20240321165520471](E:\go\go-orm-learn\gorm教程.assets\image-20240321165520471.png)

### 4、查询数据

##### 查询单个对象

GORM 提供了 `First`、`Take`、`Last` 方法，以便从数据库中检索单个对象。当查询数据库时它添加了 `LIMIT 1` 条件，且没有找到记录时，它会返回 `ErrRecordNotFound` 错误

###### first查询

按照主键升序或者第一个实体的第一个字段升序获取第一条记录

```go
	// 获取第一条记录（主键升序）
	var user User

	// 获取第一条记录（主键升序）
	// 相当于SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	result := db.First(&user)
	fmt.Println(user)
```

###### Take方法查询

获取第一条记录，没有进行排序

```go
	user = User{}
	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;
	fmt.Println(user)
```

###### Last方法查询

按照主键升序或者第一个实体的第一个字段降序获取第一条记录

```go
	user = User{}
	// 获取最后一条记录（主键降序）
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;
	db.Last(&user)
	user = User{}
```

##### 根据主键查询

```go
	var users []User
	var user2 User
	// 根据主键查询索引

	// SELECT * FROM users WHERE id = 1;
	db.First(&user2, 1)

	fmt.Println(result1)
    
    // 把user2恢复到默认状态
     user2 = User{}
	// SELECT * FROM users WHERE id = 1;
	db.First(&user2, "1")

	// SELECT * FROM users WHERE id IN (1,2,3);
	db.Find(&users, []int{1, 2, 3})
```

运行后的输出结果为：

```txt
2024/03/25 13:52:50 E:/go/go-orm-learn/chapter01/grom_crud.go:99
[10.932ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
&{0xc000188510 <nil> 1 0xc000084c40 0}

2024/03/25 13:52:50 E:/go/go-orm-learn/chapter01/grom_crud.go:105
[1.030ms] [rows:1] SELECT * FROM `users` WHERE `users`.`id` = '1' ORDER BY `users`.`id` LIMIT 1

2024/03/25 13:52:50 E:/go/go-orm-learn/chapter01/grom_crud.go:108
[0.517ms] [rows:3] SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
--- PASS: TestSelectById (0.02s)
```

在目标结构有一个主键值时，可以使用主键构建查询条件，如下面的所示：

```go
	var user = User{ID: 1}
	// SELECT * FROM `users` WHERE `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	db.First(&user)
```

##### 查询全部对象

```go
func SelectAll() {
	var users []User
	// 这里需要替换成自己的mysql地址及账号等
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// SELECT * FROM `users`
	db.Find(&users)
	fmt.Println(users)
}
```

##### 基于条件查询

###### string条件查询

​	等于查询

```go
// string条件查询
func SelectByConditionString() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	var user User
	// SELECT * FROM `users` WHERE name = 'jinzhu' AND age = 18
	db = db.Where("name = ?", "jinzhu")
	db.Where("age = ?", 18).Find(&user)
    
	fmt.Println(user)
}
```

不等于查询

```go
	var users []User
	// SELECT * FROM `users` WHERE name <> 'jinzhu'
	// 不等于查询
	db.Where("name <> ?", "jinzhu").Find(&users)
	fmt.Println(users)
```

in查询

```go
// SELECT * FROM `users` WHERE name IN ('jinzhu','jinzhu2')
	// in查询
	db.Where("name IN ?", []string{"jinzhu", "jinzhu2"}).Find(&users)
	fmt.Println(users)
```

like模糊查询

```go
// SELECT * FROM `users` WHERE name like '%jinzhu%'
	// like模糊查询
	db.Where("name like ?", "%jinzhu%").Find(&users)
	fmt.Println(users)
```

and查询

```go
	// SELECT * FROM `users` WHERE name = 'jinzhu' and age >= 10
	// and查询
	db.Where("name = ? and age >= ?", "jinzhu", 10).Find(&users)
	fmt.Println(users)
```



###### Struct条件和map条件查询

struct查询

```go
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// Struct
	// SELECT * FROM `users` WHERE `users`.`name` = 'jinzhu' AND `users`.`age` = 18 ORDER BY `users`.`id` LIMIT 1;
	db.Where(&User{Name: "jinzhu", Age: 18}).First(&user)
```

map查询

```go
	// Map
	// SELECT * FROM `users` WHERE `age` = 18 AND `name` = 'jinzhu'
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)
```

通过主键id进行查询

```go
	// Slice of primary keys
	// SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
	db.Where([]int64{1, 2, 3}).Find(&users)
```

在通过结构体进行查询的时候，如果是零值数据，那么不会构建查询条件。

如下所示：

```go
db.Where(&User{Name: "jinzhu", Age: 0}).Find(&users)
```

由于age字段的值为0，所以不会基于age字段来构建查询条件，最终的查询语句则为`SELECT * FROM users WHERE name = "jinzhu";`

当需要零值进行查询的时候，那么需要使用map结构来进行查询

```go
	// SELECT * FROM `users` WHERE `Age` = 0 AND `Name` = 'jinzhu'
	db.Where(map[string]interface{}{"Name": "jinzhu", "Age": 0}).Find(&users)
```

指定查询字段来规避零值问题

```go
// SELECT * FROM users WHERE name = "jinzhu" AND age = 0;
db.Where(&User{Name: "jinzhu"}, "name", "Age").Find(&users)

// SELECT * FROM users WHERE age = 0;
db.Where(&User{Name: "jinzhu"}, "Age").Find(&users)
```

###### 内联条件

除了使用where条件外，我们可以在使用可以使用内联条件方式在Find、First、Last等方法中使用，下面是具体示例：

```go
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// Get by primary key if it were a non-integer type
	// SELECT * FROM users WHERE id = 1;
	db.First(&user, "id = ?", 1)

	// Plain SQL
	db.Find(&user, "name = ?", "jinzhu")
	// SELECT * FROM users WHERE name = "jinzhu";

	db.Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age > 20;

	// Struct
	db.Find(&users, User{Age: 20})
	// SELECT * FROM users WHERE age = 20;

	// Map
	db.Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM users WHERE age = 20;
```

###### Not 条件

not条件跟where有点相似

```go
	var user User
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Not("name = ?", "jinzhu").First(&user)
	// SELECT * FROM users WHERE NOT name = "jinzhu" ORDER BY id LIMIT 1;

	// Not In
	db.Not(map[string]interface{}{"name": []string{"jinzhu", "jinzhu 2"}}).Find(&users)
	// SELECT * FROM users WHERE name NOT IN ("jinzhu", "jinzhu 2");

	// Struct
	db.Not(User{Name: "jinzhu", Age: 18}).First(&user)
	// SELECT * FROM users WHERE name <> "jinzhu" AND age <> 18 ORDER BY id LIMIT 1;

	// Not In slice of primary keys
	db.Not([]int64{1, 2, 3}).First(&user)
	// SELECT * FROM users WHERE id NOT IN (1,2,3) ORDER BY id LIMIT 1;
```

###### Or 条件

```go
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	
	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (name = 'jinzhu 2' and age = 18)
	db.Where("name = ?", "jinzhu").Or("name = ? and age = ?", "jinzhu 2", 18).Find(&users)

	// Struct
	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (`users`.`name` = 'jinzhu 2' AND `users`.`age` = 18)
	db.Where("name = 'jinzhu'").Or(User{Name: "jinzhu 2", Age: 18}).Find(&users)

	// Map
	// SELECT * FROM `users` WHERE name = 'jinzhu' OR (`age` = 18 AND `name` = 'jinzhu 2')
	db.Where("name = 'jinzhu'").Or(map[string]interface{}{"name": "jinzhu 2", "age": 18}).Find(&users)
```

###### 查询指定字段

通过Select方法来指定需要查询的字段，如果没有指定，则默认查询所有的字段。

```go
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	// SELECT `name`,`age` FROM `users`
	db.Select("name", "age").Find(&users)
```

###### 排序

通过order方法指定排序的字段及排序方式

```go
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Order("age desc, name").Find(&users)
	// SELECT * FROM users ORDER BY age desc, name;

	// Multiple orders
	// SELECT * FROM users ORDER BY age desc, name;
	db.Order("age desc").Order("name").Find(&users)
```

###### Limit & Offset

```go
	var users []User
	var users1 []User
	var users2 []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Limit(3).Find(&users)
	// SELECT * FROM users LIMIT 3;

	// Cancel limit condition with -1
	db.Limit(10).Find(&users1).Limit(-1).Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	// SELECT * FROM `users` LIMIT 10 OFFSET 5
	// limit表示获取多少条数据  OFFSET表示跳过多少条数据
	db.Offset(1).Limit(10).Find(&users)
```

###### Group By & Having

```go
	var result result
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	db.Model(&User{}).Select("name, sum(age) as total").Where("name LIKE ?", "group%").Group("name").Order("name desc").Find(&result)

	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "group").Find(&result)
```

###### Distinct

```go
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	var result map[string]interface{}
	db.Model(&User{}).Distinct("name", "age").Find(&result)
	fmt.Println(result)
```

### 4、修改数据

#### 保存所有字段

可以使用save字段来更新所有字段，及时字段是零值

```
	var user User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.First(&user)

	user.Name = "jinzhu 2"
	user.Age = 100
	// UPDATE `users` SET `name`='jinzhu 2',`email`=NULL,
	//`age`=100,`birthday`='2024-03-21 16:52:50.665',`member_number`=NULL,
	//`activated_at`=NULL,`created_at`='2024-03-21 16:52:50.69',
	//`updated_at`='2024-03-25 16:25:08.448' WHERE `id` = 1
	db.Save(&user)
```

从上面输出的sql语句可以看到，User结构体中的所有字段都被更新了，及时字段为空也会被更新。

save方法是一个组合方法，如果保存的值不存在主键，就会进行新建！

如果存在主键，那么就会进行更新。

```go
	// INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) 
	// VALUES ('jinzhu8',NULL,22,NULL,NULL,NULL,'2024-03-25 16:38:58.857','2024-03-25 16:38:58.857')
	db.Save(&User{
		Name:      "jinzhu8",
		Age:       22,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	// UPDATE `users` SET `name`='jinzhu',`email`=NULL,`age`=100,`birthday`=NULL,`member_number`=NULL,
	//`activated_at`=NULL,`created_at`='2024-03-25 16:38:58.859',`updated_at`='2024-03-25 16:38:58.86' WHERE `id` = 1
	db.Save(&User{ID: 1, Name: "jinzhu", Age: 100, CreatedAt: time.Now()})
```

上面的第一次的Save方法调用，由于主键Id没有赋值，所以执行了保存操作。第二次调用，由于指定ID为1，那么执行的是一个更新操作。

#### 更新单个列

当使用 `Update` 更新单列时，需要有一些条件，否则将会引起`ErrMissingWhereClause` 错误，查看 [阻止全局更新](https://gorm.io/zh_CN/docs/update.html#block_global_updates) 了解详情。 当使用 `Model` 方法，并且它有主键值时，主键将会被用于构建条件，例如：

```go
	var user User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// 根据条件更新
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.216' WHERE age = 18;
	db.Model(&User{}).Where("age = ?", 18).Update("name", "hello")

	// User 的 ID 是 1
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.227' WHERE `id` = 1
	db.Model(&User{ID: 1}).Update("name", "hello")


	// 根据条件和 model 的值进行更新
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 16:57:52.23' WHERE age = 100
	db.Model(&user).Where("age = ?", 100).Update("name", "hello")
```

#### 更新多列

gorm的Updates方法支持更新struct和map[string]interface{}参数，当参数是在结构体的时候，默认情况下只会更新非零值的字段

```go
	var user = User{
		ID: 1,
	}
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// 根据 `struct` 更新属性，只会更新非零值的字段
	// UPDATE `users` SET `name`='hello',`updated_at`='2024-03-25 17:14:10.48' WHERE `id` = 1
	db.Model(&user).Updates(User{Name: "hello", Age: 0})

	// 根据 `map` 更新属性
	// UPDATE `users` SET `age`=0,`name`='hello',`updated_at`='2024-03-25 17:14:10.494' WHERE `id` = 1
	db.Model(&user).Updates(map[string]interface{}{"name": "hello", "age": 0})
```

#### 更新选定字段

```
// UPDATE `users` SET `name`='hello2',`updated_at`='2024-03-25 17:33:01.714' WHERE `id` = 1
	db.Model(&User{ID: 1}).Select("name").Updates(map[string]interface{}{"name": "hello2", "age": 18, "active": false})
```

可以通过Select方法来选定特定字段来进行更新

#### 批量更新

```go
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	
  // UPDATE `users` SET `name`='hello',`age`=18,`updated_at`='2024-03-25 17:41:13.394' WHERE name = 'hello2'
	db.Model(User{}).Where("name = ?", "hello2").Updates(User{
		Name: "hello",
		Age:  18,
	})
```

### 5、删除数据

#### 删除一条记录

删除一条记录的时候，需要指定对应的主键，不然会发生批量删除。

```sql
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// 删除Id = 1的数据
	db.Delete(&User{
		ID: 1,
	})
```

测试代码如下：

```go
import (
	"go-orm-learn/chapter01"
	"testing"
)

func TestDeleteSingle(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chapter01.DeleteSingle()
		})
	}
}
```

观测数据库的结果：

![image-20240326094835718](E:\go\go-orm-learn\gorm教程.assets\image-20240326094835718.png)

说明数据已经被删除了！



#### 通过主键删除记录

GORM 允许通过主键(可以是复合主键)和内联条件来删除对象

```go
// 通过主键删除数据
func DeleteById() {
	var users []User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, 10)
	
	// DELETE FROM users WHERE id = 10;
	db.Delete(&User{}, "10")
	
	
	// DELETE FROM `users` WHERE `users`.`id` IN (1,2,3)
	db.Delete(&users, []int{1, 2, 3})
}
```

测试代码如下：

```go
func TestDeleteById(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "DeleteById"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chapter01.DeleteById()
		})
	}
}
```

运行测试代码后，可以发现数据库中对应的数据已经被成功删除

#### 带条件删除

如果指定的值不包括主属性，那么 GORM 会执行批量删除，它将删除所有匹配的记录。

```go
// 按照条件删除对应的
func DeleteBatch() {
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")

	// DELETE FROM `users` WHERE name = 'jinzhu'
	db.Delete(&User{}, "name = ?", "jinzhu")
}
```

测试代码如下：

```sql
func TestDeleteByBatch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "DeleteByBatch"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chapter01.DeleteBatch()
		})
	}
}
```

运行测试代码后，发现所有的姓名为jinzhu记录都已经被删除成功删除了

### 6、使用原生sql操作数据

尽管gorm框架已经提供很多强大的功能，但是在一些特殊的场景下，我们希望直接使用sql语句来跟数据库进行交互，此时我们可以按照下面的步骤进行：

```go
// 通过原生sql操作数据
func SelectBySql() {
	var user User
	db, _ := GetMysqlDb("root", "123456", "192.168.188.155", 3306, "szkfpt")
	db.Raw("select * from users;").Scan(&user)
	fmt.Println(user)
}
```

测试代码：

```go

func Test_selectBySql(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "selectBySql",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chapter01.SelectBySql()
		})
	}
}
```

测试输出的结果为：

```go
[11.968ms] [rows:1] select * from users;
{1 Jinzhu <nil> 18 2024-03-27 16:38:01.052 +0800 CST { false} {0001-01-01 00:00:00 +0000 UTC false} 2024-03-27 16:38:01.104 +0800 CST 2024-03-27 16:38:01.104 +0800 CST}
--- PASS: Test_selectBySql (0.02s)
    --- PASS: Test_selectBySql/selectBySql (0.02s)
```


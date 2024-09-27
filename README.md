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

- **数据库连接**：
    
    - 使用 GORM 连接数据库，演示 DSN 字符串的配置。
    
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
    
      
- **模型定义**：
    
    - 如何定义模型及其与数据库表的映射（字段类型、标签）
    
      ```go
      type User struct {
        ID           uint           // Standard field for the primary key
        Name         string         // 一个常规字符串字段
        Email        *string        // 一个指向字符串的指针, allowing for null values
        Age          uint8          // 一个未签名的8位整数
        Birthday     *time.Time     // A pointer to time.Time, can be null
        MemberNumber sql.NullString // Uses sql.NullString to handle nullable strings
        ActivatedAt  sql.NullTime   // Uses sql.NullTime for nullable time fields
        CreatedAt    time.Time      // 创建时间（由GORM自动管理）
        UpdatedAt    time.Time      // 最后一次更新时间（由GORM自动管理）
      }
      ```
    
      ### 约定
    
      1. **主键**：GORM 使用一个名为`ID` 的字段作为每个模型的默认主键。
      2. **表名**：默认情况下，GORM 将结构体名称转换为 `snake_case` 并为表名加上复数形式。 例如，一个 `User` 结构体在数据库中的表名变为 `users` 。
      3. **列名**：GORM 自动将结构体字段名称转换为 `snake_case` 作为数据库中的列名。
      4. **时间戳字段**：GORM使用字段 `CreatedAt` 和 `UpdatedAt` 来自动跟踪记录的创建和更新时间。
- **CRUD 操作**：
    - **Create**：插入单条和多条记录。
    
      
    
    - **Read**：查询单条、多条记录，条件查询、分页、排序。
    
    - **Update**：更新单个字段和多个字段。
    
    - **Delete**：删除记录（软删除与硬删除）

* GORM配置：
  * 连接池配置
  * 



---

### 5. **GORM 高级功能**

- **自动迁移**：使用 `AutoMigrate` 生成和更新数据库表。
- **事务处理**：GORM 事务的使用，回滚与提交操作。
- **钩子（Hooks）**：Before/After 钩子的使用场景和实现方式。
- **多数据库** ：Database Resolver
- **读写分离**: Database Resolver
- **自定义插件**
- **自定义 Logger**
- 

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

#### 8.5 **项目实战：简易博客系统**

- 项目需求
    - 开发一个简易的博客系统，包括用户管理、文章管理、评论管理等功能。
- GORM Gen 使用
    - 利用 GORM Gen 生成用户、文章、评论的数据库模型和 CRUD 操作。
- 系统集成
    - 结合 Gin 框架，开发 RESTful API。


### 1. **创建本地目录**
   先在宿主机上创建两个目录，一个用于保存 MySQL 数据，一个用于保存 MySQL 配置文件：
   ```bash
   mkdir -p /mydata/mysql-data
   mkdir -p /mydata/mysql-config
   ```

   - `/mydata/mysql-data`：将映射为 MySQL 容器中的数据目录。
   - `/mydata/mysql-config`：将映射为 MySQL 容器中的配置文件目录。

### 2. **生成默认 MySQL 配置文件**
   如果需要自定义 MySQL 配置文件，可以从 MySQL 镜像中提取默认配置文件：
   1. **运行 MySQL 容器：**
      ```bash
      docker run --name temp-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:latest
      ```

   2. **将配置文件从容器拷贝到宿主机：**
      
      ```bash
      docker cp temp-mysql:/etc/mysql /mydata/mysql-config
      docker cp temp-mysql:/etc/my.cnf /mydata/mysql-config/mysql/conf.d
      ```

  

   3. **停止并删除临时容器：**

  ```bash
docker stop temp-mysql
docker rm temp-mysql
  ```

  现在，宿主机的 `/mydata/mysql-config` 目录中应该有 MySQL 的默认配置文件。

### 3. **运行 MySQL 容器并挂载卷**
   使用 `docker run` 命令来运行 MySQL 容器，同时将数据目录和配置文件映射到宿主机：

   ```bash
   docker run --name mysql-container \
     -e MYSQL_ROOT_PASSWORD=my-secret-pw \
     -v /mydata/mysql-data:/var/lib/mysql \
     -v /mydata/mysql-config/mysql:/etc/mysql \
     -p 3306:3306 \
     -d mysql:latest
   ```

   参数说明：
   - `-v /mydata/mysql-data:/var/lib/mysql`：将宿主机的 `/mydata/mysql-data` 映射为容器中的数据目录 `/var/lib/mysql`。
   - `-v /mydata/mysql-config/mysql:/etc/mysql`：将宿主机的 `/mydata/mysql-config/mysql` 目录映射为容器中的配置目录 `/etc/mysql`。
   - `-p 3306:3306`：映射 MySQL 的端口。

### 4. **验证数据持久化和配置文件**
   - **数据持久化**：停止并删除容器后，重新启动一个容器时，数据会保存在宿主机的 `/mydata/mysql-data` 中，可以验证数据持久化功能。
   - **配置文件修改**：你可以编辑宿主机上的 `/mydata/mysql-config/my.cnf` 文件，修改 MySQL 配置，重启容器使配置生效：
     ```bash
     docker restart mysql-container
     ```

通过这种方式，你的 MySQL 数据和配置文件都将保存到宿主机中，确保数据的持久性和配置的灵活性。
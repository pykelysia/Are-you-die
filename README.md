# Are You Die?

## 介绍
这是一个用来记录你连续打卡所用的idea，当你连续三天没有完成打卡任务的话，那么将联系你的紧急联系人来检查你的生命状态。

## 技术栈
前端：待定。

后端：Golang + Gin + Gorm。

数据库：MySQL。

## 功能
可参考 [介绍](#介绍) 部分。

## 快速开始
确保数据库可正常运行。

1. 克隆仓库：
   ```bash
   git clone 
    ```
2. 进入项目目录：
    ```bash
    cd Are-you-die
    ```
3. 配置数据库连接信息：
   - 创建 `config_mysql.json` 文件，内容如下：
     ```json
     {
       "username": "your_username",
       "password": "your_password",
       "database_host": "localhost",
       "database_port": 3306,
       "database_name": "are_you_die"
     }
     ```
   - 替换 `your_username` 和 `your_password` 为你的 MySQL 用户名和密码。
4. 启动服务器：
    ```bash
    go run main.go
    ```

## 贡献
欢迎任何形式的贡献！请提交 Pull Request 或打开 Issue。

## Created By Pyke elysia
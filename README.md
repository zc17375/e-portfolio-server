# e-portfolio-server

e-portfolio-server

# mysql database

```
    docker run -p 3306:3306 -d --name mysql -e MYSQL_ROOT_PASSWORD=root -e TZ=Asia/Taipei mysql:latest
```

# Install related Dependency packages

```
# HTTP WEB 框架
go get -u github.com/gin-gonic/gin

# DB ORM
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql

# 設置配置
go get github.com/spf13/viper

# LOG日誌配置
go get -u go.uber.org/zap
go get github.com/lestrrat-go/file-rotatelogs // 輪換日誌

# SWAGGER 文件配置
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

# UUID
go get github.com/google/uuid

# JWT
go get github.com/golang-jwt/jwt/v5

# MONGODB
go get go.mongodb.org/mongo-driver/mongo


```

docker related command
```
    # 建立docker build
    docker build -f Dockerfile -t go-app .

    # 運行建立的docker build
    docker run -p 8080:8888 --name my-go-app --link mysql:mysql go-app  # 連單個container
    docker run -p 8080:8888 --name my-go-app --network MONGO go-app # 加入到network

    # 查看所有images
    docker images

    # 顯示所有容器的清單，包括正在執行的和已停止的
    docker ps -a

    # 停止運行 docker container
    docker stop  <容器名稱或ID>

    # 刪除 docker container
    docker rm <容器名稱或ID>

    # 





```

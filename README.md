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




```
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

mongodb collection desing

```
{
  "_id": ObjectId("..."),          // MongoDB自動生成的唯一識別符
  "name": "姓名",
  "job_title": "職稱",
  "head_img_path": "頭像圖片路徑",
  "email": "電子郵件",
  "media": {
    "Github": "Github 連結",
    "Linkedin": "Linkedin 連結",
    "Facebook": "Facebook 連結"
  },
  "skills": ["JavaScript", "TypeScript", "React", "Vue", "..."],
  "projects": [
    {
      "name": "專案1名稱",
      "skills": ["專案1相關技能"],
      "introduce": "專案1介紹",
      "demo_link": "專案1演示連結",
      "github_rep_link": "專案1 Github 倉庫連結"
    },
    {
      "name": "專案2名稱",
      "skills": ["專案2相關技能"],
      "introduce": "專案2介紹",
      "demo_link": "專案2演示連結",
      "github_rep_link": "專案2 Github 倉庫連結"
    }
  ],
  "resume_link": "履歷連結"
}

```

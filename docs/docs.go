// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/auth/delete": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "註銷帳號",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/auth/login": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "會員登入",
                "parameters": [
                    {
                        "description": "信箱或手機號碼, 密碼",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "返回使用者資訊,token,過期時間",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.LoginResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/auth/refresh-token": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "刷新token",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/auth/register": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "會員註冊",
                "parameters": [
                    {
                        "description": "使用者名稱, 密碼, 暱稱, 頭貼網址, 手機號碼, 信箱",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Register"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "會員註冊,返回會員註冊訊息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/common.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/response.UserResponse"
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/individual/create": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Individual"
                ],
                "summary": "新增個人作品集資料",
                "parameters": [
                    {
                        "description": "個人作品集資料",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Individual"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/individual/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Individual"
                ],
                "summary": "刪除個人作品集資料",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/individual/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Individual"
                ],
                "summary": "更新個人作品集資料",
                "parameters": [
                    {
                        "description": "更新個人作品集資料",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Individual"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/portfolio/list": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "個人作品集列表",
                "parameters": [
                    {
                        "description": "取得分頁列表",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/common.Pagination"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/portfolio/{userName}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Portfolio"
                ],
                "summary": "個人作品集",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user Name",
                        "name": "userName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "common.Pagination": {
            "type": "object",
            "properties": {
                "currentPage": {
                    "description": "當前頁碼",
                    "type": "integer",
                    "example": 1
                },
                "keywords": {
                    "description": "關鍵字",
                    "type": "string",
                    "example": "software"
                },
                "pageSize": {
                    "description": "每頁大小",
                    "type": "integer",
                    "example": 10
                },
                "sortBy": {
                    "type": "string",
                    "example": "name"
                },
                "totalPages": {
                    "description": "總頁數",
                    "type": "integer"
                },
                "totalRecords": {
                    "description": "總記錄數",
                    "type": "integer"
                }
            }
        },
        "common.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "model.Individual": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "john@example.com"
                },
                "head_img_path": {
                    "type": "string",
                    "example": "/images/john_doe.jpg"
                },
                "job_title": {
                    "type": "string",
                    "example": "Software Engineer"
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "projects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Project"
                    }
                },
                "resume_link": {
                    "type": "string",
                    "example": "https://example.com/john_doe_resume.pdf"
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "Golang",
                        "JavaScript",
                        "Python"
                    ]
                },
                "social_media": {
                    "$ref": "#/definitions/model.SocialMedia"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "model.Project": {
            "type": "object",
            "properties": {
                "demo_link": {
                    "type": "string",
                    "example": "https://example.com/project_x_demo"
                },
                "github_rep_link": {
                    "type": "string",
                    "example": "https://github.com/johndoe/project_x"
                },
                "introduce": {
                    "type": "string",
                    "example": "A project for implementing microservices architecture."
                },
                "name": {
                    "type": "string",
                    "example": "Project X"
                },
                "skills": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "['Golang'",
                        " 'Docker'",
                        " 'Kubernetes']"
                    ]
                }
            }
        },
        "model.SocialMedia": {
            "type": "object",
            "properties": {
                "facebook": {
                    "type": "string",
                    "example": "johndoe"
                },
                "github": {
                    "type": "string",
                    "example": "johndoe"
                },
                "linkedin": {
                    "type": "string",
                    "example": "johndoe"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "description": "創建時間",
                    "type": "string"
                },
                "disable": {
                    "description": "帳號是否註銷",
                    "type": "integer"
                },
                "email": {
                    "description": "電子信箱",
                    "type": "string"
                },
                "headerImg": {
                    "description": "會員頭像連結",
                    "type": "string"
                },
                "id": {
                    "description": "主鍵ID",
                    "type": "integer"
                },
                "last_login_at": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "phone": {
                    "description": "手機號碼",
                    "type": "string"
                },
                "updatedAt": {
                    "description": "更新時間",
                    "type": "string"
                },
                "userName": {
                    "description": "會員名稱",
                    "type": "string"
                },
                "uuid": {
                    "description": "會員UUID",
                    "type": "string"
                }
            }
        },
        "request.Login": {
            "type": "object",
            "required": [
                "account",
                "password"
            ],
            "properties": {
                "account": {
                    "description": "帳號",
                    "type": "string",
                    "example": "0912345678"
                },
                "password": {
                    "description": "密碼",
                    "type": "string",
                    "example": "123456"
                }
            }
        },
        "request.Register": {
            "type": "object",
            "required": [
                "email",
                "passWord",
                "phone",
                "userName"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "123@gmail.com"
                },
                "headerImg": {
                    "type": "string",
                    "example": "http://headimgurl.com"
                },
                "nickName": {
                    "type": "string",
                    "example": "Ed"
                },
                "passWord": {
                    "type": "string",
                    "example": "123456"
                },
                "phone": {
                    "type": "string",
                    "example": "0912345678"
                },
                "userName": {
                    "type": "string",
                    "example": "Edward"
                }
            }
        },
        "response.LoginResponse": {
            "type": "object",
            "properties": {
                "expiresAt": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        },
        "response.UserResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/model.User"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "x-token",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1.0.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "E-Portfolio Swagger API接口文件",
	Description:      "使用golang gin開發全端個人作品集網站",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

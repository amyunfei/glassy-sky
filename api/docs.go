// Code generated by swaggo/swag. DO NOT EDIT
package api

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
        "/category": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "分类信息"
                ],
                "summary": "分页查询分类",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "分类信息"
                ],
                "summary": "创建分类",
                "parameters": [
                    {
                        "description": "分类信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryResponse"
                        }
                    }
                }
            }
        },
        "/category/{id}": {
            "put": {
                "tags": [
                    "分类信息"
                ],
                "summary": "修改分类",
                "parameters": [
                    {
                        "description": "分类信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ModifyCategoryRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "分类信息"
                ],
                "summary": "删除分类",
                "parameters": [
                    {
                        "type": "string",
                        "description": "分类id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessEmptyResponse"
                        }
                    }
                }
            }
        },
        "/label": {
            "get": {
                "tags": [
                    "标签信息"
                ],
                "summary": "分页查询标签",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryResponse"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "标签信息"
                ],
                "summary": "创建标签",
                "parameters": [
                    {
                        "description": "标签信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateLabelResponse"
                        }
                    }
                }
            }
        },
        "/label/{id}": {
            "put": {
                "tags": [
                    "标签信息"
                ],
                "summary": "修改标签",
                "parameters": [
                    {
                        "description": "标签信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ModifyLabelRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateLabelResponse"
                        }
                    }
                }
            },
            "delete": {
                "tags": [
                    "标签信息"
                ],
                "summary": "删除标签",
                "parameters": [
                    {
                        "type": "string",
                        "description": "标签id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessEmptyResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "tags": [
                    "用户信息"
                ],
                "summary": "分页查询用户",
                "parameters": [
                    {
                        "type": "integer",
                        "name": "current",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "size",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "nickname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateCategoryResponse"
                        }
                    }
                }
            }
        },
        "/user/email-code/{email}": {
            "get": {
                "tags": [
                    "用户信息"
                ],
                "summary": "发送邮箱验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱地址",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessEmptyResponse"
                        }
                    }
                }
            }
        },
        "/user/email-verify/{email}": {
            "get": {
                "tags": [
                    "用户信息"
                ],
                "summary": "验证邮箱可用",
                "parameters": [
                    {
                        "type": "string",
                        "description": "邮箱地址",
                        "name": "email",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SuccessEmptyResponse"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "tags": [
                    "用户信息"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "description": "登录信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/user/register": {
            "post": {
                "tags": [
                    "用户信息"
                ],
                "summary": "注册用户",
                "parameters": [
                    {
                        "description": "用户注册信息",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "put": {
                "tags": [
                    "用户信息"
                ],
                "summary": "修改用户信息",
                "parameters": [
                    {
                        "description": "用户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.ModifyUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateCategoryRequest": {
            "type": "object",
            "required": [
                "color",
                "name"
            ],
            "properties": {
                "color": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                }
            }
        },
        "dto.CreateCategoryResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dto.CreateLabelRequest": {
            "type": "object",
            "required": [
                "color",
                "name"
            ],
            "properties": {
                "color": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.CreateLabelResponse": {
            "type": "object",
            "properties": {
                "color": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserRequest": {
            "type": "object",
            "required": [
                "code",
                "email",
                "password",
                "username"
            ],
            "properties": {
                "code": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.CreateUserResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "dto.ModifyCategoryRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                }
            }
        },
        "dto.ModifyLabelRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "color": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dto.ModifyUserRequest": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.SuccessEmptyResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "string",
                    "example": ""
                },
                "msg": {
                    "type": "string",
                    "example": "success"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Glassy Sky API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

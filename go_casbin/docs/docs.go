// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/alive": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "鉴权"
                ],
                "summary": "检查token是否过期",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/auth": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "鉴权"
                ],
                "summary": "登录获取登录token 信息",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/authority/add": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Authority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/authority/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "删除角色",
                "parameters": [
                    {
                        "description": "删除角色",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Authority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/set": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "设置角色资源权限",
                "parameters": [
                    {
                        "description": "设置角色资源权限",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Authority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/authority/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色"
                ],
                "summary": "更新角色信息",
                "parameters": [
                    {
                        "description": "权限id, 权限名, 父角色id",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Authority"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"code\":200,\"data\":{},\"msg\":\"ok\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/policy": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "访问权限"
                ],
                "summary": "获取权限列表",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "访问权限"
                ],
                "summary": "增加访问权限",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserPolicy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "访问权限"
                ],
                "summary": "删除访问权限",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserPolicy"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取所有用户",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "增加用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/user/delete": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "刪除用户",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/user/get": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuthc": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取单一用户",
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        },
        "/user/update": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{ \"code\": 200, \"data\": {}, \"msg\": \"ok\" }",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/utils.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Authority": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "string"
                },
                "authorityName": {
                    "type": "string"
                },
                "parentId": {
                    "type": "string"
                }
            }
        },
        "models.UserLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "description": "10003 账号不存在\n20001 30002 登录异常\n400 参数错误\n40002 用户不存在\n40003 账号密码错误",
                    "type": "string"
                }
            }
        },
        "models.UserPolicy": {
            "type": "object",
            "properties": {
                "type": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.UserRegister": {
            "type": "object",
            "properties": {
                "company_id": {
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
        "utils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "错误码",
                    "type": "integer"
                },
                "data": {
                    "description": "详细错误信息",
                    "type": "string"
                },
                "msg": {
                    "description": "错误信息",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.0.1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "code server",
	Description: "Go 学习综合demo",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
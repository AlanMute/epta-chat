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
        "/chat": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создать чат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Создать чат",
                "parameters": [
                    {
                        "description": "Данные для создания чата",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.AddChat"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Чат создан"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chat/add/members": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Добавить участника",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Добавить участника",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID пользователя",
                        "name": "user-id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "description": "Список users_id",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.AddMember"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Чат создан"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chat/all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить список чатов пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Получить список чатов пользователя",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/core.Chat"
                            }
                        }
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chat/members/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить список участников чата",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Получить список участников чата",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID чата",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/core.UserInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/chat/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить чат по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Получить чат по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID чата",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Chat"
                        }
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удалить чат",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Chat"
                ],
                "summary": "Удалить чат",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID чата",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Чат удален"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/contact": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Создать контакт",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact"
                ],
                "summary": "Создать контакт",
                "responses": {
                    "201": {
                        "description": "Контакт создан"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/contact/all": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить список контактов пользователя",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact"
                ],
                "summary": "Получить список контактов пользователя",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/core.UserInfo"
                            }
                        }
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/contact/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Получить контакт по ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact"
                ],
                "summary": "Получить контакт по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID контакта",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.UserInfo"
                        }
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Удалить контакт",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Contact"
                ],
                "summary": "Удалить контакт",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID контакта",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Контакт удален"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/messenger/connect": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Установить websocket соединение с мессенджером",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Messenger"
                ],
                "summary": "Подключиться к мессенджеру",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID чата подключения",
                        "name": "chat-id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла непредвиденная ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/refresh": {
            "post": {
                "description": "Обновить токены",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Обновить токены",
                "parameters": [
                    {
                        "description": "Данные для регистрации",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Refresh"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Токены обновлены"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/sign-in": {
            "post": {
                "description": "Войти",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Войти",
                "parameters": [
                    {
                        "description": "Данные для регистрации",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Sign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Вход выполнен"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/sign-up": {
            "post": {
                "description": "Зарегистрироваться",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Зарегистрироваться",
                "parameters": [
                    {
                        "description": "Данные для регистрации",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.Sign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Регистрация выполнена"
                    },
                    "400": {
                        "description": "Запрос не правильно составлен",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Возникла внутренняя ошибка",
                        "schema": {
                            "$ref": "#/definitions/resp.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Chat": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "isDirect": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "owner": {
                    "$ref": "#/definitions/core.User"
                },
                "ownerId": {
                    "type": "integer"
                }
            }
        },
        "core.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "core.UserInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "login": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "handler.AddChat": {
            "type": "object",
            "properties": {
                "is_direct": {
                    "type": "boolean"
                },
                "members_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "handler.AddMember": {
            "type": "object",
            "properties": {
                "chat_id": {
                    "type": "integer"
                },
                "members_ids": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "handler.Refresh": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "handler.Sign": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "resp.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "description": "\"Укажите 'Bearer', а затем ваш JWT токен.\"",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Messenger API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

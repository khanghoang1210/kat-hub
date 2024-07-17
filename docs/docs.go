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
        "/accounts/login": {
            "post": {
                "description": "Login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "accounts"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Accounts",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            }
        },
        "/posts": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Get All Posts",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Update Post",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Post ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Posts",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreatePostReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Posts"
                ],
                "summary": "Create Posts",
                "parameters": [
                    {
                        "description": "Posts",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreatePostReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Get All Users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get All Users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            },
            "put": {
                "description": "Update Users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update Users",
                "parameters": [
                    {
                        "description": "Users",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.UpdateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create Users",
                "parameters": [
                    {
                        "description": "Users",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/requests.CreateUserReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "requests.CreatePostReq": {
            "type": "object",
            "properties": {
                "textContent": {
                    "type": "string"
                }
            }
        },
        "requests.CreateUserReq": {
            "type": "object",
            "required": [
                "confirmPassword",
                "email",
                "fullName",
                "password",
                "userName"
            ],
            "properties": {
                "confirmPassword": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullName": {
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
        "requests.LoginReq": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "requests.UpdateUserReq": {
            "type": "object",
            "required": [
                "userName"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "fullName": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "responses.ResponseData": {
            "type": "object",
            "properties": {
                "data": {},
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Kathub API",
	Description:      "Kathub API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

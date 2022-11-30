// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/ping": {
            "get": {
                "description": "this for check is alive service",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ping"
                ],
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "creating new transaction and return id string",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "transactions",
                "parameters": [
                    {
                        "description": "Create Param",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.CreateTransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "203": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{transaction_id}": {
            "get": {
                "description": "gives user transactions and user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transactions"
                ],
                "summary": "transactions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Transaction ID",
                        "name": "transaction_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/structs.TransactionResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "get all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/db.User"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "creating new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "users",
                "parameters": [
                    {
                        "description": "Create Param",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/structs.CreateUserRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/users/{user_id}": {
            "get": {
                "description": "gives user transactions and user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.UserTransactions"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/structs.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.User": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "db.UserTransactions": {
            "type": "object",
            "properties": {
                "transactions": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/structs.TransactionResponse"
                    }
                },
                "user": {
                    "$ref": "#/definitions/db.User"
                }
            }
        },
        "structs.CreateTransactionRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "operation_type": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "structs.CreateUserRequest": {
            "type": "object",
            "properties": {
                "username": {
                    "type": "string"
                }
            }
        },
        "structs.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "structs.TransactionResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "typeOperation": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
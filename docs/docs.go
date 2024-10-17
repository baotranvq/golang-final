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
        "/items": {
            "get": {
                "description": "This endpoint retrieves a list of all items.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Get all items",
                "responses": {
                    "200": {
                        "description": "List of items retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoint allows authenticated users to create an item.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Create a new item",
                "parameters": [
                    {
                        "description": "Item creation payload",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemCreation"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item successfully created",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        },
        "/items/{id}": {
            "get": {
                "description": "This endpoint retrieves a single item by its unique identifier.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Get an item by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item retrieved successfully",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format or bad request",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "404": {
                        "description": "Item not found",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            },
            "put": {
                "description": "This endpoint allows updating the properties of an existing item by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Update an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Item update payload",
                        "name": "item",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.ItemUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item updated successfully",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid input or bad request",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "404": {
                        "description": "Item not found",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            },
            "delete": {
                "description": "This endpoint deletes an item identified by its unique ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Items"
                ],
                "summary": "Delete an item",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Item ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Item deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid ID format or bad request",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "404": {
                        "description": "Item not found",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        },
        "/users": {
            "patch": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "This endpoint logs a user into the system and returns a token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "update user",
                "parameters": [
                    {
                        "description": "User login payload",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token successfully generated",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "This endpoint logs a user into the system and returns a token.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User login payload",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token successfully generated",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        },
        "/users/profile": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieve detailed information of a user based on their unique UUID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user by ID",
                "responses": {
                    "200": {
                        "description": "User successfully retrieved",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID or user not found",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "This endpoint registers a new user in the system.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User creation payload",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserCreate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully registered",
                        "schema": {
                            "$ref": "#/definitions/clients.SuccessRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/clients.AppError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "clients.AppError": {
            "type": "object",
            "properties": {
                "error_key": {
                    "type": "string"
                },
                "log": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                }
            }
        },
        "clients.SuccessRes": {
            "type": "object",
            "properties": {
                "data": {},
                "filter": {},
                "paging": {}
            }
        },
        "domain.ItemCreation": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.ItemUpdate": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "title": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "domain.Status": {
            "type": "integer",
            "enum": [
                0,
                1,
                2
            ],
            "x-enum-varnames": [
                "Deleted",
                "Active",
                "Done"
            ]
        },
        "domain.UserCreate": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.UserUpdate": {
            "type": "object",
            "properties": {
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
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
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
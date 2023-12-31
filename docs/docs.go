// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "murasame29",
            "email": "oogiriminister@gamil.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/encounters": {
            "get": {
                "description": "Get All Encounters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encounter"
                ],
                "summary": "Get All Encounters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "pageID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "pageSize",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "userID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/output.ListEncounterResponse"
                            }
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            },
            "post": {
                "description": "Create Encount entory",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encounter"
                ],
                "summary": "Create Encounter",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create encounter request",
                        "name": "CreateEncounterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateEncounterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.CreateEncounterResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/encounters/{encounter_id}": {
            "get": {
                "description": "Get All Encounters",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Encounter"
                ],
                "summary": "Get All Encounters",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "list encounter request",
                        "name": "encounter_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.ListEncounterResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/game": {
            "post": {
                "description": "Create Game",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Create Game",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create game request",
                        "name": "CreateGameRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateGameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.CreateGameResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/game/{game_id}": {
            "post": {
                "description": "Join Game",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Join Game",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "game id",
                        "name": "game_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "create game request",
                        "name": "JoinGameRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.JoinGameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.JoinGameResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/game/{game_id}/action": {
            "post": {
                "description": "Game Action",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Game Action",
                "parameters": [
                    {
                        "type": "string",
                        "description": "game id",
                        "name": "game_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create game request",
                        "name": "ReadyGameRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.ActionGameRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response"
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/game/{game_id}/ready": {
            "post": {
                "description": "Game Ready",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Game"
                ],
                "summary": "Game Ready",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create game request",
                        "name": "ReadyGameRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.ReadyGameRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "game id",
                        "name": "game_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response"
                    },
                    "400": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/users": {
            "post": {
                "description": "Create User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "create user request",
                        "name": "CreateUser",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/input.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "409": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        },
        "/v1/users/{user_id}": {
            "get": {
                "description": "Get any User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "create user request",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.ReadUserResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "409": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            },
            "delete": {
                "description": "Delete User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "create user request",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success response",
                        "schema": {
                            "$ref": "#/definitions/output.DeleteUserResponse"
                        }
                    },
                    "400": {
                        "description": "error response"
                    },
                    "409": {
                        "description": "error response"
                    },
                    "500": {
                        "description": "error response"
                    }
                }
            }
        }
    },
    "definitions": {
        "input.ActionGameRequest": {
            "type": "object",
            "properties": {
                "command_id": {
                    "type": "integer"
                },
                "gameID": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "input.CreateEncounterRequest": {
            "type": "object",
            "properties": {
                "encounted_user_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "input.CreateGameRequest": {
            "type": "object",
            "properties": {
                "owner_id": {
                    "type": "string"
                }
            }
        },
        "input.CreateUser": {
            "type": "object",
            "properties": {
                "github_id": {
                    "type": "string"
                }
            }
        },
        "input.JoinGameRequest": {
            "type": "object",
            "properties": {
                "gameID": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "input.ReadyGameRequest": {
            "type": "object",
            "properties": {
                "gameID": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "output.CreateEncounterResponse": {
            "type": "object",
            "properties": {
                "encounter_id": {
                    "type": "string"
                }
            }
        },
        "output.CreateGameResponse": {
            "type": "object",
            "properties": {
                "game_id": {
                    "type": "string"
                },
                "game_server_token": {
                    "type": "string"
                }
            }
        },
        "output.CreateUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "description": "ユーザの情報",
                    "type": "string"
                }
            }
        },
        "output.DeleteUserResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "output.JoinGameResponse": {
            "type": "object",
            "properties": {
                "game_id": {
                    "type": "string"
                },
                "game_server_token": {
                    "type": "string"
                }
            }
        },
        "output.ListEncounterResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "encounter_id": {
                    "type": "string"
                },
                "encouted_user_id": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "output.ReadUserResponse": {
            "type": "object",
            "properties": {
                "github_id": {
                    "type": "string"
                },
                "user_id": {
                    "description": "ユーザの情報",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "localhost:8080",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "PDD-GitHub-Go-Backend API",
	Description:      "This is a PDD-GitHub-Go-Backend API server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

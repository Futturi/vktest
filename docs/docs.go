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
        "/api/actors": {
            "get": {
                "description": "get all actors",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "GetAllActors",
                "operationId": "get-actors",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Actor"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "update actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "UpdateActor",
                "operationId": "update-actor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "insert actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "InsertActor",
                "operationId": "insert-actor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete actor",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "actors"
                ],
                "summary": "DeleteActor",
                "operationId": "delete-actor",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/cinemas": {
            "get": {
                "description": "get all cinemas",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cinemas"
                ],
                "summary": "GetAllCinemas",
                "operationId": "get-cinemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cinema"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "update cinema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cinemas"
                ],
                "summary": "UpdateCinema",
                "operationId": "update-cinemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "insert cinema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cinemas"
                ],
                "summary": "InsertCinema",
                "operationId": "insert-cinemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete cinema",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cinemas"
                ],
                "summary": "DeleteCinemas",
                "operationId": "get-cinemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/cinemas/search": {
            "post": {
                "description": "search cinema",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "cinemas"
                ],
                "summary": "SearchCinema",
                "operationId": "search-cinemas",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Cinema"
                            }
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/admin/signin": {
            "post": {
                "description": "login account 4 admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authAdmin"
                ],
                "summary": "SingInAdmin",
                "operationId": "login-account-admin",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/admin/signup": {
            "post": {
                "description": "create account 4 admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authAdmin"
                ],
                "summary": "SingUpAdmin",
                "operationId": "create-account-admin",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Admin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signin": {
            "post": {
                "description": "login account 4 user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SingInUser",
                "operationId": "login-account-user",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "create account 4 user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "SingUpUser",
                "operationId": "create-account-user",
                "parameters": [
                    {
                        "description": "account info",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "default": {
                        "description": "",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Actor": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "string"
                },
                "genre": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.Admin": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.Cinema": {
            "type": "object",
            "properties": {
                "actors": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "data": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "rating": {
                    "type": "integer"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "username": {
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

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Cinema App API",
	Description:      "API Server 4 Cinema Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

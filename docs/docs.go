// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
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
        "/auth/login": {
            "post": {
                "description": "Authenticates a user and provides a JWT to Authorize API calls",
                "produces": [
                    "application/json"
                ],
                "summary": "Generates a JWT",
                "operationId": "Authentication",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            }
        },
        "/health": {
            "get": {
                "description": "Checks Server Health While Running",
                "produces": [
                    "application/json"
                ],
                "summary": "Checks Health of Server",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            }
        },
        "/movies": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Gets all the existing Movies from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Displays all Movies Info",
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Adds a movie to database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Adds a Movie Info",
                "parameters": [
                    {
                        "description": "Add Movie",
                        "name": "movie",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/operation.Movie"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            }
        },
        "/movies/{id}": {
            "get": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Displays a movie in database",
                "produces": [
                    "application/json"
                ],
                "summary": "Displays one movie detail",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Display Movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Updates a movie in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Updates a Movie Info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Update Movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "bearerAuth": []
                    }
                ],
                "description": "Deletes a movie in database",
                "produces": [
                    "application/json"
                ],
                "summary": "Deletes a Movie Info",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Delete Movie",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "401": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "operation.Actor": {
            "type": "object",
            "required": [
                "firstName"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        },
        "operation.Movie": {
            "type": "object",
            "required": [
                "actor",
                "price",
                "trailer"
            ],
            "properties": {
                "actor": {
                    "$ref": "#/definitions/operation.Actor"
                },
                "desc": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "price": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "trailer": {
                    "type": "string"
                }
            }
        },
        "user.User": {
            "type": "object",
            "properties": {
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
        "bearerAuth": {
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
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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

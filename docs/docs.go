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
        "contact": {
            "name": "Juan Manuel Abanto Mera",
            "email": "jmanuelabanto@gmail.com"
        },
        "license": {
            "name": "MIT License",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/users": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create a new user.",
                "parameters": [
                    {
                        "description": "Object to be created.",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Id of the created object",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": ""
                    },
                    "422": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update a user.",
                "parameters": [
                    {
                        "description": "Object to be modified.",
                        "name": "command",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/command.UpdateUser"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "422": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get a user by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
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
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete a user by Id.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": ""
                    },
                    "400": {
                        "description": ""
                    },
                    "404": {
                        "description": ""
                    },
                    "500": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "command.CreateUser": {
            "type": "object",
            "required": [
                "birthDate",
                "email",
                "firstName",
                "lastName",
                "phoneNumber",
                "userName"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 50
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 20
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 20
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 50
                },
                "userName": {
                    "type": "string",
                    "maxLength": 10
                }
            }
        },
        "command.UpdateUser": {
            "type": "object",
            "required": [
                "birthDate",
                "email",
                "firstName",
                "id",
                "lastName",
                "phoneNumber",
                "userName"
            ],
            "properties": {
                "birthDate": {
                    "type": "string"
                },
                "email": {
                    "type": "string",
                    "maxLength": 50
                },
                "firstName": {
                    "type": "string",
                    "maxLength": 20
                },
                "id": {
                    "type": "string",
                    "maxLength": 24,
                    "minLength": 24
                },
                "lastName": {
                    "type": "string",
                    "maxLength": 20
                },
                "phoneNumber": {
                    "type": "string",
                    "maxLength": 50
                },
                "userName": {
                    "type": "string",
                    "maxLength": 10
                }
            }
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
	Version:     "v1",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "User API",
	Description: "Specifying services for users.",
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
	swag.Register("swagger", &s{})
}

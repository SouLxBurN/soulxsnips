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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/md": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "md"
                ],
                "summary": "Retrieve All Markdown Snippets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/md.MDListItem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "md"
                ],
                "summary": "Create new a markdown snippet",
                "parameters": [
                    {
                        "description": "Post Body",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/md.CreateMDReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/md.MarkdownSnippet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
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
                    "md"
                ],
                "summary": "Updates a markdown snippet",
                "parameters": [
                    {
                        "description": "Patch Body",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/md.UpdateMDReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/md.MarkdownSnippet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/md/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "md"
                ],
                "summary": "Retrieve Markdown Snippet",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Snippet uuid",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/md.MarkdownSnippet"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
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
                    "md"
                ],
                "summary": "Removes MarkdownSnippet permanantly",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Snippet uuid",
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
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "format": "int",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "format": "string"
                }
            }
        },
        "md.CreateMDReq": {
            "type": "object",
            "required": [
                "body",
                "title"
            ],
            "properties": {
                "body": {
                    "description": "Markdown body to save.",
                    "type": "string",
                    "maxLength": 64000,
                    "minLength": 1,
                    "example": "# Markdown Snippet\nSome Text"
                },
                "title": {
                    "description": "Markdown snippet title.",
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1,
                    "example": "SouLxBurN Is Awesome!"
                }
            }
        },
        "md.MDListItem": {
            "type": "object",
            "properties": {
                "createDate": {
                    "description": "Date markdown snippet was created",
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "description": "Markdown snippet guid.",
                    "type": "string",
                    "format": "uuid"
                },
                "title": {
                    "description": "Markdown snippet title.",
                    "type": "string",
                    "example": "SouLxBurN Is Awesome!"
                }
            }
        },
        "md.MarkdownSnippet": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Markdown body to save.",
                    "type": "string",
                    "example": "# Markdown Snippet\nSome Text"
                },
                "createDate": {
                    "description": "Date markdown snippet was created.",
                    "type": "string",
                    "format": "date-time"
                },
                "id": {
                    "description": "Markdown snippet guid.",
                    "type": "string",
                    "format": "uuid"
                },
                "title": {
                    "description": "Markdown snippet title.",
                    "type": "string",
                    "example": "SouLxBurN Is Awesome!"
                },
                "updateKey": {
                    "description": "Update hash key allowing the snippet to be updated.",
                    "type": "string",
                    "format": "uuid"
                }
            }
        },
        "md.UpdateMDReq": {
            "type": "object",
            "required": [
                "body",
                "id",
                "title",
                "updateKey"
            ],
            "properties": {
                "body": {
                    "description": "Markdown body to save.",
                    "type": "string",
                    "maxLength": 64000,
                    "minLength": 1,
                    "example": "# Markdown Snippet\nSome Text"
                },
                "id": {
                    "description": "Markdown snippet guid.",
                    "type": "string",
                    "format": "uuid"
                },
                "title": {
                    "description": "Markdown snippet title.",
                    "type": "string",
                    "maxLength": 64,
                    "minLength": 1,
                    "example": "SouLxBurN Is Awesome!"
                },
                "updateKey": {
                    "description": "UpdateKey required for updating snippet.",
                    "type": "string",
                    "format": "uuid"
                }
            }
        }
    },
    "tags": [
        {
            "name": "md"
        }
    ]
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "SouLxSnippets",
	Description: "Backend API for storing and retrieving markdown snippets",
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

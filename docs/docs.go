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
        "contact": {
            "name": "Sw-Saturn"
        },
        "license": {
            "name": "MIT License",
            "url": "https://github.com/MiCasa-IoT/MainAPI/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/db/document/create": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "新規のドキュメントを作成する",
                "parameters": [
                    {
                        "description": "Connection",
                        "name": "create",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Connection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.InsertOneResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/db/document/delete": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "既存のドキュメントを削除する",
                "parameters": [
                    {
                        "description": "UUID",
                        "name": "delete",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Connection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.DeleteResult"
                        }
                    },
                    "400": {
                        "description": "UUID Not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/db/document/read": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "指定したIDのドキュメントを取得する",
                "parameters": [
                    {
                        "description": "UUID",
                        "name": "read",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Connection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Connection"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/db/document/readall": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "全てのドキュメントを取得する",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Connection"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/db/document/update": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "既存のドキュメントを更新する",
                "parameters": [
                    {
                        "description": "UUID",
                        "name": "read",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Connection"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UpdateResult"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Connection": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "models.DeleteResult": {
            "type": "object",
            "properties": {
                "deletedCount": {
                    "description": "The number of documents deleted.",
                    "type": "integer"
                }
            }
        },
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                }
            }
        },
        "models.InsertOneResult": {
            "type": "object",
            "properties": {
                "insertedID": {
                    "description": "The _id of the inserted document. A value generated by the driver will be of type primitive.ObjectID.",
                    "type": "object"
                }
            }
        },
        "models.UpdateResult": {
            "type": "object",
            "properties": {
                "matchedCount": {
                    "description": "The number of documents matched by the filter.",
                    "type": "integer"
                },
                "modifiedCount": {
                    "description": "The number of documents modified by the operation.",
                    "type": "integer"
                },
                "upsertedCount": {
                    "description": "The number of documents upserted by the operation.",
                    "type": "integer"
                },
                "upsertedID": {
                    "description": "The _id field of the upserted document, or nil if no upsert was done.",
                    "type": "object"
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
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "MiCasa Main API",
	Description: "MiCasaのAPIドキュメント",
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
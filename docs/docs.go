// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "github.com/tetrex"
        },
        "license": {
            "name": "MIT License"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "returns server time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Get Health check status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/health.HelloResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrResponse"
                        }
                    }
                }
            }
        },
        "/home/articles": {
            "get": {
                "description": "returns array of articles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Homepage"
                ],
                "summary": "Get homepage articles",
                "parameters": [
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "description": "int valid",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 0,
                        "type": "integer",
                        "description": "int valid",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/homepage.ArticlesResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrResponse"
                        }
                    }
                }
            }
        },
        "/home/featured": {
            "get": {
                "description": "returns featured article",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Homepage"
                ],
                "summary": "Get featured article",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/homepage.FeaturedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrResponse"
                        }
                    }
                }
            }
        },
        "/home/press": {
            "get": {
                "description": "returns array of articles",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Homepage"
                ],
                "summary": "Get Press articles",
                "parameters": [
                    {
                        "maximum": 100,
                        "minimum": 1,
                        "type": "integer",
                        "description": "int valid",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "maximum": 100,
                        "minimum": 0,
                        "type": "integer",
                        "description": "int valid",
                        "name": "skip",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/homepage.PressResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/types.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "health.HelloResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "homepage.ArticlesResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/homepage.Blog"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "homepage.Blog": {
            "type": "object",
            "properties": {
                "author_name": {
                    "type": "string"
                },
                "author_profile_url": {
                    "type": "string"
                },
                "blog_created_at": {
                    "type": "string"
                },
                "blog_id": {
                    "type": "integer"
                },
                "blog_thumbnail_url": {
                    "type": "string"
                },
                "category": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "read_time": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "homepage.FeaturedResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/homepage.Blog"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "homepage.Press": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "description": {
                    "type": "string"
                },
                "external_url": {
                    "type": "string"
                },
                "press_id": {
                    "type": "integer"
                },
                "press_published_at": {
                    "type": "string"
                },
                "press_thumbnail_url": {
                    "type": "string"
                },
                "publisher_name": {
                    "type": "string"
                },
                "publisher_profile_img_link": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "homepage.PressResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/homepage.Press"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "types.ErrResponse": {
            "type": "object",
            "properties": {
                "err": {
                    "type": "string"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "server api",
	Description:      "This is a backend api server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

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
        "/ping": {
            "get": {
                "description": "连接测试",
                "summary": "ping",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/user/log/retrieve": {
            "get": {
                "summary": "获取用户日志",
                "parameters": [
                    {
                        "type": "string",
                        "description": "事件名称",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询开始时间 YYYY-MM-DD",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "查询结束时间 YYYY-MM-DD",
                        "name": "end",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        },
        "/user/log/upload": {
            "post": {
                "summary": "添加用户日志",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.PushJson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.EventInfo": {
            "type": "object",
            "properties": {
                "act_time": {
                    "type": "string"
                },
                "append": {
                    "type": "string"
                }
            }
        },
        "models.PushJson": {
            "type": "object",
            "properties": {
                "device_id": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "user_events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserEvent"
                    }
                }
            }
        },
        "models.UserEvent": {
            "type": "object",
            "properties": {
                "event_infos": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.EventInfo"
                    }
                },
                "event_name": {
                    "type": "string"
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
	Version:     "2.0",
	Host:        "localhost:9101",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "swagger 接口文档",
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

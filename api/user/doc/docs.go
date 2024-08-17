// Code generated by swaggo/swag. DO NOT EDIT
package doc

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
        "/provider": {
            "get": {
                "tags": [
                    "provider"
                ],
                "operationId": "listProviderAccount",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProviderList"
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "provider"
                ],
                "operationId": "addProviderAccount",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content - success without response body"
                    }
                }
            }
        },
        "/provider/{provider}": {
            "get": {
                "tags": [
                    "provider"
                ],
                "operationId": "getProviderAccount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    }
                }
            }
        },
        "/provider/{provider}/charge": {
            "post": {
                "description": "This endpoint allows you to add fund to an account",
                "tags": [
                    "provider"
                ],
                "operationId": "charge",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Provider"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        },
        "/provider/{provider}/refund": {
            "post": {
                "description": "This endpoint allows you to refund from an account",
                "tags": [
                    "provider"
                ],
                "operationId": "refund",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Refund"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        },
        "/provider/{provider}/service/{service}": {
            "get": {
                "tags": [
                    "service"
                ],
                "operationId": "getService",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service name",
                        "name": "service",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Service"
                        }
                    }
                }
            },
            "post": {
                "description": "This endpoint allows you to retrieve data based on provider and service. This endpoint acts as a proxy to retrieve data from various external services. The response type can vary depending on the service being accessed",
                "tags": [
                    "data"
                ],
                "operationId": "getData",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service name",
                        "name": "service",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Binary stream response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/provider/{provider}/service/{service}/{suffix}": {
            "post": {
                "description": "This endpoint acts as a proxy to retrieve data from various external services based on the provided ` + "`" + `provider` + "`" + ` and ` + "`" + `service` + "`" + ` parameters. The response type can vary depending on the external service being accessed. An optional ` + "`" + `suffix` + "`" + ` parameter can be appended to further specify the request for external services",
                "tags": [
                    "data"
                ],
                "operationId": "getDataWithSuffix",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Service name",
                        "name": "service",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Suffix",
                        "name": "suffix",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Binary stream response",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/provider/{provider}/sync": {
            "post": {
                "description": "This endpoint allows you to synchronize information of single account from the contract",
                "tags": [
                    "provider"
                ],
                "operationId": "syncProviderAccount",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Provider address",
                        "name": "provider",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        },
        "/service": {
            "get": {
                "tags": [
                    "service"
                ],
                "operationId": "listService",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ServiceList"
                        }
                    }
                }
            }
        },
        "/sync": {
            "post": {
                "description": "This endpoint allows you to synchronize information of all accounts from the contract",
                "tags": [
                    "provider"
                ],
                "operationId": "syncProviderAccounts",
                "responses": {
                    "202": {
                        "description": "Accepted"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ListMeta": {
            "type": "object",
            "properties": {
                "total": {
                    "type": "integer"
                }
            }
        },
        "model.Provider": {
            "type": "object",
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "lastResponseTokenCount": {
                    "type": "integer"
                },
                "nonce": {
                    "type": "integer"
                },
                "pendingRefund": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "refunds": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Refund"
                    }
                }
            }
        },
        "model.ProviderList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Provider"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/model.ListMeta"
                }
            }
        },
        "model.Refund": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string",
                    "readOnly": true
                },
                "index": {
                    "type": "integer",
                    "readOnly": true
                },
                "processed": {
                    "type": "boolean"
                },
                "provider": {
                    "type": "string"
                }
            }
        },
        "model.Service": {
            "type": "object",
            "properties": {
                "inputPrice": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "outputPrice": {
                    "type": "integer"
                },
                "provider": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "model.ServiceList": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Service"
                    }
                },
                "metadata": {
                    "$ref": "#/definitions/model.ListMeta"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "0G Serving User Agent API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

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
        "/auth/login": {
            "post": {
                "description": "Allows users to login into their account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login route",
                "parameters": [
                    {
                        "description": "User's email and password",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse_doc"
                        }
                    },
                    "400": {
                        "description": "Please provide with sufficient credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "401": {
                        "description": "Invalid Credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "404": {
                        "description": "User is not registered",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "422": {
                        "description": "Email already registered, please verify your email address",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    }
                }
            }
        },
        "/auth/otp": {
            "post": {
                "description": "Allows users to validate OTP and complete the registration process.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Validation route",
                "parameters": [
                    {
                        "description": "User's email address and otp",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.OTP"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful response, User already verified. Please login.",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse_doc"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON data, Invalid Email",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "401": {
                        "description": "Invalid OTP",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "404": {
                        "description": "User does not exist. Please register to generate OTP.",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "422": {
                        "description": "Please provide with sufficient credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Allows users to create a new account.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register route",
                "parameters": [
                    {
                        "description": "User name, email, password",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Register"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successful response",
                        "schema": {
                            "$ref": "#/definitions/responses.UserResponse_doc"
                        }
                    },
                    "400": {
                        "description": "Invalid JSON data, Invalid Email",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "401": {
                        "description": "Invalid Credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "409": {
                        "description": "User already exists",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "422": {
                        "description": "Please provide with sufficient credentials",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error, Error in inserting the document",
                        "schema": {
                            "$ref": "#/definitions/responses.ErrorResponse_doc"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Login": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.OTP": {
            "type": "object",
            "required": [
                "email",
                "otp"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "otp": {
                    "type": "string"
                }
            }
        },
        "model.Register": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "responses.ErrorResponse_doc": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "responses.UserResponse_doc": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "additionalProperties": true
                },
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/",
	Schemes:          []string{},
	Title:            "Registration API",
	Description:      "This is a registration api for an application.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

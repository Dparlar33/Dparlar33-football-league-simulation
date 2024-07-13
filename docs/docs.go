// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/leagues/": {
            "get": {
                "description": "Get a list of all leagues",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Get all leagues",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.League"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/leagues/add": {
            "post": {
                "description": "Create a new league",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Create a new league",
                "parameters": [
                    {
                        "description": "League to create",
                        "name": "league",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.League"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.League"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/leagues/league-table": {
            "get": {
                "description": "Get the current league table",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Get league table",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.TeamPosition"
                            }
                        }
                    }
                }
            }
        },
        "/leagues/play-all-week": {
            "get": {
                "description": "Simulate all remaining weeks and update the league",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Play all remaining weeks",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.TeamPosition"
                            }
                        }
                    }
                }
            }
        },
        "/leagues/play-one-week": {
            "get": {
                "description": "Simulate one week of matches and update the league",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Play one week of matches",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.WeeklyResult"
                        }
                    }
                }
            }
        },
        "/leagues/prediction": {
            "get": {
                "description": "Get the current championship predictions for each team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Get championship predictions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.TeamPrediction"
                            }
                        }
                    }
                }
            }
        },
        "/leagues/reset": {
            "get": {
                "description": "Reset all league data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leagues"
                ],
                "summary": "Reset the league",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/players/": {
            "get": {
                "description": "Get a list of all players",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Get all players",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.Player"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/players/add": {
            "post": {
                "description": "Create a new player",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "players"
                ],
                "summary": "Create a new player",
                "parameters": [
                    {
                        "description": "Player to create",
                        "name": "player",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Player"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Player"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/teams/": {
            "get": {
                "description": "Get a list of all teams",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "summary": "Get all teams",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dtos.Team"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/teams/add": {
            "post": {
                "description": "Create a new team",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "teams"
                ],
                "summary": "Create a new team",
                "parameters": [
                    {
                        "description": "Team to create",
                        "name": "team",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dtos.Team"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dtos.Team"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/dtos.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dtos.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "dtos.Match": {
            "type": "object",
            "properties": {
                "awayTeam": {
                    "type": "string"
                },
                "awayTeamScore": {
                    "type": "integer"
                },
                "hostTeam": {
                    "type": "string"
                },
                "hostTeamScore": {
                    "type": "integer"
                }
            }
        },
        "dtos.Player": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "teamID": {
                    "type": "integer"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "dtos.Team": {
            "type": "object",
            "properties": {
                "capacityOfPlayer": {
                    "type": "integer"
                },
                "drawn": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "leagueID": {
                    "type": "integer"
                },
                "lost": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "point": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                },
                "title": {
                    "type": "integer"
                },
                "valueOfPlayers": {
                    "type": "number"
                },
                "won": {
                    "type": "integer"
                }
            }
        },
        "dtos.TeamPosition": {
            "type": "object",
            "properties": {
                "club": {
                    "type": "string"
                },
                "drawn": {
                    "type": "integer"
                },
                "lost": {
                    "type": "integer"
                },
                "points": {
                    "type": "integer"
                },
                "position": {
                    "type": "integer"
                },
                "won": {
                    "type": "integer"
                }
            }
        },
        "dtos.TeamPrediction": {
            "type": "object",
            "properties": {
                "odds": {
                    "type": "number"
                },
                "teamName": {
                    "type": "string"
                }
            }
        },
        "dtos.WeeklyResult": {
            "type": "object",
            "properties": {
                "matches": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dtos.Match"
                    }
                },
                "week": {
                    "type": "integer"
                }
            }
        },
        "models.League": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "limit_of_teams": {
                    "type": "integer"
                },
                "names": {
                    "type": "string"
                },
                "teams": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Team"
                    }
                },
                "weeks": {
                    "type": "integer"
                }
            }
        },
        "models.Player": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "position": {
                    "type": "string"
                },
                "team_id": {
                    "type": "integer"
                },
                "value": {
                    "type": "integer"
                }
            }
        },
        "models.Team": {
            "type": "object",
            "properties": {
                "capacity_of_players": {
                    "type": "integer"
                },
                "deuces": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "league_id": {
                    "type": "integer"
                },
                "loses": {
                    "type": "integer"
                },
                "names": {
                    "type": "string"
                },
                "players": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Player"
                    }
                },
                "points": {
                    "type": "integer"
                },
                "scores": {
                    "type": "integer"
                },
                "titles": {
                    "type": "integer"
                },
                "value_of_players": {
                    "type": "number"
                },
                "wins": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "League Management API",
	Description:      "This is a sample server for managing leagues.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
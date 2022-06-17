{
    "openapi": "3.0.2",
    "servers": [
        {
            "url": "/"
        }
    ],
    "info": {
        "title": "{{.Title}}",
        "version": "there is no version :)",
        "license": {
            "name": "GNU GENERAL PUBLIC LICENSE VERSION 3",
            "url": "https://www.gnu.org/licenses/gpl-3.0.txt"
        },
        "description": "{{.Description}}",
        "contact": {
            "email": "info@karimi.dev",
            "name": "Muhammed Hussein Karimi",
            "url": "https://karimi.dev"
        }
    },
    "paths": {
        "/healthz/{name}": {
            "get": {
                "operationId": "healthz",
                "summary": "Healthz",
                "description": "Say that the service is healthy to the name",
                "parameters": [
                    {
                        "in": "path",
                        "name": "name",
                        "schema":{
                            "type": "string"
                        },
                        "required": true,
                        "example": "Steve"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/greeting":{
            "get":{
                "operationId": "getGreeting",
                "summary": "Greeting",
                "description": "Say hello and time",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    }
}
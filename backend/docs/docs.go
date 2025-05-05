// Paquete docs Código generado por swaggo/swag. NO EDITAR
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
    "paths": {}
}`

// SwaggerInfo contiene la información exportada de Swagger para que los clientes puedan modificarla
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",                                                    // Versión de la API
	Host:             "localhost:8080",                                           // Host donde se ejecuta la API
	BasePath:         "/",                                                        // BasePath de la API
	Schemes:          []string{"http"},                                           // Esquema utilizado (http o https)
	Title:            "Documentación de la API",                                  // Título de la documentación
	Description:      "Esta es la documentación de la API generada con Swagger.", // Descripción de la API
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}

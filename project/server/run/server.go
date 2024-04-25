package main

import (
	"fmt"
	"log"

	dbRoadProduct "RoadProduct"
	dbRoadPropiedades "RoadPropiedades"

	WritingCart "WritingCart" // Importa el paquete WritingCart

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
    // Inicializa el servidor de Fiber
    app := fiber.New()
    app.Use(cors.New())
    app.Use(logger.New())

	// Ruta POST para guardar un string
	app.Post("/addrocart", WritingCart.AddToCartHandler) // Asigna el handler AddToCartHandler para la ruta POST /cart
    app.Post("/clearcart", WritingCart.ClearCartHandler)
    

    // Agrega un nuevo endpoint para obtener el contenido del carrito
    app.Get("/getCartContent", func(c fiber.Ctx) error {
        // Obtén el contenido del carrito desde WritingCart
        cartIDs := WritingCart.GetCartIDs()
        fmt.Printf("El tipo de dato de cartIDs es: %T\n", cartIDs)

        // Devuelve el contenido del carrito como respuesta en formato JSON
        return c.JSON(cartIDs)
    })

    // Ruta para manejar la solicitud y devolver el JSON de propiedades
    app.Get("/propiedades", func(c fiber.Ctx) error {
        // Llama a la función ObtenerPropiedades para obtener el JSON de propiedades
        propiedadesJSON, err := dbRoadPropiedades.ObtenerPropiedades()
        if err != nil {
            log.Printf("Error al obtener propiedades: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener propiedades")
        }
        fmt.Printf("El tipo de dato de propiedades es: %T\n", propiedadesJSON)
        // Devuelve el JSON como respuesta
        return c.JSON(propiedadesJSON)
    })

    // Ruta para manejar la solicitud y devolver el JSON de productos
    app.Get("/product", func(c fiber.Ctx) error {
        // Llama a la función ObtenerProduct para obtener el JSON de productos
        productJSON, err := dbRoadProduct.ObtenerProduct()
        if err != nil {
            log.Printf("Error al obtener productos: %v", err)
            return c.Status(fiber.StatusInternalServerError).SendString("Error al obtener productos")
        }
        fmt.Printf("El tipo de dato de productos es: %T\n", productJSON)
        // Devuelve el JSON como respuesta
        return c.JSON(productJSON)
    })

    // Lanza el servidor Fiber
    fmt.Println("Iniciando servidor Fiber...")
    log.Fatal(app.Listen(":3000"))
}
package main

import (
	"log"

	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/migration"
	"github.com/Artzy1401/clone-cineplex-backend-4/route"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	// "database/sql"
	// "fmt"
	//_ "github.com/lib/pq"
)

// const (
// 	host = "db.vjpgdoelfnmgwmpdokkn.supabase.co"
// 	port = 5432
// 	user = "postgres"
// 	password = "SuretyBond2023!"
// 	dbname = "Cineplex_Team_4"
//   )

func main() {
	// INITIAL DATABASE
	database.DatabaseInit()

	// MIGRATE DATABASE
	migration.RunMigration()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	// // INITIAL ROUTE
	route.RouteInit(app)

    // err := app.Listen(":8080")
    // if err != nil {
    //     return
    // }

    app.Get("/", func (c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })

    log.Fatal(app.Listen(":8080"))

	// INITIAL 

	// app.Listen(":8080")
}
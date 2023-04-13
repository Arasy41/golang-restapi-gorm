package main

import (
	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/migration"
	"github.com/Artzy1401/clone-cineplex-backend-4/route"
	"github.com/gofiber/fiber/v2"
	// "database/sql"
	// "fmt"
	_ "github.com/lib/pq"

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
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	//   panic(err)
	// }
	// defer db.Close()
	// err = db.Ping()
	// if err != nil {
	//   panic(err)
	// }
	// fmt.Println("Established a successful connection!")

	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	// INITIAL 

	app.Listen(":8080")
}
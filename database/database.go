package database

import (
	//"database/sql"
	"fmt"
	//"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

const (
	host = "db.vjpgdoelfnmgwmpdokkn.supabase.co"
	port = 5432
	user = "postgres"
	password = "SuretyBond2023!"
	dbname = "Cineplex_Team_4"
  )

var DB *gorm.DB

func DatabaseInit() {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not Connect to database")
	}
	fmt.Println("Connected to database")

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " + "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// DB, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	//   panic(err)
	// }
	// defer DB.Close()
	// err = DB.Ping()
	// if err != nil {
	//   panic(err)
	// }
	// fmt.Println("Established a successful connection!")
}

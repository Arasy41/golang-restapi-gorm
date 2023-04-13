package migration

import (
	"log"
	"fmt"
	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Film{}, &entity.Theather{}, &entity.Comment{}, &entity.TheaterList{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
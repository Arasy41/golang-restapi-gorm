package migration

import (
	"log"
	"fmt"
	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.CommentList{})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Database Migrated")
}
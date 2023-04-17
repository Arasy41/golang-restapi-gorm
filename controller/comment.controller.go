package controller

import (
	"log"

	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/entity"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/request"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func CommentControllerGetComments(ctx *fiber.Ctx) error {
	// filmId := ctx.QueryInt("filmId")
    // var film []entity.Comment
    // err := database.DB.Raw(`
	// 	SELECT f.id, f.name, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like, c.comment
	// 	FROM films f
	// 	INNER JOIN comments c ON c.film_id = f.id
	// 	WHERE c.film_id = ?`, filmId).Scan(&film)

    // if err.Error != nil{
    //     log.Println(err.Error)
    // }
    // return ctx.JSON(film)

	var film []entity.Comment
	result := database.DB.Find(&film)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&film).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(film)
}

func CommentControllerCreate(ctx *fiber.Ctx) error {
	Comment := new(request.CommentCreateRequest)
	if err := ctx.BodyParser(Comment); err != nil {
		return err	
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(Comment)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}

	newComment := entity.Comment{
		FilmID: Comment.FilmID,
		Comment: Comment.Comment,
	}

	errComment := database.DB.Create(&newComment).Error
	if errComment != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create comment",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newComment,
	})
}

func CommentControllerDelete(ctx *fiber.Ctx) error {
	commentid := ctx.Params("id")
	var comment entity.Comment

	// CHECK AVAILABLE COMMENT
	err := database.DB.Debug().First(&comment, "id=?" ,&commentid).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "comment Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&comment).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "comment deleted",
	})
}
package controller 

import (
	"log"

	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/entity"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/request"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func TheaterControllerGetAll(ctx *fiber.Ctx) error {
	// userInfo := ctx.Locals("userInfo")
	// log.Println("user info data :: ", userInfo)
	var Theater []entity.Theather
	result := database.DB.Find(&Theater)
	if result.Error != nil {
		log.Println(result.Error)
	}

	// err := database.DB.Find(&User).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(Theater)
}

func TheaterControllerGetDetails(ctx *fiber.Ctx) error {
	theaterid := ctx.QueryInt("theaterid")

	var theater entity.Theather
	err := database.DB.First(&theater, "id = ?", theaterid).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "theater not found",
		})
	}

	var film []entity.TheaterId
    err = database.DB.Raw(`
		SELECT f.id, f.name, l.theater_id AS theater_id, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like
		FROM films f
		INNER JOIN theater_lists l ON l.film_id = f.id
		WHERE l.theater_id = ?`, theaterid).Scan(&film).Error
	
	var theaterdetails entity.TheaterDetails
	theaterdetails.ID = theater.ID
	theaterdetails.Film = film

    return ctx.JSON(fiber.Map{
		"theater": theater,
		"film" : film,
		"message": "successfully",
	})
}

func TheaterControllerCreate(ctx *fiber.Ctx) error {
	Theater := new(request.TheaterCreateRequest)
	if err := ctx.BodyParser(Theater); err != nil {
		return err
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(Theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}

	newTheater := entity.Theather{
		Kota: Theater.Kota,
		Cinema: Theater.Cinema,
		Contact: Theater.Contact,
	}

	errCreateUser := database.DB.Create(&newTheater).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newTheater,
	})
}

func TheaterControllerCreateTheaterList(ctx *fiber.Ctx) error {
	Theater := new(request.TheaterListCreateRequest)
	if err := ctx.BodyParser(Theater); err != nil {
		return err	
	}

	// VALIDASI REQUEST
 	validate := validator.New()
	errValidate := validate.Struct(Theater)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message" : "failed to validate",
			"error" : errValidate.Error(),
		})
	}

	newTheaterList := entity.TheaterList{
		TheaterID: Theater.TheaterID,
		FilmID: Theater.FilmID,
	}

	errTheater := database.DB.Create(&newTheaterList).Error
	if errTheater != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": newTheaterList,
	})
}

func TheaterControllerGetById(ctx *fiber.Ctx) error{
	TheaterId := ctx.Params("id")

	var Theater entity.Theather
	err := database.DB.First(&Theater, "id = ?", TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Theater not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": Theater,
	})
}

func TheaterControllerUpdate(ctx *fiber.Ctx) error {
	TheaterRequest := new(request.TheaterUpdateRequest)
	if err := ctx.BodyParser(TheaterRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var Theater entity.Theather

	TheaterId := ctx.Params("id")
	// CHECK AVALAIBLE Theater
	err := database.DB.First(&Theater, "id = ?", TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Theater not found",
		})
	}

	// UPDATE Theater DATA
	Theater.Kota = TheaterRequest.Kota
	Theater.Cinema = TheaterRequest.Cinema
	Theater.Contact = TheaterRequest.Contact

	errUpdate := database.DB.Save(&Theater).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": Theater,
	})
}

func TheaterControllerDelete(ctx *fiber.Ctx) error {
	TheaterId := ctx.Params("id")
	var Theater entity.Theather

	// CHECK AVAILABLE Theater
	err := database.DB.Debug().First(&Theater, "id=?" ,&TheaterId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Theater Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&Theater).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Theater deleted",
	})
}
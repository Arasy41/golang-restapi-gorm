package controller

import (
	// "database/sql"
	"fmt"
	"log"
	// "sort"
	"strconv"

	"github.com/Artzy1401/clone-cineplex-backend-4/database"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/entity"
	"github.com/Artzy1401/clone-cineplex-backend-4/model/request"
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

func FilmControllerGetAll(ctx *fiber.Ctx) error {
	var film []entity.Film
	
	sql := "SELECT * FROM films"

	if s := ctx.Query("s"); s != "" {
		sql = fmt.Sprintf("%s WHERE title LIKE '%%%s%%' OR description LIKE '%%%s%%'", sql, s, s)
	}

	if sort := ctx.Query("sort"); sort != "" {
		sql = fmt.Sprintf("%s ORDER BY id %s", sql, sort)
	}

	page, _ := strconv.Atoi(ctx.Query("page", "1"))
		perPage := 9
		var total int64

		database.DB.Raw(sql).Count(&total)

		sql = fmt.Sprintf("%s LIMIT %d OFFSET %d", sql, perPage, (page-1)*perPage)

	database.DB.Raw(sql).Scan(&film)

	// err := database.DB.Find(&film).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(fiber.Map{
		"film": film,
		"total": total,
		"page": page,
	})
}

func FilmHandlerGetByTheaterId(ctx *fiber.Ctx) error {
    theaterid := ctx.QueryInt("theaterid")
    var film []entity.TheaterId
    err := database.DB.Raw(`
		SELECT f.id, f.name, l.theater_id AS theater_id, f.jenis_film, f. produser, f.sutradara, f.penulis, f.produksi, f.casts, f.sinopsis, f.like
		FROM films f
		INNER JOIN theater_lists l ON l.film_id = f.id
		WHERE l.theater_id = ?`, theaterid).Scan(&film)

    if err.Error != nil{
        log.Println(err.Error)
    }

    return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}


func FilmControllerCreate(ctx *fiber.Ctx) error {
    film := new(request.FilmCreateRequest)
    if err := ctx.BodyParser(film); err != nil {
        return err
    }

    // VALIDASI REQUEST
    validate := validator.New()
    errValidate := validate.Struct(film)
    if errValidate != nil {
        return ctx.Status(400).JSON(fiber.Map{
            "message": "Gagal",
            "error":   errValidate.Error(),
        })
    }

    //HANDLE FILE
    file, errFile := ctx.FormFile("cover")
    if errFile != nil {
        log.Println("Error File: ", errFile)
    }

    filename := file.Filename

    errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/asset/%s", filename))
    if errSaveFile != nil {
        log.Println("File gagal disimpan")
    }

    newFilm := entity.Film{
        Name:     film.Name,
        JenisFilm: film.JenisFilm,
        Produser:  film.Produser,
        Sutradara: film.Sutradara,
        Penulis:   film.Penulis,
        Produksi:  film.Produksi,
        Casts:     film.Casts,
        Sinopsis:  film.Sinopsis,
        Cover:     filename,
    }

    errCreateFilm := database.DB.Create(&newFilm).Error
    if errCreateFilm != nil {
        return ctx.Status(500).JSON(fiber.Map{
            "message": "Tidak berhasil menyimpan data",
        })
    }

    return ctx.JSON(fiber.Map{
        "message": "Berhasil",
        "data":    newFilm,
    })
}



func FilmControllerGetById(ctx *fiber.Ctx) error{
	filmId := ctx.Params("id")

	var film entity.Film
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}

func FilmControllerUpdate(ctx *fiber.Ctx) error {
	filmRequest := new(request.FilmUpdateRequest)
	if err := ctx.BodyParser(filmRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var film entity.Film

	filmId := ctx.Params("id")
	// CHECK AVALAIBLE FILM
	err := database.DB.First(&film, "id = ?", filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	// UPDATE USER DATA
	if filmRequest.Name != "" {
		film.Name = filmRequest.Name
	}
	film.JenisFilm = filmRequest.JenisFilm
	film.Produksi = filmRequest.Produksi
	film.Sutradara = filmRequest.Sutradara
	film.Penulis = filmRequest.Penulis
	film.Produksi =	filmRequest.Produksi
	film.Casts = filmRequest.Casts
	film.Sinopsis =	filmRequest.Sinopsis
	film.Like = filmRequest.Like

	errUpdate := database.DB.Save(&film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": film,
	})
}

func FilmControllerDelete(ctx *fiber.Ctx) error {
	filmId := ctx.Params("id")
	var film entity.Film

	// CHECK AVAILABLE FILM
	err := database.DB.Debug().First(&film, "id=?" ,&filmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "Film Not Found",
		})
	}

	errDelete := database.DB.Debug().Delete(&film).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "film deleted",
	})
}

func FilmControllerLikeUpdate(ctx *fiber.Ctx) error{
	LikeRequest := new(request.FilmLikeUpdateRequest)
	if err := ctx.BodyParser(LikeRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "Bad Request",
		})
	}

	var Film entity.Film
	FilmId := ctx.Params("id")
	// CHECK AVALAIBLE Film
	err := database.DB.First(&Film, "id = ?", FilmId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "film not found",
		})
	}

	// UPDATE USER DATA
	Film.Like = LikeRequest.Like

	errUpdate := database.DB.Save(&Film).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "successfully",
		"data": Film,
	})
}
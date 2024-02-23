package controllers

import (
	"goravel/app/models"
	"log"

	"github.com/goravel/framework/contracts/http"
)

type UserController struct {
	//Dependent services
}

func NewUserController() *UserController {
	return &UserController{
		//Inject services
	}
}

func (r *UserController) Show(ctx http.Context) http.Response {
	return ctx.Response().Success().Json(http.Json{
		"Hello": "Goravel",
	})
}

func GetAll() []models.User {

	rows, err := facades.DB.Query("SELECT id_pesanan, tanggal, nomor_polisi, nama_motor, jenis_pesanan FROM customer")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var customers []models.User

	for rows.Next() {
		var customer models.User
		if err := rows.Scan(&customer.Id_pesanan, &customer.Tanggal, &customer.Nomor_polisi, &customer.Nama_motor, &customer.Jenis_pesanan); err != nil {
			log.Fatal(err)
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return customers
}

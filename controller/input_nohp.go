package controller

import (
	"jtn/database"
	"jtn/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

var UserList = make(map[int]model.No_handphone, 0)
var LastUser int = 1

func GetData(c echo.Context) (err error) { //menampilkan seluruh data user
	//static API
	var result []model.No_handphone

	for k := range UserList {
		result = append(result, UserList[k])
	}

	// Mengambil data dari database
	db := database.DBManager()
	user := []model.No_handphone{}
	if err := db.Find(&user).Error; err != nil {
		// Handle kesalahan yang terjadi saat mengambil data dari database
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Gagal mengambil data pengguna"})
	}

	// Menggabungkan data dari database dengan data dari "static API"
	result = append(result, user...)

	return c.JSON(http.StatusOK, result)
}

func CreateData(c echo.Context) (err error) { //menambahkan data konsumen
	req := new(model.No_handphone)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, nil)
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	//static API
	user := model.No_handphone{
		Id:         LastUser,
		No_hp:      req.No_hp,
		Provider:   req.Provider,
		Keterangan: req.Keterangan,
	}
	UserList[LastUser] = user
	LastUser++

	db := database.DBManager()
	if err := db.Create(&user).Error; err != nil {
		// Handle kesalahan yang terjadi saat menyimpan ke database
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result := db.Create(&user)

	return c.JSON(http.StatusCreated, result)
}

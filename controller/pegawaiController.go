package controller

import (
	"echo/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func FecthAll(c echo.Context) error {
	result, err := models.FetchAllPegawai()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func StorePegawai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	hp := c.FormValue("hp")

	result, err := models.StorePegawai(id, nama, alamat, hp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, result)
}

func UpdatePegawai(c echo.Context) error {
	id := c.FormValue("id")
	nama := c.FormValue("nama")
	alamat := c.FormValue("alamat")
	hp := c.FormValue("hp")

	id_pegawai, _ := strconv.Atoi(id)

	result, err := models.UpdatePegawai(id_pegawai, nama, alamat, hp)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeletePegawai(c echo.Context) error {
	id := c.FormValue("id")
	id_pegawai, _ := strconv.Atoi(id)
	result, err := models.DeletePegawai(id_pegawai)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)

}

package controllers

import (
	"delegaciafacil/configDB"
	"delegaciafacil/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetDelegacias(c *gin.Context) {
	var delegacias []models.Delegacia
	configDB.GetDB().Find(&delegacias)
	c.JSON(http.StatusOK, delegacias)
}

func GetDelegaciasPorHorario(c *gin.Context) {
	horario24hStr := c.Query("horario24h")
	var delegacias []models.Delegacia

	horario24h, err := strconv.ParseBool(horario24hStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parâmetro 'horario24h' deve ser 'true' ou 'false'"})
		return
	}

	configDB.GetDB().Where("horario24h = ?", horario24h).Find(&delegacias)
	c.JSON(http.StatusOK, delegacias)
}
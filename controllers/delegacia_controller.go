package controllers

import (
	"delegaciafacil/configDB"
	"delegaciafacil/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// func GetDelegacias(c *gin.Context) {
// 	var delegacias []models.Delegacia
// 	configDB.GetDB().Find(&delegacias)
// 	c.JSON(http.StatusOK, delegacias)
// }

// func GetDelegaciaByID(c *gin.Context) {
// 	id := c.Param("id")
// 	var delegacia models.Delegacia
// 	if err := configDB.GetDB().First(&delegacia, id).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Delegacia não encontrada"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, delegacia)
// }

// func GetDelegaciasByHorario(c *gin.Context) {
// 	var delegacias []models.Delegacia
// 	configDB.GetDB().Where("horario24h = ?", true).Find(&delegacias)
// 	c.JSON(http.StatusOK, delegacias)
// }

// func GetDelegaciasByTipo(c *gin.Context) {
// 	tipo := c.Query("tipo")
// 	var delegacias []models.Delegacia
// 	configDB.GetDB().Joins("JOIN delegacia_tipos ON delegacias.id = delegacia_tipos.delegacia_id").
// 		Joins("JOIN tipos ON tipos.id = delegacia_tipos.tipo_id").
// 		Where("tipos.nome = ?", tipo).
// 		Find(&delegacias)
// 	c.JSON(http.StatusOK, delegacias)
// }

func GetDelegacias(c *gin.Context) {
    var delegacias []models.Delegacia
    db := configDB.GetDB()

    // Filtrar por horário 24h
    horario24h := c.Query("horario24h")
    if horario24h != "" {
        db = db.Where("horario24h = ?", horario24h == "true")
    }

    // Filtrar por tipo
    tipos := c.Query("tipos")
    if tipos != "" {
        tipoList := strings.Split(tipos, ",")
        db = db.Where("tipo IN ?", tipoList)
    }

    db.Find(&delegacias)
    c.JSON(http.StatusOK, delegacias)
}
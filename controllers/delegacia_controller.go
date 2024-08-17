package controllers

import (
	"delegaciafacil/configDB"
	"delegaciafacil/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDelegacias(c *gin.Context) {
	var delegacias []models.Delegacia
	configDB.GetDB().Find(&delegacias)
	c.JSON(http.StatusOK, delegacias)
}

func GetDelegaciaByID(c *gin.Context) {
	id := c.Param("id")
	var delegacia models.Delegacia
	if err := configDB.GetDB().First(&delegacia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Delegacia não encontrada"})
		return
	}
	c.JSON(http.StatusOK, delegacia)
}

//func CreateDelegacia(c *gin.Context) {
//    var delegacia models.Delegacia
//    if err := c.ShouldBindJSON(&delegacia); err != nil {
//        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//        return
//    }
//    configDB.GetDB().Create(&delegacia)
//    c.JSON(http.StatusCreated, delegacia)
//}

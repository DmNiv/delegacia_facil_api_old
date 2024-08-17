package main

import (
	"delegaciafacil/configDB"
	"delegaciafacil/models"
	"delegaciafacil/routes"
)

func main() {
	configDB.Connect()
	configDB.GetDB().AutoMigrate(&models.Delegacia{}, &models.Tipo{})
	
	r := routes.SetupRoutes()
	r.Run(":8080")
}


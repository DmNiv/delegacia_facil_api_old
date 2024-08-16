package models

import "github.com/jinzhu/gorm"

type Delegacia struct {
	gorm.Model
	Nome      string  `json:"nome"`
	Endereco  string  `json:"endereco"`
	Tipo      string  `json:"tipo"`
	Horario24 bool    `json:"horario24"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

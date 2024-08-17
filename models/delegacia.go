package models

type Delegacia struct {
	ID         int     `json:"id" gorm:"primary_key"`
	Nome       string  `json:"nome"`
	Endereco   string  `json:"endereco"`
	Horario24h bool    `json:"horario24h"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}
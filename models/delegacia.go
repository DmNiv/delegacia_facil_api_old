package models

type Delegacia struct {
	ID         int     `json:"id" gorm:"primary_key"`
	Nome       string  `json:"nome"`
	Endereco   string  `json:"endereco"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Horario24h bool    `json:"horario24h"`
	Tipos      []Tipo  `json:"tipos" gorm:"many2many:delegacia_tipos;"`
}

type Tipo struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Nome string `json:"nome"`
}

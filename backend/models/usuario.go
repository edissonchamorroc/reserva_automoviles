package models

type Usuario struct {
	Cedula     string `json:"cedula" db:"cedula"`
	Contrasena string `json:"contrasena" db:"contrasena"`
}

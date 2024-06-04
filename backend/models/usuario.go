package models

type Usuario struct {
	Cedula     string `json:"cedula" db:"cedula"`
	Nombre     string `json:"nombre" db:"nombre"`
	Celular    string `json:"celular" db:"celular"`
	Email      string `json:"email" db:"email"`
	Contrasena string `json:"contrasena" db:"contrasena"`
}

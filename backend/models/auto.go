package models

type Auto struct {
	Auto_id     int    `json:"id" db:"auto_id"`
	Marca       string `json:"marca" db:"marca"`
	Modelo      int    `json:"modelo" db:"modelo"`
	Combustible string `json:"combustible" db:"combustible"`
	Transmision string `json:"transmision" db:"transmision"`
	Lujos       string `json:"lujos" db:"lujos"`
	Sillas      int    `json:"sillas" db:"sillas"`
	Sillabb     string `json:"sillabb" db:"sillabb"`
	Seguros     string `json:"seguros" db:"seguros"`
	Cedula      string `json:"cedula" db:"cedula"`
	Img         string `json:"img" db:"img"`
	Precio      int    `json:"precio" db:"precio"`
}

package models

type Auto struct {
	marca string `json:"marca" db:"marca"`
	modelo int `json:"modelo" db:"modelo"`
	combustible string `json:"combustible" db:"combustible"`
	transmision string `json:"transmision" db:"transmision"`
}
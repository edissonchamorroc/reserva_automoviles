package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/models"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/repositories"
)


type ControllerAuto struct {
	repo repositories.Repository[models.Auto]
}

func NewControllerAuto(repo repositories.Repository[models.Auto]) (*ControllerAuto, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es neceario un repositorio valido")
	}
	return &ControllerAuto{
		repo: repo,
	}, nil
}

func (c *ControllerAuto) ListarAutos(cedula string) ([]byte, error) {

	var buscarAutoPorCedula = fmt.Sprintf("SELECT * FROM autos where cedula='%s' limit $1 offset $2", cedula)
	autos, _, err := c.repo.List(context.Background(), buscarAutoPorCedula, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("fallo al listar autos, con error: %s", err.Error())
	}
	jsonAutos, err := json.Marshal(autos)
	if err != nil {
		return nil, fmt.Errorf("fallo convertir autos a json: %s", err.Error())
	}
	return jsonAutos, nil
}

func (c *ControllerAuto) ActualizarReserva(id string, cedula string) error {

	var query = fmt.Sprintf("update autos SET cedula=:cedula WHERE auto_id='%s'",id)
	valoresActualizar := make(map[string]any)
	valoresActualizar["cedula"] = cedula
	err := c.repo.Update(context.Background(), query, valoresActualizar)
	if err != nil {
		log.Printf("fallo en la actualizacion reserva: %s", err.Error())
		return fmt.Errorf("fallo en la actualizacion reserva: %s", err.Error())
	}
	return nil
}


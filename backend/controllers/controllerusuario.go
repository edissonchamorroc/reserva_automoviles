package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/models"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/repositories"
)

var (
	buscarUsuarioByCedula = "SELECT * FROM usuarios where cedula=$1"
	registrarUsuario = "insert into usuarios (cedula, nombre, celular, email, contrasena) values (:cedula, :nombre, :celular, :email, :contrasena)"
)

type ControllerUsuario struct {
	repo repositories.Repository[models.Usuario]
}

func NewControllerUsuario(repo repositories.Repository[models.Usuario]) (*ControllerUsuario, error) {
	if repo == nil {
		return nil, fmt.Errorf("para un controlador es neceario un repositorio valido")
	}
	return &ControllerUsuario{
		repo: repo,
	}, nil
}

func (c *ControllerUsuario) BuscarUsuario(cedula string) (bool, error) {
	
	_, err := c.repo.Read(context.Background(), buscarUsuarioByCedula, cedula)
	if err != nil {
		log.Printf("Usuario no registrado: %s", err.Error())
		return false, fmt.Errorf("usuario no registrado, con error: %s", err.Error())
	}
	return true, nil
}

func (c *ControllerUsuario) RegistrarUsuario(body []byte) (bool, error) {

	nuevoUsuario := &models.Usuario{}

	err := json.Unmarshal(body, nuevoUsuario)
	if err != nil {
		log.Printf("datos no validos: %s", err.Error())
		return false, fmt.Errorf("falla por datos no validos, con error: %s", err.Error())
	}

	valoresColumna := map[string]any{
		"cedula":     nuevoUsuario.Cedula,
		"nombre":     nuevoUsuario.Nombre,
		"celular":    nuevoUsuario.Celular,
		"email":      nuevoUsuario.Email,
		"contrasena": nuevoUsuario.Contrasena,
	}

	_, err = c.repo.Create(context.Background(), registrarUsuario, valoresColumna)
	if err != nil {
		log.Printf("falló al registrar usuario: %s", err.Error())
		return false, fmt.Errorf("falló al registrar usuario, con error: %s", err.Error())
	}
	return true, nil
}

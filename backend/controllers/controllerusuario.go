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
	registrarUsuario      = "insert into usuarios (cedula, contrasena) values (:cedula, :contrasena)"
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

func (c *ControllerUsuario) RegistrarUsuario(cedula string, contrasena string) ([]byte, error) {

	valoresColumna := map[string]any{
		"cedula":     cedula,
		"contrasena": contrasena,
	}

	_, err := c.repo.Create(context.Background(), registrarUsuario, valoresColumna)
	if err != nil {
		log.Printf("falló al registrar usuario: %s", err.Error())
		return nil, fmt.Errorf("falló al registrar usuario, con error: %s", err.Error())
	}
	usuarioJson, _ := json.Marshal(valoresColumna)
	return usuarioJson, nil
}

func (c *ControllerUsuario) Authenticate(cedula string, contrasena string) ([]byte, bool) {

	usuario, err := c.repo.Read(context.Background(), buscarUsuarioByCedula, cedula)

	if err != nil {
		log.Printf("Usuario no registrado: %s", err.Error())
		return nil, false
	}
	if contrasena != usuario.Contrasena {
		log.Printf("Las credenciales no son validas")
		return nil, false
	}
	valoresRetornar := map[string]any{
		"cedula": cedula,
	}
	usuarioJson, _ := json.Marshal(valoresRetornar)

	return usuarioJson, true
}

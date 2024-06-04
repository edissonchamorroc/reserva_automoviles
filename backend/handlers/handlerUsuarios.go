package handlers

import (
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/controllers"
)

type HandlerUsuarios struct {
	cUsuario *controllers.ControllerUsuario
}

func NewHandlerUsuario(controller *controllers.ControllerUsuario) (*HandlerUsuarios, error){
	if controller == nil{
		return nil, fmt.Errorf("controlador usuario nulo")
	}
	return &HandlerUsuarios{
		cUsuario: controller,
	}, nil
}

func (hc *HandlerUsuarios) RegistrarUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		body,err := io.ReadAll(r.Body)

		if err != nil {
			http.Error(w, "error en datos de usuario", http.StatusBadRequest)
			return
		}
		_ , err = hc.cUsuario.RegistrarUsuario(body)
		if err != nil {
			http.Error(w, "error registrando usuario nuevo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerUsuarios) BuscarUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		cedula := r.PathValue("cedula")

		if cedula == ""{
			http.Error(w, "no se ingresó cedula", http.StatusBadRequest)
			return
		}
		_ , err := strconv.Atoi(cedula)
		if err != nil{
			http.Error(w, "no se encontró cedula valida", http.StatusBadRequest)
			return
		}
		_ , err = hc.cUsuario.BuscarUsuario(cedula)
		if err != nil {
			http.Error(w, "no se encuentra usuario", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
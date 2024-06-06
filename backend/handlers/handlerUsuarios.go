package handlers

import (
	"fmt"
	"net/http"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/controllers"
	"github.com/gorilla/mux"
)

type HandlerUsuarios struct {
	cUsuario *controllers.ControllerUsuario
}

func NewHandlerUsuario(controller *controllers.ControllerUsuario) (*HandlerUsuarios, error) {
	if controller == nil {
		return nil, fmt.Errorf("controlador usuario nulo")
	}
	return &HandlerUsuarios{
		cUsuario: controller,
	}, nil
}

func (hc *HandlerUsuarios) RegistrarUsuario() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		cedula := vars["cedula"]
		pass := vars["contrasena"]

		usuario, err := hc.cUsuario.RegistrarUsuario(cedula,pass)
		if err != nil {
			http.Error(w, "error registrando usuario nuevo", http.StatusInternalServerError)
			return
		}
		w.Write(usuario)
		w.WriteHeader(http.StatusCreated)
	})
}

func (hc *HandlerUsuarios) LoginHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		vars := mux.Vars(r)
		cedula := vars["cedula"]
		pass := vars["contrasena"]
		if r.Method == http.MethodGet {
			usuario, flag := hc.cUsuario.Authenticate(cedula, pass)
			if flag  {
				w.WriteHeader(http.StatusOK)
				w.Write(usuario)
				return
			}
			http.Error(w, "Credenciales inválidas", http.StatusForbidden)
			return
		}
	})
}


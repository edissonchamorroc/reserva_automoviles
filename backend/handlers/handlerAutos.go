package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/controllers"
	"github.com/gorilla/mux"
)

type HandlerAutos struct {
	cAuto *controllers.ControllerAuto
}

func NewHandlerAutos(controller *controllers.ControllerAuto) (*HandlerAutos, error){
	if controller == nil{
		return nil, fmt.Errorf("controlador autos nulo")
	}
	return &HandlerAutos{
		cAuto: controller,
	}, nil
}

func (hc *HandlerAutos) ListarAutos() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		vars := mux.Vars(r)
		cedula := vars["cedula"]
		autos, err := hc.cAuto.ListarAutos(cedula)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write(nil)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(autos)
	})
}


func (hc *HandlerAutos) ActualizarReserva() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		vars := mux.Vars(r)
		cedula := vars["cedula"]
		id := vars["id"]

		if id == "" {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}
		_, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "no se encontro un id valido", http.StatusBadRequest)
			return
		}

		err = hc.cAuto.ActualizarReserva(id, cedula)
		if err != nil {
			log.Printf("fallo al actualizar la reserva: %s", err.Error())
			http.Error(w, "fallo al actualizar la reserva", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}

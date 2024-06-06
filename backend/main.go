package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/controllers"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/handlers"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/models"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/repositories"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/utils"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)


func main() {

	conn, err := conectarDB(fmt.Sprintf("postgres://%s:%s@db:%s/%s?sslmode=disable", os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME")), "postgres")

	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
	}
	handlerUsuarios := adecuacionAmbienteUsuarios(conn)
	handlerAutos := adecuacionAmbienteAutos(conn)

	r := mux.NewRouter()
	r.Use(utils.CORSMiddleware)
	r.HandleFunc("/login/{cedula}/{contrasena}", handlerUsuarios.LoginHandler()).Methods("GET")
	r.HandleFunc("/registro/{cedula}/{contrasena}", handlerUsuarios.RegistrarUsuario()).Methods("GET")
	r.HandleFunc("/autos/{cedula}", handlerAutos.ListarAutos()).Methods("GET")
	r.HandleFunc("/reserva/{id}/{cedula}", handlerAutos.ActualizarReserva()).Methods("GET")
	http.ListenAndServe(":8080", r)

}

func conectarDB(url, driver string) (*sqlx.DB, error) {
	pgUrl, _ := pq.ParseURL(url)
	db, err := sqlx.Connect(driver, pgUrl)
	if err != nil {
		log.Printf("falló la conexión a base de datos %s", err.Error())
		return nil, err
	}
	log.Printf("Conexion exitosa %#v", db)
	return db, err
}

func adecuacionAmbienteUsuarios(conn *sqlx.DB) *handlers.HandlerUsuarios {

	baseDatosUsuarios, err := repositories.NewRepository[models.Usuario](conn)
	if err != nil {
		log.Fatalln("falló creando instancia en bd de usuarios", err.Error())
	}

	controllerUsuarios, err := controllers.NewControllerUsuario(baseDatosUsuarios)
	if err != nil {
		log.Fatalln("error creando controlador para usuarios", err.Error())

	}
	handlerUsuarios, err := handlers.NewHandlerUsuario(controllerUsuarios)
	if err != nil {
		log.Fatalln("error creando handler para usuarios", err.Error())
	}
	return handlerUsuarios
}

func adecuacionAmbienteAutos(conn *sqlx.DB) *handlers.HandlerAutos {

	baseDatosAutos, err := repositories.NewRepository[models.Auto](conn)
	if err != nil {
		log.Fatalln("falló creando instancia en bd de usuarios", err.Error())
	}

	controllerAutos, err := controllers.NewControllerAuto(baseDatosAutos)
	if err != nil {
		log.Fatalln("error creando controlador para usuarios", err.Error())

	}
	handlerAuto, err := handlers.NewHandlerAutos(controllerAutos)
	if err != nil {
		log.Fatalln("error creando handler para usuarios", err.Error())
	}
	return handlerAuto
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/controllers"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/handlers"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/models"
	"github.com/edissonchamorroc/reserva_automoviles/backend/reserva/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var (
	server   = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	database = "Reserva_Autos"
)

func main() {

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		user, password, server, port, database)

	conn, err := conectarDB(url, "postgres")
	if err != nil {
		log.Fatalln("error conectando a la base de datos", err.Error())
	}

	baseDatosUsuarios, err := repositories.NewRepository[models.Usuario](conn)
	if err != nil {
		log.Fatalln("falló creando instancia en bd", err.Error())
	}

	controllerUsuarios, err := controllers.NewControllerUsuario(baseDatosUsuarios)
	if err != nil {
		log.Fatalln("error creando controlador", err.Error())

	}
	handlerUsuarios, err := handlers.NewHandlerUsuario(controllerUsuarios)
	if err != nil {
		log.Fatalln("error creando handler", err.Error())

	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /usuarios/{cedula}", handlerUsuarios.BuscarUsuario())
	mux.HandleFunc("POST /usuarios/registro", handlerUsuarios.RegistrarUsuario())
	log.Fatal(http.ListenAndServe(":8080", mux))
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
package main

import (
	"log"
	"net/http"
	"os"

	"Cartera-Mongo-Backend/config"
	"Cartera-Mongo-Backend/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "Cartera-Mongo-Backend/docs"

	"github.com/joho/godotenv"
)

// ... (tus anotaciones de Swagger)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se encontró archivo .env en main.go o error al cargarlo.")
	}

	// 1. Conectar a la base de datos MongoDB
	config.ConnectDB()

	// 2. Obtener la colección de MongoDB *después* de que el cliente DB esté conectado
	//    y luego inyectarla en los manejadores.
	projectCol := config.GetCollection()
	handlers.SetProjectCollection(projectCol)

	// Inicializar el router
	r := mux.NewRouter()

	apiV1 := r.PathPrefix("/api/v1").Subrouter()

	// Rutas para Proyectos
	apiV1.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	apiV1.HandleFunc("/projects", handlers.GetProjects).Methods("GET")
	apiV1.HandleFunc("/projects/{id}", handlers.GetProjectByID).Methods("GET")
	apiV1.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	apiV1.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

	// Ruta para la documentación de Swagger
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	listenAddr := ":" + port

	log.Printf("Server running on http://localhost%s", listenAddr)
	log.Printf("Swagger UI available at http://localhost%s/swagger/index.html", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, r))
}
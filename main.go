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

// @title Cartera-Mongo-Backend Project API
// @version 1.0
// @description This is a sample server for a Project API with Go and MongoDB, named Cartera-Mongo-Backend.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host cartera-mongo-backend-production.up.railway.app
// @BasePath /
// @schemes https // ¡IMPORTANTE! Railway usa HTTPS
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env (posiblemente esperado en entorno de despliegue). Asumiendo que las variables de entorno están configuradas.")
	}

	config.ConnectDB()

	projectCol := config.GetCollection()
	handlers.SetProjectCollection(projectCol)

	r := mux.NewRouter()

	// ❌ ELIMINAMOS esta línea: apiV1 := r.PathPrefix("/api/v1").Subrouter()

	// ✅ Y AÑADIMOS las rutas directamente al router principal 'r'
	// O puedes crear un subrouter sin prefijo para mantener la organización si lo deseas:
	// mainRouter := r.PathPrefix("/").Subrouter() // Esto es opcional, si quieres un subrouter para todas las rutas

	// Rutas para Proyectos - Ahora directamente bajo la raíz
	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/funciones/data", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.GetProjectByID).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

	// Ruta para la documentación de Swagger (sigue igual, no tiene prefijo /api/v1)
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
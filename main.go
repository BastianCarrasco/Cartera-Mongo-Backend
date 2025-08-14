package main

import (
	"log"
	"net/http"
	"os"

	"Cartera-Mongo-Backend/config"
	"Cartera-Mongo-Backend/handlers"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/joho/godotenv"
)

// NOTA IMPORTANTE:
// Con go-swagger, las anotaciones globales (@title, @version, @host, etc.)
// se mueven a un archivo swagger.yaml/openapi.yaml separado.
// Las anotaciones para los endpoints individuales (parámetros, respuestas)
// se ponen dentro de las funciones de handler correspondientes.

func main() {
	// Cargar variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env (posiblemente esperado en entorno de despliegue). Asumiendo que las variables de entorno están configuradas.")
	}

	// Leer variables de entorno para la configuración de la base de datos
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")

	// Validar que las variables de la DB estén presentes
	if mongoURI == "" || dbName == "" || collectionName == "" {
		log.Fatal("Error: Las variables de entorno MONGO_URI, DB_NAME o COLLECTION_NAME no están configuradas.")
	}

	// Conectar a la base de datos usando las variables de entorno
	// (Asegúrate de que tu función ConnectDB en config.go acepte estos parámetros)
	// Si tu ConnectDB no acepta parámetros, tendrás que modificarla para que los lea
	// o que ella misma lea las variables de entorno.
	config.ConnectDB() // Modificado para pasar parámetros si ConnectDB los espera

	projectCol := config.GetCollection()
	handlers.SetProjectCollection(projectCol)

	r := mux.NewRouter()

	r.HandleFunc("/projects", handlers.CreateProject).Methods("POST")
	r.HandleFunc("/funciones/data", handlers.GetProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.GetProjectByID).Methods("GET")
	r.HandleFunc("/projects/{id}", handlers.UpdateProject).Methods("PUT")
	r.HandleFunc("/projects/{id}", handlers.DeleteProject).Methods("DELETE")

	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	// Leer el puerto de las variables de entorno
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Valor por defecto para desarrollo local
	}
	listenAddr := ":" + port

	// Leer el host de la aplicación para los logs (puede ser SWAGGER_HOST o RAILWAY_APP_HOST)
	// SWAGGER_HOST de tu .env es una buena opción ya que es la URL de Railway.
	appHost := os.Getenv("SWAGGER_HOST")
	if appHost == "" {
		// Fallback si SWAGGER_HOST no está configurada (por ejemplo, en desarrollo local sin .env)
		log.Println("Advertencia: SWAGGER_HOST no está configurada. Usando 'localhost' para los logs.")
		appHost = "localhost" // Para desarrollo local, podrías querer "localhost:" + port
	}


	log.Printf("Server running on https://%s (listening on port %s)", appHost, port)
	log.Printf("Swagger UI available at https://%s/swagger/index.html", appHost)
	log.Fatal(http.ListenAndServe(listenAddr, r))
}
package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DBClient *mongo.Client
var dbName string // Para almacenar el nombre de la base de datos
var collectionName string // Para almacenar el nombre de la colección

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se encontró archivo .env o error al cargarlo. Intentando leer variables de entorno del sistema.")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("Error: La variable de entorno MONGO_URI no está configurada.")
	}

	dbName = os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "CARTERA" // Valor por defecto si no está en .env
		log.Printf("Advertencia: DB_NAME no configurado en .env, usando '%s' por defecto.", dbName)
	}

	collectionName = os.Getenv("COLLECTION_NAME")
	if collectionName == "" {
		collectionName = "PROYECTOS" // Valor por defecto si no está en .env
		log.Printf("Advertencia: COLLECTION_NAME no configurado en .env, usando '%s' por defecto.", collectionName)
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error al conectar a MongoDB:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error al hacer ping a MongoDB:", err)
	}

	log.Println("Conexión exitosa a MongoDB!")
	DBClient = client
}

// GetCollection devuelve la colección de proyectos.
func GetCollection() *mongo.Collection { // No recibe collectionName ahora
	if DBClient == nil {
		log.Fatal("Error: El cliente de MongoDB no ha sido inicializado. Llama a ConnectDB() primero.")
	}
	return DBClient.Database(dbName).Collection(collectionName) // Usa la variable global
}
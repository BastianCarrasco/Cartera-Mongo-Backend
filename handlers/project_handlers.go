package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Cartera-Mongo-Backend/models"
)

// ProjectCollection será inicializada desde main.go.
// Es importante que sea exportada (empieza con P mayúscula) para que main.go pueda acceder a ella.
var ProjectCollection *mongo.Collection

// SetProjectCollection se usa para inyectar la colección de MongoDB
func SetProjectCollection(collection *mongo.Collection) {
	ProjectCollection = collection
}

// CreateProject godoc
// @Summary Create a new project
// @Description Add a new project to the database
// @Tags Projects
// @Accept json
// @Produce json
// @Param project body models.Project true "Project object to be created"
// @Success 201 {object} models.Project
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects [post]
func CreateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var project models.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ¡CORRECCIÓN AQUÍ! Usar ProjectCollection
	result, err := ProjectCollection.InsertOne(ctx, project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Could not convert InsertedID to ObjectID for project")
	} else {
		project.ID = oid
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(project)
}

// GetProjects godoc
// @Summary Get all projects
// @Description Retrieve a list of all projects
// @Tags Projects
// @Produce json
// @Success 200 {array} models.Project
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects [get]
func GetProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var projects []models.Project
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ¡CORRECCIÓN AQUÍ! Usar ProjectCollection
	cursor, err := ProjectCollection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var project models.Project
		cursor.Decode(&project)
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(projects)
}

// GetProjectByID godoc
// @Summary Get a project by ID
// @Description Retrieve a single project by its ID
// @Tags Projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 200 {object} models.Project
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/{id} [get]
func GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var project models.Project
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ¡CORRECCIÓN AQUÍ! Usar ProjectCollection
	err = ProjectCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(project)
}

// UpdateProject godoc
// @Summary Update an existing project
// @Description Update project details by ID
// @Tags Projects
// @Accept json
// @Produce json
// @Param id path string true "Project ID"
// @Param project body models.Project true "Updated project object"
// @Success 200 {object} models.Project
// @Failure 400 {string} string "Invalid ID or Bad Request"
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/{id} [put]
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var project models.Project
	err = json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	update := bson.D{{Key: "$set", Value: bson.M{
		"nombre":            project.Nombre,
		"academicos":        project.Academicos,
		"estudiantes":       project.Estudiantes,
		"monto":             project.Monto,
		"fecha_postulacion": project.FechaPostulacion,
		"unidad":            project.Unidad,
		"tematica":          project.Tematica,
		"estatus":           project.Estatus,
		"convocatoria":      project.Convocatoria,
		"tipo_convocatoria": project.TipoConvocatoria,
		"inst_conv":         project.InstConv,
		"detalle_apoyo":     project.DetalleApoyo,
		"apoyo":             project.Apoyo,
		"id_kth":            project.IdKth,
		"comentarios":       project.Comentarios,
	}}}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	// ¡CORRECCIÓN AQUÍ! Usar ProjectCollection
	result := ProjectCollection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, opts)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}
		http.Error(w, result.Err().Error(), http.StatusInternalServerError)
		return
	}

	var updatedProject models.Project
	result.Decode(&updatedProject)
	json.NewEncoder(w).Encode(updatedProject)
}

// DeleteProject godoc
// @Summary Delete a project
// @Description Delete a project by ID
// @Tags Projects
// @Produce json
// @Param id path string true "Project ID"
// @Success 204 "No Content"
// @Failure 400 {string} string "Invalid ID"
// @Failure 404 {string} string "Project not found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /projects/{id} [delete]
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ¡CORRECCIÓN AQUÍ! Usar ProjectCollection
	result, err := ProjectCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
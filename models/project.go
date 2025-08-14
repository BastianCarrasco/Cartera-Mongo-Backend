package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Academic representa un académico asociado a un proyecto.
type Academic struct {
	// Puedes añadir los campos específicos que tenga un académico
	// dentro de este objeto, por ejemplo:
	// Name  string `json:"name" bson:"name"`
	// Email string `json:"email" bson:"email"`
	// Otros campos...
}

// Student representa un estudiante asociado a un proyecto.
type Student struct {
	// Puedes añadir los campos específicos que tenga un estudiante
	// dentro de este objeto, por ejemplo:
	// Name  string `json:"name" bson:"name"`
	// Major string `json:"major" bson:"major"`
	// Otros campos...
}

// Project representa un proyecto en la base de datos MongoDB.
type Project struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre           string             `json:"nombre" bson:"nombre"`
	Academicos       []Academic         `json:"academicos" bson:"academicos"`
	Estudiantes      []Student          `json:"estudiantes" bson:"estudiantes"`
	Monto            float64            `json:"monto" bson:"monto"`
	FechaPostulacion time.Time          `json:"fecha_postulacion" bson:"fecha_postulacion"` // Usar time.Time para fechas ISO
	Unidad           string             `json:"unidad" bson:"unidad"`
	Tematica         string             `json:"tematica" bson:"tematica"`
	Estatus          string             `json:"estatus" bson:"estatus"`
	Convocatoria     string             `json:"convocatoria" bson:"convocatoria"`
	TipoConvocatoria string             `json:"tipo_convocatoria" bson:"tipo_convocatoria"`
	InstConv         string             `json:"inst_conv" bson:"inst_conv"`
	DetalleApoyo     string             `json:"detalle_apoyo" bson:"detalle_apoyo"`
	Apoyo            string             `json:"apoyo" bson:"apoyo"`
	IdKth            *string            `json:"id_kth" bson:"id_kth"`             // Usar puntero para manejar 'null'
	Comentarios      string             `json:"comentarios" bson:"comentarios"`
}
package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Academic representa un académico asociado a un proyecto.
type Academic struct {
	// ¡Aquí están los campos que faltaban!
	Nombre    string `json:"nombre" bson:"nombre"`
	APaterno  string `json:"a_paterno" bson:"a_paterno"`
	AMaterno  string `json:"a_materno" bson:"a_materno"`
	// Si hay más campos en tus académicos en la BD, añádelos aquí.
	// Por ejemplo:
	// Email string `json:"email,omitempty" bson:"email,omitempty"`
}

// Student representa un estudiante asociado a un proyecto.
type Student struct {
	Nombre    string `json:"nombre" bson:"nombre"`
	APaterno  string `json:"a_paterno" bson:"a_paterno"`
	AMaterno  string `json:"a_materno" bson:"a_materno"`
}

// Project representa un proyecto en la base de datos MongoDB.
type Project struct {
	ID               primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Nombre           string             `json:"nombre" bson:"nombre"`
	Academicos       []Academic         `json:"academicos" bson:"academicos"`
	Estudiantes      []Student          `json:"estudiantes" bson:"estudiantes"`
	Monto            float64            `json:"monto" bson:"monto"`
	FechaPostulacion time.Time          `json:"fecha_postulacion" bson:"fecha_postulacion"`
	Unidad           string             `json:"unidad" bson:"unidad"`
	Tematica         string             `json:"tematica" bson:"tematica"`
	Estatus          string             `json:"estatus" bson:"estatus"`
	Convocatoria     string             `json:"convocatoria" bson:"convocatoria"`
	TipoConvocatoria string             `json:"tipo_convocatoria" bson:"tipo_convocatoria"`
	InstConv         string             `json:"inst_conv" bson:"inst_conv"`
	DetalleApoyo     string             `json:"detalle_apoyo" bson:"detalle_apoyo"`
	Apoyo            string             `json:"apoyo" bson:"apoyo"`
	IdKth            *string            `json:"id_kth" bson:"id_kth"`
	Comentarios      string             `json:"comentarios" bson:"comentarios"`
}
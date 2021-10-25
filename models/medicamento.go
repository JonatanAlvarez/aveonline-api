package models

import "time"

// medicamento represents data about a record medicament.
type Medicamento struct {
	ID					uint   		`json:"id" gorm:"primary_key;auto_increment"`
	Nombre  		string  	`json:"nombre" gorm:"size:50" binding:"required"`
	Precio 			float64  	`json:"precio"`
	Ubicacion  	string 		`json:"ubicacion" gorm:"size:50" binding:"required"`
	CreatedAt 	time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 	time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt 		*time.Time 		`json:"deleted_at" sql:"index"`
}
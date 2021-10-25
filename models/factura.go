package models

import "time"

// medicamento represents data about a record medicament.
type Factura struct {
	ID						uint   				`json:"id" gorm:"primary_key;auto_increment"`
	Fecha_crear  	time.Time 		`json:"fecha_crear" binding:"required"`
	Pago_total 		float64  			`json:"pago_total"`
	PromocionID		uint					
	Promocion			Promocion			`json:"promocion"`
	Medicamentos 	[]Medicamento	`json:"medicamentos" gorm:"many2many:factura_medicamentos"`
	CreatedAt 		time.Time 		`json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 		time.Time 		`json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt 		*time.Time 		`json:"deleted_at" sql:"index"`
}
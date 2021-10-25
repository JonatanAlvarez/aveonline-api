package models

import "time"

const MAX_PORCENTAJE = 0.7

// medicamento represents data about a record medicament.
type Promocion struct {
	ID						uint   		`json:"id" gorm:"primary_key;auto_increment"`
	Descripcion		string  	`json:"descripcion" gorm:"size:100"`
	Porcentaje 		float64  	`json:"porcentaje"`
	Fecha_inicio  time.Time `json:"fecha_inicio"`
	Fecha_fin  		time.Time `json:"fecha_fin"`
	CreatedAt 		time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt 		time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt 		*time.Time 		`json:"deleted_at" sql:"index"`
}

// CheckDatecollision return true if
func (p *Promocion) CheckDateCollision (promociones []Promocion) bool {
	pStart := p.Fecha_inicio.Unix()
	pEnd := p.Fecha_fin.Unix()

	for _, promocion := range promociones {
		start := promocion.Fecha_inicio.Unix()
		end := promocion.Fecha_fin.Unix()

		if (pStart == start || pStart == end || pEnd == start || pEnd == end) || 
			(pStart >= start && pStart <= end) || (pEnd >= start && pEnd <= end) || 
			(pStart < start && pEnd > end) {
			return true
		}
	}

	return false
}
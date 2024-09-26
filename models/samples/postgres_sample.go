package samples

import "gorm.io/gorm"

type Postgres_Sample struct {
	gorm.Model
	Label string  `json:"label" gorm:"type:varchar(255);not null"`
	Value float64 `json:"value" gorm:"type:numeric;not null"`
}

package samples

import "gorm.io/gorm"

type Postgres_Sample struct {
	gorm.Model
	Value float64 `json:"value" gorm:"type:numeric;not null"`
}

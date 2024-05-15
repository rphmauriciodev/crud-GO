package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	gorm.Model
	Title string
	Year  int16
	Genre string
}

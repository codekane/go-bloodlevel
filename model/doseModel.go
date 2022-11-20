package model

import (
  "gorm.io/gorm"
)

type Dose struct {

  gorm.Model
  Date string `json:"date"`
  Time string `json:"time"`
  Substance string `json:"substance"`
  Dosage int `json:"dosage"`
  Dosage_Unit string `json:dosage_unit"`
}

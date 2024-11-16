package model

import (
  "gorm.io/gorm"
  "time"
)

type Dose struct {
  ID uint `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  DeletedAt gorm.DeletedAt `gorm:"indes" json:"deleted_at"`

  //gorm.Model
  Date string `json:"date"`
  Time string `json:"time"`
  Substance string `json:"substance"`
  Dosage int `json:"dosage"`
  Dosage_Unit string `json:"dosage_unit"`
}

package model

import (
  "gorm.io/gorm"
)


type ROI struct {
  gorm.Model
  Name                          string
}



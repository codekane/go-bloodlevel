package model

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
  "log"
)

func Database() (*gorm.DB, error) {

  db, err := gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

  if err != nil {
    log.Fatal(err.Error())
  }

  if err = db.AutoMigrate(&Dose{}); err != nil {
    log.Println(err)
  }

  return db, err

}

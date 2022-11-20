package main

import (
  "log"
  "github.com/gin-gonic/gin"
  "bloodlevel/dose"
  "bloodlevel/model"
)

func main() {
  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }
  db.DB()

  router := gin.Default()

  router.GET("/doses", dose.GetDoses)
  router.GET("/dose/:id", dose.GetDose)
  router.POST("/dose", dose.PostDose)
  router.PUT("/dose/:id", dose.UpdateDose)
  router.DELETE("/dose/:id", dose.DeleteDose)

  log.Fatal(router.Run(":10000"))
}


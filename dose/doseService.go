package dose

import (
  "net/http"

  "bloodlevel/model"
  "github.com/gin-gonic/gin"
  "strconv"
  "log"
)

type NewDose struct {
  Date string `json:"date" binding:"required"`
  Time string `json:"time" binding:"required"`
  Substance string `json:"substance" binding:"required"`
  Dosage int `json:"dosage" binding:"required"`
  Dosage_Unit string `json:"dosage_unit" binding:"required"`
}

type DoseUpdate struct  {
  Date string `json:"date"`
  Time string `json:"time"`
  Substance string `json:"substance"`
  Dosage int `json:"dosage"`
  Dosage_Unit string `json:"dosage_unit"`
}

func GetDoses(c *gin.Context) {

  var doses []model.Dose

  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }

  if err := db.Find(&doses).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, doses)
}

func GetDose(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var dose model.Dose

  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }

  if err := db.Where("id= ?", id).First(&dose).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Dose not found"})
    return
  }

  c.JSON(http.StatusOK, dose)
}

func PostDose(c *gin.Context) {
  var dose NewDose

  if err := c.ShouldBindJSON(&dose); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  newDose := model.Dose{Date: dose.Date, Time: dose.Time, Substance: dose.Substance, Dosage: dose.Dosage, Dosage_Unit: dose.Dosage_Unit}

  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }

  if err := db.Create(&newDose).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, newDose)

}

func UpdateDose(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var dose model.Dose

  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }

  if err := db.Where("id = ?", id).First(&dose).Error; err !=nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Dose not found!"})
    return
  }

  var updateDose DoseUpdate

  if err := c.ShouldBindJSON(&updateDose); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := db.Model(&dose).Updates(model.Dose{Date: dose.Date, Time: dose.Time, Substance: dose.Substance, Dosage: dose.Dosage, Dosage_Unit: dose.Dosage_Unit}).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }
  c.JSON(http.StatusOK, dose)
}

func DeleteDose(c *gin.Context) {
  id, _ := strconv.Atoi(c.Param("id"))

  var dose model.Dose

  db, err := model.Database()
  if err != nil {
    log.Println(err)
  }

  if err := db.Where("id = ?", id).First(&dose).Error; err != nil{
    c.JSON(http.StatusNotFound, gin.H{"error": "Dose not found!"})
    return
  }

  if err := db.Delete(&dose).Error; err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
    return
  }


  c.JSON(http.StatusOK, gin.H{"message": "Dose deleted"})
}


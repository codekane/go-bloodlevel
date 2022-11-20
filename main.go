package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "strconv"
)

type dose struct {
  ID    int `json:"id"`
  Date  string `json:"date"`
  Time  string `json:"time"`
  Substance string `json:"substance"`
  Dosage int `json:"dosage"`
  Dosage_Unit string `json:"dosage_unit"`
}

var doses = []dose{
  {ID: 1, Date: "2022-11-19", Time: "17:30", Substance: "Dexedrine IR", Dosage: 15, Dosage_Unit: "mg"},
  {ID: 2, Date: "2022-11-19", Time: "18:15", Substance: "Dexedrine IR", Dosage: 10, Dosage_Unit: "mg"},
  {ID: 3, Date: "2022-11-19", Time: "22:15", Substance: "Dexedrine IR", Dosage: 15, Dosage_Unit: "mg"},
}

func main() {
  router := gin.Default()
  router.GET("/doses", getDoses)
  router.GET("/doses/:id", getDoseByID)
  router.POST("/doses", postDoses)

  router.Run("localhost:8080")
}



// getDoses responds with the list of all registered medication doses
func getDoses(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, doses)
}

func postDoses(c *gin.Context) {
  var newDose dose

  // Call BindJSON to bind the received JSON to newDose.
  if err := c.BindJSON(&newDose); err != nil {
    return
  }

  doses = append(doses, newDose)
  c.IndentedJSON(http.StatusCreated, newDose)
}

func getDoseByID(c *gin.Context) {
  id := c.Param("id")

  // Loop over the list of doses, looking for a dose that matches the ID value of the parameter
  for _, a := range doses {
    intID, _ := strconv.Atoi(id)
    if a.ID == intID {
      c.IndentedJSON(http.StatusOK, a)
      return
    }
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "dose not found"})
}

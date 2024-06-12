package routes

import (
	"net/http"
	"pvmsproject/models"

	"github.com/gin-gonic/gin"
)

//Inserting Route
func insertAgency(context *gin.Context) {
	var agency models.Agency

	err := context.ShouldBindJSON(&agency)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could Not Parse Agency Data"})
		return
	}

	_, err = agency.Insert()
	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could Not Insert Agency"})
		return
	}

	context.IndentedJSON(http.StatusCreated, gin.H{"message": "Agency added successfully", "agency": agency})
}

//Update Route
func updateAgency(context *gin.Context) {
	agencyId := context.Param("agencyId")
	var agencyStateName models.Agency
	err := context.ShouldBindJSON(&agencyStateName)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could Not Parse State"})
		return
	}
	agency, err := models.Update(agencyId, agencyStateName.State)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Could not Update the Agency"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"message": "Agency updated successfully", "agency": agency})
}

//Getting Agency By Id Route
func getAgencyById(context *gin.Context) {

	agencyId := context.Param("agencyId")

	agency, err := models.Get(&agencyId)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Could Not Fetch Agency"})
		return
	}
	context.IndentedJSON(http.StatusOK, gin.H{"agency": agency})

}

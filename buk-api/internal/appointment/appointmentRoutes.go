package appointment

import (
	"buk/buk-backend-service/pkg/appointment"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

type appointmentRoutes struct {
	db *sqlx.DB
}

func Routes(r *gin.Engine, json *sqlx.DB) {
	routes := appointmentRoutes{json}

	v1 := r.Group("/v1/appointment")
	{
		//v1.GET("/:email", routes.getAppointment)
		v1.POST("/:email", routes.PostNewAppointment)
	}
}

func (s appointmentRoutes) PostNewAppointment(c *gin.Context) {
	var requestBody appointment.AppointmentDTO
	jsonErr := c.ShouldBindJSON(&requestBody)
	if jsonErr != nil {
		log.Println("Error while biding appointmentRequest", jsonErr)
		c.AbortWithStatusJSON(http.StatusBadRequest, "The request Body has invalid inputs")
		return
	}

	err := CreateNewAppointment(s.db, &requestBody)
	if err != nil {
		c.AbortWithStatusJSON(err.HttpStatus, err.Message)
		return
	}

	c.JSON(http.StatusOK, "Appointment created")
}

//func (s appointmentRoutes) getAppointment(c *gin.Context) {
//	email := c.Param("email")
//
//	account, err := services.GetAccountByEmail(a.json, email)
//	if err != nil {
//		c.AbortWithStatusJSON(err.HttpStatus, err.Message)
//		return
//	}
//
//	c.JSON(http.StatusOK, account)
//}

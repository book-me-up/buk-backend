package appointment

import (
	"buk/buk-backend-service/pkg/appointment"
	"github.com/jmoiron/sqlx"
)

func CreateNewAppointment(db *sqlx.DB, appointmentRequest *appointment.AppointmentDTO) error {
	// Find if exists any appointment within the same time boundries

	// IF NOT save new booking
	err := SaveNewAppointmentByEmail(db, appointmentRequest)

	return err
}

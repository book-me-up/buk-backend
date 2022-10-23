package appointment

import (
	"buk/buk-backend-service/pkg/appointment"
	"github.com/jmoiron/sqlx"
	"log"
)

func SaveNewAppointmentByEmail(db *sqlx.DB, appointment *appointment.AppointmentDTO) error {
	_, err := db.NamedExec(`
		INSERT INTO appointment (partner_id, customer_id, begin_at, end_at, duration, description) 
			VALUES (:PartnerID, :CustomerID, :BeginAt, :EndAt, :Duration, :Description)`, appointment)
	if err != nil {
		log.Println("Error while persisting new appointment", err)
	}

	return nil
}

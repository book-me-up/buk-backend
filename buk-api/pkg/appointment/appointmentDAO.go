package appointment

import (
	"github.com/google/uuid"
	"time"
)

type AppointmentDAO struct {
	Id          uuid.UUID `db:"appointment_id"`
	PartnerID   uuid.UUID `db:"partner_id"`
	CustomerID  uuid.UUID `db:"customer_id"`
	BeginAt     time.Time `db:"begin_at"`
	EndAt       time.Time `db:"end_at"`
	Duration    int       `db:"duration"`
	Description string    `db:"description"`
}

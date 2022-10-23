package appointment

import (
	"github.com/google/uuid"
	"time"
)

type AppointmentDTO struct {
	Id          uuid.UUID
	PartnerID   uuid.UUID
	CustomerID  uuid.UUID
	BeginAt     time.Time
	EndAt       time.Time
	Duration    int
	Description string
}

package repository

import (
	"github.com/lildannylin/bookings-app-golang/internal/models"
	"time"
)

type DatabaseRepo interface {
	AllUsers() bool
	InertReservation(re models.Reservation) (int, error)
	InsertRoomRestriction(r models.RoomRestriction) error
	SearchAvailabilityByDateByRoomID(start, end time.Time, roomID int) (bool, error)
	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	GetRoomByID(id int) (models.Room, error)
}

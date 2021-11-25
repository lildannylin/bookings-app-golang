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
	GetUserByID(id int) (models.User, error)
	UpdateUser(u models.User) error
	Authenticate(email, testPassword string) (int, string, error)
	AllReservation() ([]models.Reservation, error)
}

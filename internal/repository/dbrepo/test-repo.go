package dbrepo

import (
	"github.com/lildannylin/bookings-app-golang/internal/models"
	"time"
)

func (m *testDBRepo) AllUsers() bool {
	return true
}

func (m *testDBRepo) InertReservation(res models.Reservation) (int, error) {

	return 1, nil
}

// InsertRoomRestriction inserts a room restriction to db
func (m *testDBRepo) InsertRoomRestriction(r models.RoomRestriction) error {

	return nil
}

func (m *testDBRepo) SearchAvailabilityByDateByRoomID(start, end time.Time, roomID int) (bool, error) {

	return false, nil
}

func (m *testDBRepo) SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error) {

	var rooms []models.Room

	return rooms, nil
}

// GetRoomByID gets a room by id
func (m *testDBRepo) GetRoomByID(id int) (models.Room, error) {

	var room models.Room

	return room, nil
}

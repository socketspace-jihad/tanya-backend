package school_rooms

import "github.com/socketspace-jihad/tanya-backend/models/schools"

type SchoolRoomsData struct {
	ID                 uint `json:"id"`
	schools.SchoolData `json:"school_data"`
	Name               string `json:"name"`
}

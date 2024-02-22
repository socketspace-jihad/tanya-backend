package event_types

type EventTypesData struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

var (
	SchoolEvents = &EventTypesData{
		ID:   1,
		Name: "School Events",
	}
	StudentEvents = &EventTypesData{
		ID:   2,
		Name: "Student Events",
	}
	ClassEvents = &EventTypesData{
		ID:   3,
		Name: "Class Events",
	}
	SchoolBatchEvents = &EventTypesData{
		ID:   4,
		Name: "School Batch Events",
	}
)

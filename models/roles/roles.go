package roles

const (
	GuruRolesID         = 1
	SiswaRolesID        = 2
	OrangTuaRolesID     = 3
	AdminSekolahRolesID = 4
)

type RolesData struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

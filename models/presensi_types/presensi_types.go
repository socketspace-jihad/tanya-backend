package presensi_types

type PresensitypesData struct {
	ID        uint8  `json:"id"`
	Name      string `json:"name"`
	Deskripsi string `json:"deskripsi"`
	Actor     string `json:"Actor"`
	Tools     string `json:"tools"`
	Scope     string `json:"scope"`
}

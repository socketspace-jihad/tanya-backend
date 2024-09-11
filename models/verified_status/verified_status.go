package verified_status

var (
	Unverified = VerifiedStatus{
		ID:   1,
		Name: "Unverified",
	}
	Verified = VerifiedStatus{
		ID:   2,
		Name: "Verified",
	}
)

type VerifiedStatus struct {
	ID   uint8  `json:"id"`
	Name string `json:"name"`
}

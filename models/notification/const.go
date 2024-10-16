package notification

import "encoding/json"

var (
	CatatanPersonalTargetPath string = "/catatan-personal"
)

func ParseDataStructToString(data interface{}) string {
	res, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(res)
}

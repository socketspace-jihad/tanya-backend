package get_all_schools

import (
	"encoding/json"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/models/schools"
)

type Schools struct{}

func (s *Schools) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data, err := schools.SchoolsDB.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(data)
}

func init() {
	http.DefaultServeMux.Handle("/v1/master-data/schools", &Schools{})
}

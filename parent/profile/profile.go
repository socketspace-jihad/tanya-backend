package profile

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/parent_profiles"
)

type ProfileParent struct{}

func (p *ProfileParent) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	parent, err := middlewares.GetParentFromRequestContext(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Println("PARENT ID", parent.ID)
	profile, err := parent_profiles.ParentProfilesDB.GetByID(parent.ID)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func init() {
	http.DefaultServeMux.HandleFunc("/v1/parent/profile", middlewares.ParentMiddleware(&ProfileParent{}))
}

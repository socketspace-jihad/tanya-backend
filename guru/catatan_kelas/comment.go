package catatan_kelas

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_comments"
)

type CatatanKelasComments struct{}

func (c *CatatanKelasComments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teacher, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	switch r.Method {
	case http.MethodPost:
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		comment := school_class_events_notes_comments.SchoolClassEventsNotesCommentsData{}
		if err := json.Unmarshal(body, &comment); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		comment.TeacherProfilesData = teacher
		if err := school_class_events_notes_comments.SchoolClassEventsNotesCommentsDB.Save(&comment); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func init() {
	http.DefaultServeMux.Handle("/v1/guru/catatan-kelas/comment", middlewares.TeacherMiddleware(&CatatanKelasComments{}))
}

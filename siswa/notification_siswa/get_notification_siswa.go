package notification_siswa

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
)

type NotificationSiswa struct{}

func (n *NotificationSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	student, err := middlewares.GetStudentFromRequestContext(r)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("STUDENT NOTIFICATION", student.ID, student.UserRolesData.UserData.ID)
	switch r.Method {
	case http.MethodGet:
		notifications, err := notification.NotificationDB.GetByUserOrStudentProfilesID(
			student.ID,
			student.UserData.ID,
		)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notifications)
	}
}

func init() {
	http.DefaultServeMux.HandleFunc(
		"/v1/student/notification",
		middlewares.StudentMiddleware((&NotificationSiswa{}).ServeHTTP),
	)
}

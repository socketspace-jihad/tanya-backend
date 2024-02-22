package guru

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/student_tasks"
	"github.com/socketspace-jihad/tanya-backend/models/teacher_profiles"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
	_ "github.com/socketspace-jihad/tanya-backend/queue/engine/nats"
)

type Tugas struct {
	q engine.QueueEngine
}

func (p *Tugas) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teacher, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data := student_tasks.StudentTasksData{
		TeacherProfilesData: *teacher,
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := student_tasks.StudentTasksDB.Save(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	teacher, err = teacher_profiles.TeacherProfilesDB.GetByID(data.TeacherProfilesData.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	data.TeacherProfilesData = *teacher
	p.q.Publish(events.PembuatanTugas{
		Message:   fmt.Sprintf("Ada tugas dari %v, Tekan notif ini untuk melihat", data.TeacherProfilesData.Name),
		SubjectID: 10,
		StudentID: data.StudentProfilesData.ID,
	}, queue.TSekolahPembuatanTugas)
}

func init() {
	platformFactory, err := engine.GetQueueEngine(os.Getenv("PEMBUATAN_TUGAS_QUEUE_ENGINE"))
	if err != nil {
		panic(err)
	}
	platform := platformFactory()
	if err := platform.Connect(&engine.EngineAuthData{
		Host:     os.Getenv("PEMBUATAN_TUGAS_QUEUE_HOST"),
		Username: os.Getenv("PEMBUATAN_TUGAS_QUEUE_USERNAME"),
		Password: os.Getenv("PEMBUATAN_TUGAS_QUEUE_PASSWORD"),
	}); err != nil {
		panic(err)
	}
	http.DefaultServeMux.Handle("/v1/guru/tugas", middlewares.TeacherMiddleware(&Tugas{
		q: platform,
	}))
}

package presensi_siswa

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

type PresensiSiswa struct {
	student_presensi.StudentPresensiData `json:"student_presensi"`
	q                                    engine.QueueEngine
}

func (p *PresensiSiswa) CreatePresensi() error {
	return student_presensi.StudentPresensiDB.Save(&p.StudentPresensiData)
}

func (p *PresensiSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		student, err := middlewares.GetStudentFromRequestContext(r)
		if err != nil {
			return
		}
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Println(p.StudentPresensiData)
		p.StudentPresensiData.StudentProfilesData = student
		err = p.CreatePresensi()
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		p.q.Publish(p.StudentPresensiData, queue.TSekolahPresensiSiswa)

	}

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
	http.DefaultServeMux.HandleFunc(
		"/v1/siswa/presensi",
		middlewares.StudentMiddleware((&PresensiSiswa{
			q: platform,
		}).ServeHTTP),
	)
}

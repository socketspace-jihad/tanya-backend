package guru

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_pictures"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
	_ "github.com/socketspace-jihad/tanya-backend/queue/engine/nats"
)

type CatatanKelas struct {
	q engine.QueueEngine
}

func (p *CatatanKelas) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teacher, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set batas ukuran maksimal file upload (misalnya, 10MB)
	const maxUploadSize = 10 << 20 // 10MB
	r.ParseMultipartForm(maxUploadSize)

	// Parsing JSON data dari form
	var eventNotes school_class_events_notes.SchoolClassEventsNotesData
	eventNotes.TeacherProfilesData = *teacher

	if err := json.Unmarshal([]byte(r.FormValue("data")), &eventNotes); err != nil {
		log.Println("ERROR DECODE!", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Mendapatkan file dari multipart form
	files := r.MultipartForm.File["files[]"]

	// Simpan data ke dalam database
	if err := school_class_events_notes.SchoolClassEventsNotesDB.Save(&eventNotes); err != nil {
		log.Println("ERROR SAVE!", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(files) > 0 {
		// Menyimpan setiap file
		for _, fileHeader := range files {
			go func(fileHeader *multipart.FileHeader) {
				if fileHeader != nil {
					// Buka file
					file, err := fileHeader.Open()
					if err != nil {
						log.Println("ERROR OPEN FILE!", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					defer file.Close()

					// Tentukan path untuk menyimpan file
					folderPath := fmt.Sprintf("./assets/images/class-events/%d", eventNotes.SchoolClassEventsData.ID)
					filePath := fmt.Sprintf("%s/%s", folderPath, fileHeader.Filename)

					// Buat folder jika belum ada
					if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
						log.Println("ERROR CREATE DIRECTORY!", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					// Simpan file ke folder
					dst, err := os.Create(filePath)
					if err != nil {
						log.Println("ERROR CREATE FILE!", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					defer dst.Close()

					if _, err := io.Copy(dst, file); err != nil {
						log.Println("ERROR SAVE FILE!", err.Error())
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}

					if err := school_class_events_notes_pictures.SchoolClassEventsNotesPicturesDB.Save(&school_class_events_notes_pictures.SchoolClassEventsNotesPicturesData{
						Path:                       filePath[1:],
						SchoolClassEventsNotesData: eventNotes,
					}); err != nil {
						log.Println(err)
					}
				}
			}(fileHeader)
		}
	}

	// Encode data untuk respons
	json.NewEncoder(w).Encode(eventNotes)

	// Publish event setelah berhasil
	p.q.Publish(events.EventKelasData{
		Title:   fmt.Sprintf("Ada catatan kelas dari %v, Tekan notif ini untuk melihat", eventNotes.TeacherProfilesData.Name),
		ClassID: eventNotes.SchoolClassData.ID,
	}, queue.TEventKelas)
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
	http.DefaultServeMux.Handle("/v1/guru/catatan-kelas", middlewares.TeacherMiddleware(&CatatanKelas{
		q: platform,
	}))
}

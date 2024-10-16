package catatan_personal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_pictures"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

type CatatanPersonal struct {
	q engine.QueueEngine
}

func (c *CatatanPersonal) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	teacher, err := middlewares.GetTeacherFromRequestContext(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		data := r.URL.Query().Get("class_events_id")
		if data == "" {
			return
		}
		classEventsID, err := strconv.Atoi(data)
		if err != nil {
			return
		}
		notes, err := school_class_events_notes_personal.SchoolClassEventsNotesPersonalDB.GetByTeacherAndClassEventsID(uint(classEventsID), teacher.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(notes)
	case http.MethodPost:
		// Set batas ukuran maksimal file upload (misalnya, 10MB)
		const maxUploadSize = 10 << 20 // 10MB
		r.ParseMultipartForm(maxUploadSize)

		// Parsing JSON data dari form
		var eventNotes school_class_events_notes_personal.SchoolClassEventsNotesPersonalData
		eventNotes.TeacherProfilesData = *teacher

		if err := json.Unmarshal([]byte(r.FormValue("data")), &eventNotes); err != nil {
			log.Println("ERROR DECODE!", err.Error())
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Mendapatkan file dari multipart form
		files := r.MultipartForm.File["files[]"]

		// Simpan data ke dalam database
		if err := school_class_events_notes_personal.SchoolClassEventsNotesPersonalDB.Save(&eventNotes); err != nil {
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
						folderPath := fmt.Sprintf("./assets/images/class-events-personal/%d", eventNotes.SchoolClassEventsData.ID)
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

						if err := school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesDB.Save(&school_class_events_notes_personal_pictures.SchoolClassEventsNotesPersonalPicturesData{
							Path:                               filePath[1:],
							SchoolClassEventsNotesPersonalData: eventNotes,
						}); err != nil {
							log.Println(err)
						}
					}
				}(fileHeader)
			}
		}

		// Encode data untuk respons
		json.NewEncoder(w).Encode(eventNotes)

		notifData := notification.ParseDataStructToString(eventNotes)
		// Publish event setelah berhasil
		c.q.Publish(events.EventSiswaData{
			Title:     fmt.Sprintf("Ada catatan kelas KHUSUS untuk SISWA Bapak/Ibu dari %v, Tekan notif ini untuk melihat", eventNotes.TeacherProfilesData.Name),
			StudentID: eventNotes.StudentProfilesData.ID,
			NotificationData: notification.NotificationData{
				Title:      "Catatan Khusus dari Guru",
				Contents:   "Ada catatan khusus dari guru",
				TargetPath: &notification.CatatanPersonalTargetPath,
				Data:       &notifData,
			},
		}, queue.TEventSiswa)
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
	http.DefaultServeMux.Handle("/v1/guru/catatan-personal", middlewares.TeacherMiddleware(&CatatanPersonal{
		q: platform,
	}))
}

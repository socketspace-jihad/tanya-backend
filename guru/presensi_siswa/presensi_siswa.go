package presensi_siswa

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/socketspace-jihad/tanya-backend/middlewares"
	"github.com/socketspace-jihad/tanya-backend/models/event_types"
	"github.com/socketspace-jihad/tanya-backend/models/notification"
	"github.com/socketspace-jihad/tanya-backend/models/presensi_types"
	"github.com/socketspace-jihad/tanya-backend/models/school_class_events"
	"github.com/socketspace-jihad/tanya-backend/models/student_presensi"
	"github.com/socketspace-jihad/tanya-backend/models/student_profiles"
	"github.com/socketspace-jihad/tanya-backend/queue"
	"github.com/socketspace-jihad/tanya-backend/queue/consumers/events"
	"github.com/socketspace-jihad/tanya-backend/queue/engine"
)

type PresensiSiswaV1 struct {
	Profiles                                  []student_profiles.StudentProfilesData `json:"student_profiles"`
	school_class_events.SchoolClassEventsData `json:"school_class_events"`
	q                                         engine.QueueEngine
}

func (p *PresensiSiswaV1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		g, err := middlewares.GetTeacherFromRequestContext(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		wg := &sync.WaitGroup{}
		for _, profile := range p.Profiles {
			wg.Add(1)
			go func(profile student_profiles.StudentProfilesData) {
				defer wg.Done()
				presensi := &student_presensi.StudentPresensiData{
					StudentProfilesData:   &profile,
					SchoolClassEventsData: &p.SchoolClassEventsData,
					EventTypesData:        *event_types.ClassEvents,
					TeacherProfilesData:   g,
					PresensitypesData: &presensi_types.PresensitypesData{
						ID: 4,
					},
				}
				err := student_presensi.StudentPresensiDB.Save(presensi)
				if err == nil {
					notifData := notification.ParseDataStructToString(presensi)
					p.q.Publish(events.EventSiswaData{
						StudentID: profile.ID,
						Title:     fmt.Sprintf("%v HADIR pada kegiatan %v", profile.Name, *p.SchoolClassEventsData.Name),
						Subtitle:  fmt.Sprintf("Guru a/n %v telah melakukan Presensi Kegiatan, ketuk untuk lebih detail", g.Name),
						NotificationData: notification.NotificationData{
							Title:      fmt.Sprintf("%v baru saja melakukan presensi kegiatan"),
							Contents:   fmt.Sprintf("%v HADIR pada kegiatan tersebut"),
							TargetPath: &notification.PresensiSiswaTargetPath,
							Data:       &notifData,
						},
					}, queue.TEventSiswa)
				}
			}(profile)
		}
		wg.Wait()
		w.Header().Set("Content-Type", "application/json")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

type PresensiSiswa struct {
	Profiles                                  []student_profiles.StudentProfilesData `json:"student_profiles"`
	school_class_events.SchoolClassEventsData `json:"school_class_events"`
	q                                         engine.QueueEngine
	Hadir                                     []student_profiles.StudentProfilesData `json:"hadir"`
	TidakHadir                                []student_profiles.StudentProfilesData `json:"tidak_hadir"`
	PresensiTypes                             presensi_types.PresensitypesData       `json:"presensi_types"`
}

func (p *PresensiSiswa) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		g, err := middlewares.GetTeacherFromRequestContext(r)
		if err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		body, err := io.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.Unmarshal(body, p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		wg := &sync.WaitGroup{}
		for _, profile := range p.Hadir {
			wg.Add(1)
			go func(profile student_profiles.StudentProfilesData) {
				defer wg.Done()
				presensi := &student_presensi.StudentPresensiData{
					StudentProfilesData:   &profile,
					SchoolClassEventsData: &p.SchoolClassEventsData,
					EventTypesData:        *event_types.ClassEvents,
					TeacherProfilesData:   g,
					Hadir:                 1,
					PresensitypesData:     &p.PresensiTypes,
				}
				err := student_presensi.StudentPresensiDB.Save(presensi)
				if err == nil {
					notifData := notification.ParseDataStructToString(presensi)
					p.q.Publish(events.EventSiswaData{
						StudentID: profile.ID,
						Title:     fmt.Sprintf("%v HADIR pada kegiatan %v", profile.Name, *p.SchoolClassEventsData.Name),
						Subtitle:  fmt.Sprintf("Guru a/n %v telah melakukan Presensi Kegiatan, ketuk untuk lebih detail", g.Name),
						NotificationData: notification.NotificationData{
							Title:      fmt.Sprintf("%v baru saja melakukan presensi kegiatan"),
							Contents:   fmt.Sprintf("%v HADIR pada kegiatan tersebut"),
							TargetPath: &notification.PresensiSiswaTargetPath,
							Data:       &notifData,
						},
					}, queue.TEventSiswa)
				}
			}(profile)
		}
		for _, profile := range p.TidakHadir {
			wg.Add(1)
			go func(profile student_profiles.StudentProfilesData) {
				defer wg.Done()
				presensi := &student_presensi.StudentPresensiData{
					StudentProfilesData:   &profile,
					SchoolClassEventsData: &p.SchoolClassEventsData,
					EventTypesData:        *event_types.ClassEvents,
					TeacherProfilesData:   g,
					Hadir:                 0,
					PresensitypesData:     &p.PresensiTypes,
				}
				err := student_presensi.StudentPresensiDB.Save(presensi)
				if err == nil {
					notifData := notification.ParseDataStructToString(presensi)
					p.q.Publish(events.EventSiswaData{
						StudentID: profile.ID,
						Title:     fmt.Sprintf("%v TIDAK HADIR pada kegiatan %v", profile.Name, *p.SchoolClassEventsData.Name),
						Subtitle:  fmt.Sprintf("Guru a/n %v telah melakukan Presensi Kegiatan, ketuk untuk lebih detail", g.Name),
						NotificationData: notification.NotificationData{
							Title:      fmt.Sprintf("%v baru saja melakukan presensi kegiatan"),
							Contents:   fmt.Sprintf("%v TIDAK HADIR pada kegiatan tersebut"),
							TargetPath: &notification.PresensiSiswaTargetPath,
							Data:       &notifData,
						},
					}, queue.TEventSiswa)
				}
			}(profile)
		}
		wg.Wait()
		w.Header().Set("Content-Type", "application/json")
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
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
	http.DefaultServeMux.HandleFunc("/v1/guru/presensi/siswa", middlewares.TeacherMiddleware(&PresensiSiswaV1{
		q: platform,
	}))
	http.DefaultServeMux.HandleFunc("/v2/guru/presensi/siswa", middlewares.TeacherMiddleware(&PresensiSiswa{
		q: platform,
	}))
}

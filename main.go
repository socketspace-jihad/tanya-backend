package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	_ "github.com/socketspace-jihad/tanya-backend/auth"
	_ "github.com/socketspace-jihad/tanya-backend/class/event_presensi"
	_ "github.com/socketspace-jihad/tanya-backend/class/events_notes"
	_ "github.com/socketspace-jihad/tanya-backend/class/student_class"
	_ "github.com/socketspace-jihad/tanya-backend/guru"
	_ "github.com/socketspace-jihad/tanya-backend/guru/catatan_kelas"
	_ "github.com/socketspace-jihad/tanya-backend/guru/catatan_personal"
	_ "github.com/socketspace-jihad/tanya-backend/guru/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/guru/profile_guru"
	_ "github.com/socketspace-jihad/tanya-backend/guru/registrasi_guru"
	_ "github.com/socketspace-jihad/tanya-backend/guru/schedule_guru"
	_ "github.com/socketspace-jihad/tanya-backend/master_data/assets"
	_ "github.com/socketspace-jihad/tanya-backend/master_data/class_events_notes_pictures"
	_ "github.com/socketspace-jihad/tanya-backend/master_data/schools/get_all_schools"
	_ "github.com/socketspace-jihad/tanya-backend/models/class_tasks/class_tasks_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/class_tasks_student/class_tasks_student_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/class_tasks_student_parent/class_tasks_student_parent_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/global_chats/global_chats_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/global_chats_detail/global_chats_detail_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/notification/notification_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/parent_profiles/parent_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/parent_student/parent_student_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/person_chats/person_chats_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events/school_class_events_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes/school_class_events_notes_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_comments/school_class_events_notes_comments_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal/school_class_events_notes_personal_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_personal_pictures/school_class_events_notes_personal_pictures_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_pictures/school_class_events_notes_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes_viewer/school_class_events_notes_viewer_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/schools/schools_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_events/student_events_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_presensi/student_presensi_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_profiles/student_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_tasks/student_tasks_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/teacher_profiles/teacher_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/user/user_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/user_roles/user_roles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/user_topics/user_topics_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/notification/fcm"
	_ "github.com/socketspace-jihad/tanya-backend/parent/catatan_kelas"
	_ "github.com/socketspace-jihad/tanya-backend/parent/catatan_personal"
	_ "github.com/socketspace-jihad/tanya-backend/parent/kegiatan_siswa_parent"
	_ "github.com/socketspace-jihad/tanya-backend/parent/registrasi_parent"
	_ "github.com/socketspace-jihad/tanya-backend/parent/schedule_parent"
	_ "github.com/socketspace-jihad/tanya-backend/parent/siswa_parent"
	_ "github.com/socketspace-jihad/tanya-backend/parent/tugas_parent"

	// _ "github.com/socketspace-jihad/tanya-backend/queue/consumers/guru/event_kelas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/event_kelas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/event_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/pembuatan_tugas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/presensi_siswa"

	// _ "github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa/event_kelas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa/pembuatan_tugas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/notification_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/profile_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/registrasi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/schedule_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/tugas_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/update_user_role_session_profile"
	_ "github.com/socketspace-jihad/tanya-backend/user/chats/get_chats"
	_ "github.com/socketspace-jihad/tanya-backend/user/chats/get_detail_chats"
	_ "github.com/socketspace-jihad/tanya-backend/user/chats/send_chats"
	_ "github.com/socketspace-jihad/tanya-backend/user/topics"
)

var proc atomic.Int32

func main() {
	proc.Store(0)
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		log.Println("LISTENING ON PORT 8082")
		if err := http.ListenAndServe(":8082", nil); err != nil {
			log.Fatalln(err)
		}
	}()
	<-exit
}

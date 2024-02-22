package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/socketspace-jihad/tanya-backend/auth"
	_ "github.com/socketspace-jihad/tanya-backend/class/events_notes"
	_ "github.com/socketspace-jihad/tanya-backend/guru"
	_ "github.com/socketspace-jihad/tanya-backend/guru/schedule_guru"
	_ "github.com/socketspace-jihad/tanya-backend/models/parent_profiles/parent_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/parent_student/parent_student_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events/school_class_events_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/school_class_events_notes/school_class_events_notes_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_events/student_events_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_presensi/student_presensi_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_profiles/student_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/student_tasks/student_tasks_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/teacher_profiles/teacher_profiles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/user/user_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/models/user_roles/user_roles_mysql"
	_ "github.com/socketspace-jihad/tanya-backend/notification/fcm"
	_ "github.com/socketspace-jihad/tanya-backend/parent/schedule_parent"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/pembuatan_tugas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/orang_tua/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa/pembuatan_tugas"
	_ "github.com/socketspace-jihad/tanya-backend/queue/consumers/siswa/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/presensi_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/profile_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/schedule_siswa"
	_ "github.com/socketspace-jihad/tanya-backend/siswa/tugas_siswa"
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

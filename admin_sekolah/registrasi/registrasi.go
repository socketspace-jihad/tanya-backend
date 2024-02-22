package registrasi

import "net/http"

type Registrasi struct {
}

// 1. Registrasi orang tua dan siswa oleh sekolah
// 2. Registrasi orang tua individu
// 3. Registrasi siswa individu

func (reg *Registrasi) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//
}

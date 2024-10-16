package assets

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ImagePersonalHandler struct{}

func (i *ImagePersonalHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan URL Path
	path := strings.TrimPrefix(r.URL.Path, "/assets/images/class-events-personal/")
	pathParts := strings.Split(path, "/")

	// Pastikan path memiliki format yang benar
	if len(pathParts) != 2 {
		http.Error(w, "Invalid image path", http.StatusBadRequest)
		return
	}

	idEvent := pathParts[0]
	imageName := pathParts[1]

	// Validasi bahwa nama file adalah gambar
	if !strings.HasSuffix(imageName, ".jpg") && !strings.HasSuffix(imageName, ".jpeg") && !strings.HasSuffix(imageName, ".png") {
		http.Error(w, "Invalid file type", http.StatusBadRequest)
		return
	}

	// Menentukan path lengkap ke file gambar
	imagePath := filepath.Join("assets", "images", "class-events", idEvent, imageName)

	// Membuka file gambar
	file, err := os.Open(imagePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	// Menentukan header Content-Type
	http.ServeFile(w, r, imagePath)
}

func init() {
	http.DefaultServeMux.HandleFunc("/assets/images/class-events-personal/", (&ImagePersonalHandler{}).ServeHTTP)
}

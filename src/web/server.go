package web

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/image", imageViewer)

	if err := http.ListenAndServe(":80", mux); err != nil {
		log.Fatalf("Error: %v", err)
	}
}

func imageViewer(w http.ResponseWriter, r *http.Request) {
	path, err := os.Getwd()
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Error Occured."))
		return
	}

	fileName := "canvas"
	if v, ok := r.URL.Query()["name"]; ok && len(v) > 0 {
		fileName = v[0]
	}

	file, err := os.ReadFile(path + "/images/" + fileName + ".png")
	if err != nil {
		w.WriteHeader(404)
		w.Write([]byte("Error Occured."))
		return
	}

	w.Header().Add("Content-Type", "image/png")
	w.Write(file)
}

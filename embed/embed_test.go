package embed_test

import (
	"embed"
	"log"
	"net/http"
	"testing"
)

func TestEmbed(t *testing.T) {
	var content embed.FS
	mutex := http.NewServeMux()
	mutex.Handle("/", http.FileServer(http.FS(content)))
	err := http.ListenAndServe(":8080", mutex)
	if err != nil {
		log.Fatal(err)
	}
}

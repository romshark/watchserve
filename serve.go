package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"

	_ "embed"

	"github.com/fsnotify/fsnotify"
	sse "github.com/r3labs/sse/v2"
)

//go:embed frame.html
var frameHTML []byte

func main() {
	fFilePath := flag.String("f", "", "file path")
	fHost := flag.String("host", ":8080", "host address")
	flag.Parse()

	if *fFilePath == "" {
		log.Fatal("missing file path, use the -f flag")
	}
	fileContents, err := os.ReadFile(*fFilePath)
	if err != nil {
		log.Fatal("reading file:", err)
	}

	sseSrv := sse.New()
	const streamUpdates = "updates"
	sseSrv.CreateStream(streamUpdates)
	defer sseSrv.Close()

	go watchFile(*fFilePath, sseSrv, "updates")
	listenHTTP(*fHost, *fFilePath, fileContents, sseSrv)
}

func watchFile(
	filePath string,
	sseSrv *sse.Server,
	streamUpdates string,
) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("initializing file watcher:", err)
	}
	defer watcher.Close()
	if err := watcher.Add(filePath); err != nil {
		log.Fatal("setting up file watcher:", err)
	}

	var updateMsgWrite = []byte("write")
	for {
		select {
		case e := <-watcher.Events:
			if e.Name != filePath || e.Op == fsnotify.Write {
				continue
			}
			log.Print("file change detected")
			sseSrv.Publish(streamUpdates, &sse.Event{
				Data: updateMsgWrite,
			})
		case err := <-watcher.Errors:
			log.Fatal("watching file:", err)
		}
	}
}

func listenHTTP(
	host string,
	filePath string,
	fileContents []byte,
	sseSrv *sse.Server,
) {
	const (
		pathRoot = "/"
		pathMeta = "/meta"
		pathFile = "/file"
		pathSSE  = "/events"
	)

	mux := http.NewServeMux()
	mux.HandleFunc(pathRoot, func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write(frameHTML); err != nil {
			log.Fatal("writing frame HTML:", err)
		}
	})
	mux.HandleFunc(pathFile, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache")
		http.ServeFile(w, r, filePath)
	})
	mux.HandleFunc(pathMeta, func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(struct {
			FilePath string `json:"file-path"`
		}{
			FilePath: filePath,
		}); err != nil {
			log.Fatal("writing meta JSON:", err)
		}
	})
	mux.HandleFunc(pathSSE, sseSrv.HTTPHandler)

	log.Printf("listening on %s", host)
	if err := http.ListenAndServe(host, mux); err != nil {
		log.Fatal("serving HTTP:", err)
	}
	log.Fatal("closed")
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

var (
	basePath string
	host     = "127.0.0.1"
	port     = "5000"
)

func init() {
	flag.StringVar(&basePath, "path", basePath, "Path where assets are placed")
	flag.StringVar(&host, "host", host, "Host where to bind server")
	flag.StringVar(&port, "port", port, "Port where to bind server")
	flag.Parse()
}

func main() {

	if len(basePath) == 0 {
		fmt.Println("No base path provided")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	connString := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("Running test server at %s\n", connString)
	fmt.Printf("Serving test data from %s\n", basePath)

	http.ListenAndServe(connString, mux)
}

func handler(w http.ResponseWriter, req *http.Request) {
	log.Printf("Requested URI %s\n", req.RequestURI)
	plugin, asset := extractParts(req)
	if plugin == "" { // list plugins
		files, err := ioutil.ReadDir(basePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		plugs := []Asset{}
		for _, file := range files {
			if file.IsDir() {
				plug := Asset{
					Name:      file.Name(),
					Size:      file.Size(),
					ModTime:   file.ModTime(),
					Mode:      file.Mode(),
					IsDir:     file.IsDir(),
					URL:       "./" + file.Name() + "/",
					IsSymlink: false,
				}
				plugs = append(plugs, plug)
			}
		}
		data, err := json.Marshal(plugs)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(data)
		return
	} else if asset == "" { // list assets
		folders, err := ioutil.ReadDir(path.Join(basePath, plugin))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// assets := Assets{}
		assets := []Asset{}
		for _, folder := range folders {
			if !folder.IsDir() {
				asset := Asset{
					Name:      folder.Name(),
					Size:      folder.Size(),
					ModTime:   folder.ModTime(),
					Mode:      folder.Mode(),
					IsDir:     folder.IsDir(),
					URL:       "./" + folder.Name() + "/",
					IsSymlink: false,
				}
				assets = append(assets, asset)
			}
		}
		data, err := json.Marshal(assets)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(data)
		return
	} else { // return asset
		file, err := os.Open(path.Join(basePath, plugin, asset))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		data, err := ioutil.ReadAll(file)
		w.Header().Set("Content-Type", "application/octet-stream")
		w.WriteHeader(200)
		w.Write(data)
		return
	}
}

func extractParts(req *http.Request) (string, string) {
	buff := strings.Split(req.URL.String(), "/")
	parts := []string{}
	for _, p := range buff {
		if p != "" {
			parts = append(parts, p)
		}
	}
	if len(parts) == 0 {
		return "", ""
	} else if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

type Asset struct {
	Name      string      `json:"Name"`
	Size      int64       `json:"Size"`
	URL       string      `json:"URL"`
	ModTime   time.Time   `json:"ModTime"`
	Mode      os.FileMode `json:"Mode"`
	IsDir     bool        `json:"IsDir"`
	IsSymlink bool        `json:"IsSymlink"`
}

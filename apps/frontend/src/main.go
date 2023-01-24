package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

// compile all templates and cache them
var templates = template.Must(template.ParseGlob("templates/*"))

var backend = GetEnvironmentVariable("BACKEND_URL", "backend:8080")

type Data struct {
	Title string // Must be exported!
	Body  string // Must be exported!
}

// Renders the templates
func renderTemplate(w http.ResponseWriter, tmpl string, page Data) {
	err := templates.ExecuteTemplate(w, tmpl, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetTitleAndBody(backend string, pageId string) Data {
	// Get Title and Body information from backend REST API
	resp, err := http.Get("http://" + backend + "/pages/" + pageId)
	if err != nil {
		log.Fatalln(err)
	}
	// parse the response body into a struct
	var respBody Data
	json.NewDecoder(resp.Body).Decode(&respBody)
	return Data{Title: respBody.Title, Body: respBody.Body}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", GetTitleAndBody(backend, "index"))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", GetTitleAndBody(backend, "about"))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", GetTitleAndBody(backend, "test3"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"alive": true}`))
}

func main() {
	httpPort := GetEnvironmentVariable("HTTP_PORT", "8080")
	httpAddress := GetEnvironmentVariable("HTTP_ADDRESS", "0.0.0.0")

	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/about/", AboutHandler)
	http.HandleFunc("/test/", TestHandler)
	http.HandleFunc("/health", HealthHandler)
	log.Fatal(http.ListenAndServe(httpAddress+":"+httpPort, nil))
}

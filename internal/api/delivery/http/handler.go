package http

import (
	"ascii-art-web/internal/api/usecase"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

type Handler struct{}

type Data struct {
	Text string
}

var data Data

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) GetPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/get.html", "./templates/500.html", "./templates/not_found.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "500.html", nil)
		return
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "not_found.html", nil)
		return
	}
	if r.Method != "GET" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("BAD REQUEST: Your method will be GET"))
		return
	}
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, nil)
}

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("BAD REQUEST: Your method will be POST"))
		return
	}
	tmpl, err := template.ParseFiles("./templates/post.html", "./templates/500.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		tmpl.ExecuteTemplate(w, "500.html", nil)
		return
	}
	r.ParseForm()
	font := r.Form.Get("fonts")
	text := r.Form.Get("text")
	ascii_art := usecase.ConvertToAscii(text, font)
	data.Text = ascii_art
	w.WriteHeader(http.StatusOK)
	tmpl.ExecuteTemplate(w, "post.html", data)
}

func (h *Handler) Api(w http.ResponseWriter, r *http.Request) {
	output := data
	jsonResponse, jsonError := json.Marshal(output)
	if jsonError != nil {
		fmt.Println("Unable to encode JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

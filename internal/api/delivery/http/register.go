package http

import "net/http"

func RegisterHTTPEndpoints(router *http.ServeMux) {
	fs := http.FileServer(http.Dir("./static"))
	h := NewHandler()
	router.Handle("/static/", http.StripPrefix("/static", fs))
	router.HandleFunc("/", h.GetPage)
	router.HandleFunc("/ascii-art", h.PostPage)
	router.HandleFunc("/api", h.Api)
}

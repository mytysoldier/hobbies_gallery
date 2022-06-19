package handler

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/template/index.html", "static/template/contents.html", "static/template/common/header_and_footer.html", "static/template/common/side_menu.html")
	t.ExecuteTemplate(w, "index", "")
}

func Handle() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", index)
	server.ListenAndServe()
}

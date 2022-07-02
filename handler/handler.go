package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

type Content struct {
	Image string
	Title string
}

func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/template/index.html", "static/template/contents.html", "static/template/pages/top.html", "static/template/common/header_and_footer.html", "static/template/common/side_menu.html")
	t.ExecuteTemplate(w, "index", "")
}

func category(w http.ResponseWriter, r *http.Request) {
	content := fmt.Sprintf("static/template/pages/%s.html", r.URL.Query().Get("content"))
	c1 := Content{
		Image: "https://bulma.io/images/placeholders/256x256.png",
		Title: "test1",
	}
	c2 := Content{
		Image: "https://bulma.io/images/placeholders/256x256.png",
		Title: "test2",
	}
	contents := []Content{c1, c2}
	t, _ := template.ParseFiles("static/template/index.html", "static/template/contents.html", content, "static/template/common/header_and_footer.html", "static/template/common/side_menu.html", "static/template/common/content_loop.html")
	t.ExecuteTemplate(w, "index", contents)
}

func Handle() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/cat", category)
	http.HandleFunc("/", index)
	server.ListenAndServe()
}
